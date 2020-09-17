package ce

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var DefaultCe *ce
var DefaultCeByCommitHash string

type ce struct {
	*log.Logger
}

func init() {
	DefaultCe = New(os.Stderr, DefaultCeByCommitHash)
}

func New(out io.Writer, ceByCodeCommitHash ...string) *ce {
	if len(ceByCodeCommitHash) == 0 {
		ceByCodeCommitHash = []string{""}
	}
	_, zoneOffset := time.Now().Zone()
	zoneOffset = zoneOffset / 3600
	if out == nil {
		out = os.Stderr
	}
	return &ce{
		Logger: log.New(out, ceByCodeCommitHash[0]+"+"+strconv.Itoa(zoneOffset)+" ", log.LstdFlags|log.Lshortfile),
	}
}

func Print(v ...interface{}) {
	DefaultCe.Output(2, fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	DefaultCe.Output(2, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	DefaultCe.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	DefaultCe.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	DefaultCe.Output(2, s)
	panic(&panicByCe{OriginalErr: errors.New(s)})
}

func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	DefaultCe.Output(2, s)
	panic(&panicByCe{OriginalErr: errors.New(s)})
}

func CheckError(err error) {
	if err != nil {
		s := fmt.Sprint(err)
		DefaultCe.Output(2, s)
		panic(&panicByCe{OriginalErr: err})
	}
}

type sync interface {
	Sync()
}

func Sync() {
	if syncObj, ok := DefaultCe.Writer().(sync); ok {
		syncObj.Sync()
	}
}

type panicByCe struct {
	OriginalErr error
}

func (p *panicByCe) Error() string {
	return p.OriginalErr.Error()
}

func IsFromCe(errByRecover interface{}) (*panicByCe, bool) {
	me, ok := errByRecover.(*panicByCe)
	return me, ok
}
