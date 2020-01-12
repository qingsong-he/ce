package ce

import (
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	Print(1, 2.2)
	Print("hello")
	Print(os.Args)
}

func TestPrintf(t *testing.T) {
	Printf("%T", 2.2)
}

func TestCheckError(t *testing.T) {
	CheckError(nil)
}

func TestDebug(t *testing.T) {
	Debug("")
}

func TestWarn(t *testing.T) {
	Warn("")
}

func TestInfo(t *testing.T) {
	Info("")
}

func TestError(t *testing.T) {
	Error("")
}

func TestPanic(t *testing.T) {
	Panic("")
}

func TestFatal(t *testing.T) {
	Fatal("")
}
