package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var sformat = "%s\n"

/**
  Captures stdout of wrapped function
*/

func captureStdout(f func(request interface{}, format string, v ...interface{}), request interface{},
	format string, v ...interface{}) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(request, format, v)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func TestLogger(t *testing.T) {
	Jsonlog = true
	t.Run("Info", func(t *testing.T) {
		message := captureStdout(Info, "http", sformat, "Fuck a Duck")
		fmt.Printf("==============> %s\n", message)
	})
	t.Run("Warning", func(t *testing.T) {
		message := captureStdout(Warning, "http", sformat, "Fuck a Duck")
		fmt.Printf("==============> %s\n", message)
	})
	t.Run("Debug", func(t *testing.T) {
		message := captureStdout(Debug, "http", sformat, "Fuck a Duck")
		fmt.Printf("==============> %s\n", message)
	})
	t.Run("Error", func(t *testing.T) {
		message := captureStdout(Error, "http", sformat, "Fuck a Duck")
		fmt.Printf("==============> %s\n", message)
	})
}
