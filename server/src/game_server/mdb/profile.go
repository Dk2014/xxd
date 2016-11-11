package mdb

import (
	"os"
)

type XXDProfile interface {
	ProfilePrintToLog()
	ProfileWriteToFile(*os.File)
}
