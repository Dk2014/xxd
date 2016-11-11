/*
Copyright 2012, Google Inc.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
    * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package mysql

// #cgo linux  CFLAGS:  -I /usr/include/mysql
// #cgo linux  LDFLAGS: -L /usr/lib/mysql -L /usr/lib64/mysql -lmysqlclient_r -lrt -ldl -lstdc++ -lm
// #cgo darwin CFLAGS:  -I /usr/local/include/mysql
// #cgo darwin LDFLAGS: -L /usr/local/lib/ -lmysqlclient_r
// #include <stdlib.h>
// #include "mysql.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type SqlError struct {
	Num     int
	Message string
	Query   string
}

func NewSqlError(number int, format string, args ...interface{}) *SqlError {
	return &SqlError{Num: number, Message: fmt.Sprintf(format, args...)}
}

func (self *SqlError) Error() string {
	if self.Query == "" {
		return fmt.Sprintf("%v (errno %v)", self.Message, self.Num)
	}
	return fmt.Sprintf("%v (errno %v) during query: %s", self.Message, self.Num, self.Query)
}

func (self *SqlError) Number() int {
	return self.Num
}

func handleError(err *error) {
	if x := recover(); x != nil {
		terr := x.(*SqlError)
		*err = terr
	}
}

type Connection struct {
	handle *C.MYSQL
}

type QueryResult struct {
	Fields       []Field
	FieldMap     map[string]int
	RowsAffected uint64
	InsertId     uint64
	Rows         []Row
	self         *Connection
	colCount     int
	result       *C.MYSQL_RES
}

type Field struct {
	Name string
	Type int64
}

func Connect(info map[string]interface{}) (conn *Connection, err error) {
	defer handleError(&err)

	host := mapCString(info, "host")
	defer cfree(host)
	port := mapint(info, "port")
	uname := mapCString(info, "uname")
	defer cfree(uname)
	pass := mapCString(info, "pass")
	defer cfree(pass)
	dbname := mapCString(info, "dbname")
	defer cfree(dbname)
	unix_socket := mapCString(info, "unix_socket")
	defer cfree(unix_socket)
	charset := mapCString(info, "charset")
	defer cfree(charset)

	conn = &Connection{}
	conn.handle = C.mysql_init(nil)
	if C.mysql_real_connect(conn.handle, host, uname, pass, dbname, C.uint(port), unix_socket, C.CLIENT_MULTI_STATEMENTS) == nil {
		defer conn.Close()
		return nil, conn.lastError(nil)
	}

	if C.mysql_set_character_set(conn.handle, charset) != 0 {
		defer conn.Close()
		return nil, conn.lastError(nil)
	}
	return conn, nil
}

func (self *Connection) ForbidAutoCommit() error {
	if C.mysql_autocommit(self.handle, 0) != 0 {
		return self.lastError([]byte("forbid autocommit"))
	}
	return nil
}

func (self *Connection) Commit() error {
	if C.mysql_commit(self.handle) != 0 {
		return self.lastError([]byte("commit error"))
	}
	return nil
}

func (self *Connection) RollBack() error {
	if C.mysql_rollback(self.handle) != 0 {
		return self.lastError([]byte("roll back error"))
	}
	return nil
}

func (self *Connection) ExecuteFetch(query []byte, maxrows int) (qr *QueryResult, err error) {
	defer handleError(&err)
	self.validate()

	if C.mysql_real_query(self.handle, (*C.char)(unsafe.Pointer(&query[0])), C.ulong(len(query))) != 0 {
		return nil, self.lastError(query)
	}

	result := C.mysql_store_result(self.handle)
	if result == nil {
		if int(C.mysql_field_count(self.handle)) != 0 { // Query was supposed to return data, but it didn't
			return nil, self.lastError(query)
		}
		qr = &QueryResult{}
		qr.RowsAffected = uint64(C.mysql_affected_rows(self.handle))
		qr.InsertId = uint64(C.mysql_insert_id(self.handle))
		return qr, nil
	}
	defer C.mysql_free_result(result)
	qr = &QueryResult{}
	qr.RowsAffected = uint64(C.mysql_affected_rows(self.handle))
	if maxrows >= 0 && qr.RowsAffected > uint64(maxrows) {
		return nil, &SqlError{0, fmt.Sprintf("Row count exceeded %d", maxrows), string(query)}
	}
	qr.Fields, qr.FieldMap = self.buildFields2(result)
	qr.Rows = self.fetchAll(result)
	return qr, nil
}

//使用mysql_use_result务必保证没有其他进程在读关联表
func (self *Connection) _ExecuteFetchAll(query []byte) (qr *QueryResult, err error) {
	defer handleError(&err)
	self.validate()
	if C.mysql_real_query(self.handle, (*C.char)(unsafe.Pointer(&query[0])), C.ulong(len(query))) != 0 {
		return nil, self.lastError(query)
	}

	result := C.mysql_use_result(self.handle)
	qr = &QueryResult{}
	if result == nil {
		return qr, nil
	}
	qr.Fields, qr.FieldMap = self.buildFields2(result)
	qr.self = self
	qr.colCount = len(qr.Fields)
	qr.result = result
	return qr, nil
}

func (self *Connection) ExecuteMultiSql(query []byte) (err error) {
	defer handleError(&err)
	self.validate()

	if C.mysql_real_query(self.handle, (*C.char)(unsafe.Pointer(&query[0])), C.ulong(len(query))) != 0 {
		return self.lastError(query)
	}
	result := C.mysql_store_result(self.handle)
	defer C.mysql_free_result(result)

	if int(C.mysql_more_results(self.handle)) == 1 {
		for {
			err := int(C.mysql_next_result(self.handle))
			if err == 0 {
				result := C.mysql_store_result(self.handle)
				defer C.mysql_free_result(result)
				continue
			} else {
				if err == 1 {
					return self.lastError(query)
				}
				break
			}
		}
	}
	return nil
}

func (self *Connection) Id() int64 {
	if self.handle == nil {
		return 0
	}
	return int64(C.mysql_thread_id(self.handle))
}

func (self *Connection) Close() {
	if self.handle == nil {
		return
	}
	C.mysql_close(self.handle)
	self.handle = nil
}

func (self *Connection) IsClosed() bool {
	return self.handle == nil
}

func (qr *QueryResult) Map(field_name string) int {
	if fi, ok := qr.FieldMap[field_name]; ok {
		return fi
	}
	return -1
}

func (qr *QueryResult) _GetRow() (row Row) {
	fr := C.mysql_fetch_row(qr.result)
	if fr == nil {
		defer C.mysql_free_result(qr.result)
		return nil
	}
	rowPtr := (*[1 << 30]*[1 << 30]byte)(unsafe.Pointer(fr))
	if rowPtr == nil {
		panic(qr.self.lastError(nil))
	}
	row = make([]interface{}, qr.colCount)
	lengths := (*[1 << 30]uint64)(unsafe.Pointer(C.mysql_fetch_lengths(qr.result)))

	totalLength := uint64(0)
	for i := 0; i < qr.colCount; i++ {
		totalLength += (*lengths)[i]
	}
	arena := NewStringArena(int(totalLength))
	for i := 0; i < qr.colCount; i++ {
		colLength := (*lengths)[i]
		colPtr := (*rowPtr)[i]
		if colPtr == nil {
			continue
		}
		row[i] = arena.NewString((*colPtr)[:colLength])
	}
	return row
}

func (qr *QueryResult) _Free() {
	C.mysql_free_result(qr.result)
}

func (self *Connection) buildFields(result *C.MYSQL_RES) (fields []Field) {
	nfields := int(C.mysql_num_fields(result))
	cfieldsptr := C.mysql_fetch_fields(result)
	cfields := (*[1 << 30]C.MYSQL_FIELD)(unsafe.Pointer(cfieldsptr))
	arena := NewStringArena(1024) // prealloc a reasonable amount of space
	fields = make([]Field, nfields)
	for i := 0; i < nfields; i++ {
		length := strlen(cfields[i].name)
		fname := (*[1 << 30]byte)(unsafe.Pointer(cfields[i].name))[:length]
		fields[i].Name = arena.NewString(fname)
		fields[i].Type = int64(cfields[i]._type)
	}
	return fields
}

func (self *Connection) buildFields2(result *C.MYSQL_RES) (fields []Field, fieldsmap map[string]int) {
	nfields := int(C.mysql_num_fields(result))
	cfieldsptr := C.mysql_fetch_fields(result)
	cfields := (*[1 << 30]C.MYSQL_FIELD)(unsafe.Pointer(cfieldsptr))
	arena := NewStringArena(1024) // prealloc a reasonable amount of space
	fields = make([]Field, nfields)
	fieldsmap = make(map[string]int, nfields)
	field_count := 0
	for i := 0; i < nfields; i++ {
		length := strlen(cfields[i].name)
		fname := (*[1 << 30]byte)(unsafe.Pointer(cfields[i].name))[:length]
		name := arena.NewString(fname)
		fields[i].Name = name
		fields[i].Type = int64(cfields[i]._type)
		fieldsmap[name] = field_count
		field_count++
	}
	return fields, fieldsmap
}

func (self *Connection) fetchAll(result *C.MYSQL_RES) (rows []Row) {
	rowCount := int(C.mysql_num_rows(result))
	if rowCount == 0 {
		return nil
	}
	rows = make([]Row, rowCount)
	colCount := int(C.mysql_num_fields(result))
	for i := 0; i < rowCount; i++ {
		rows[i] = self.fetchNext(result, colCount)
	}
	return rows
}

func (self *Connection) fetchNext(result *C.MYSQL_RES, colCount int) (row Row) {
	rowPtr := (*[1 << 30]*[1 << 30]byte)(unsafe.Pointer(C.mysql_fetch_row(result)))
	if rowPtr == nil {
		panic(self.lastError(nil))
	}
	row = make([]interface{}, colCount)
	lengths := (*[1 << 30]uint64)(unsafe.Pointer(C.mysql_fetch_lengths(result)))

	totalLength := uint64(0)
	for i := 0; i < colCount; i++ {
		totalLength += (*lengths)[i]
	}
	arena := NewStringArena(int(totalLength))
	for i := 0; i < colCount; i++ {
		colLength := (*lengths)[i]
		colPtr := (*rowPtr)[i]
		if colPtr == nil {
			continue
		}
		row[i] = arena.NewString((*colPtr)[:colLength])
	}
	return row
}

func (self *Connection) lastError(query []byte) error {
	if err := C.mysql_error(self.handle); *err != 0 {
		return &SqlError{Num: int(C.mysql_errno(self.handle)), Message: C.GoString(err), Query: string(query)}
	}
	return &SqlError{0, "Dummy", string(query)}
}

func (self *Connection) validate() {
	if self.handle == nil {
		panic(NewSqlError(2006, "Connection is closed"))
	}
}

func mapCString(info map[string]interface{}, key string) *C.char {
	ival, ok := info[key]
	if !ok {
		panic(NewSqlError(0, "Missing connection parameter %s", key))
	}
	sval, ok := ival.(string)
	if !ok {
		panic(NewSqlError(0, "Expecting string for %s, received %T", key, ival))
	}
	if sval == "" {
		return nil
	}
	return C.CString(sval)
}

func mapint(info map[string]interface{}, key string) int {
	ival, ok := info[key]
	if !ok {
		panic(NewSqlError(0, "Missing connection parameter %s", key))
	}
	intval, ok := ival.(int)
	if !ok {
		panic(NewSqlError(0, "Expecting int for %s, received %T", key, ival))
	}
	return intval
}

func cfree(str *C.char) {
	if str != nil {
		C.free(unsafe.Pointer(str))
	}
}

// Muahahaha
func strlen(str *C.char) int {
	gstr := (*[1 << 30]byte)(unsafe.Pointer(str))
	length := 0
	for gstr[length] != 0 {
		length++
	}
	return length
}
