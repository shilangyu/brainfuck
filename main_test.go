package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/require"
)

func Test_Main(t *testing.T) {
	req := require.New(t)

	tests := []struct {
		command, output, stdin string
	}{
		{
			command: "+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.",
			output:  "hello world",
		},
		{
			command: ",>,>,.<.<.",
			output:  "yay",
			stdin:   "yay",
		},
		{
			command: "+++++[>+++++++++++++<-]>.",
			output:  "A",
		},
	}

	for _, tt := range tests {
		tempFile, err := ioutil.TempFile(".", "tempbf")
		defer os.Remove(tempFile.Name())
		req.NoError(err)

		tempFile.Write([]byte(tt.command))

		argss := [][]string{
			{tempFile.Name()},
			{"-c", tt.command},
		}
		for _, args := range argss {
			os.Args = append([]string{""}, args...)
			out := capturer.CaptureStdout(func() {
				r, w, _ := os.Pipe()
				os.Stdin = r
				fmt.Fprint(w, tt.stdin)
				main()
			})
			req.Equal(tt.output, out)
		}
	}
}
