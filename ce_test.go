package ce

import (
	"io"
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

func TestCheckError(t *testing.T) {
	CheckError(io.EOF)
}

func TestPanic(t *testing.T) {
	Panic("")
}

func TestFatal(t *testing.T) {
	Fatal("")
}
