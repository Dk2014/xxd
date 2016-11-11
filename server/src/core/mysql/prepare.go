package mysql

// #include <stdlib.h>
// #include <mysql.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type StmtError struct {
	Num  int
	Msg  string
	Stmt *Stmt
}

func (self *StmtError) Error() string {
	if len(self.Stmt.sql) == 0 {
		return fmt.Sprintf("%v (errno %v)", self.Msg, self.Num)
	}
	return fmt.Sprintf("%v (errno %v) during query: %s", self.Msg, self.Num, string(self.Stmt.sql))
}

type Stmt struct {
	st       *C.MYSQL_STMT
	binds    []C.MYSQL_BIND
	bind_pos int //绑定的位置
	sql      []byte
	Conn     *Connection
}

func (self *Connection) Prepare(sql []byte) (*Stmt, error) {
	stmt := new(Stmt)
	stmt.Conn = self
	stmt.sql = sql
	stmt.st = C.mysql_stmt_init(self.handle)

	if C.mysql_stmt_prepare(stmt.st, (*C.char)(unsafe.Pointer(&sql[0])), C.ulong(len(sql))) != 0 {
		return nil, &StmtError{Num: int(C.mysql_stmt_errno(stmt.st)), Msg: C.GoString(C.mysql_stmt_error(stmt.st)), Stmt: stmt}
	}

	bind_count := int(C.mysql_stmt_param_count(stmt.st))
	stmt.binds = make([]C.MYSQL_BIND, bind_count)

	return stmt, nil
}

func (stmt *Stmt) CleanBind() {
	stmt.bind_pos = 0
}

var (
	THE_TRUE  = true
	THE_FALSE = false
)

func (stmt *Stmt) BindInt(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_LONG
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))

	stmt.bind_pos++
}

func (stmt *Stmt) BindTinyInt(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_TINY
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

func (stmt *Stmt) BindSmallInt(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_SHORT
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

func (stmt *Stmt) BindBigInt(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_LONGLONG
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

func (stmt *Stmt) BindFloat(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_FLOAT
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

func (stmt *Stmt) BindDouble(valuePtr unsafe.Pointer) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_DOUBLE
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

func (stmt *Stmt) BindVarchar(valuePtr unsafe.Pointer, length int) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_VAR_STRING
	stmt.binds[stmt.bind_pos].buffer = valuePtr
	stmt.binds[stmt.bind_pos].buffer_length = (C.ulong)(length)
	stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	stmt.bind_pos++
}

// 备注：binary若有值肯定是固定长度的，那么长度就必须大于0，所以对应的go变量长度大于0；无值的时候为null，对应的go值为nil
func (stmt *Stmt) BindBinary(valuePtr unsafe.Pointer, length int) {
	stmt.binds[stmt.bind_pos].buffer_type = C.MYSQL_TYPE_VAR_STRING
	if valuePtr == nil {
		stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_TRUE))
	} else {
		stmt.binds[stmt.bind_pos].buffer = valuePtr
		stmt.binds[stmt.bind_pos].is_null = (*C.my_bool)(unsafe.Pointer(&THE_FALSE))
	}
	stmt.binds[stmt.bind_pos].buffer_length = (C.ulong)(length)
	stmt.bind_pos++
}

func (stmt *Stmt) Execute() error {
	C.mysql_stmt_bind_param(stmt.st, &stmt.binds[0])
	if C.mysql_stmt_execute(stmt.st) != 0 {
		return &StmtError{Num: int(C.mysql_stmt_errno(stmt.st)), Msg: C.GoString(C.mysql_stmt_error(stmt.st)), Stmt: stmt}
	}
	return nil
}

func (stmt *Stmt) Close() {
	C.mysql_stmt_close(stmt.st)
}
