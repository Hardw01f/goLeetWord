package errorprint

import (
	"log"
	"os"
)

type LogFilters struct {
	Warn  *log.Logger
	Err   *log.Logger
	Debug *log.Logger
}

var logs LogFilters

func init() {
	logs.Err = log.New(os.Stdout, "[Error]:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logs.Warn = log.New(os.Stdout, "[WARNING]:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logs.Debug = log.New(os.Stdout, "[Debug]:", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

func ErrorMessage(err error) {
	logs.Err.Fatal(err.Error())
}

func WarningMessage(err error) {
	logs.Warn.Print(err.Error())
}

func DebugMessage(mes string) {
	logs.Debug.Println(mes)
}
