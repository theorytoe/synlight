//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	out, err := os.Create("lexers.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't create: %s", err)
		os.Exit(1)
	}
	io.WriteString(out, "package synlight\n\n")
	for _, fname := range os.Args[1:] {
		fd, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't open: %s", err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "var lexer_%s = \"\" + \n", fname[7:len(fname)-4])
		reader := bufio.NewReader(fd)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "\\", "\\\\")
			line = strings.ReplaceAll(line, "\"", "\\\"")
			io.WriteString(out, "\"")
			io.WriteString(out, line)
			io.WriteString(out, "\\n\" + \n")
		}
		io.WriteString(out, "\"\"\n")
		fd.Close()
	}
	out.Close()
}
