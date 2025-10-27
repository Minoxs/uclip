package main

import (
	"fmt"
	"io"
	"os"

	"github.com/minoxs/uclip/clipboard"
)

func IsInteractive() bool {
	var stat, err = os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	return stat.Mode()&os.ModeCharDevice != 0
}

func ConsumeStdin() string {

	var bytes, err = io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func SetClipboardString(value string) {
	if err := clipboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		if err := clipboard.Close(); err != nil {
			fmt.Println("closing", err)
		}
	}()

	if err := clipboard.SetClipboardData(value); err != nil {
		panic(err)
	}
}

const HELP = `UCLIP

Description:
Redirects output of command line tools to the Windows clipboard.
This text output can then be pasted into other programs.

This tool expects that only valid UTF8 text will be passed onto it.
To copy files and images use the standard clip.exe bundled with windows.

Examples:
(Piping)
cat file.txt | uclip  Places a copy of the text from stdin into windows clipboard.
(File Redirection)
uclip < file.txt      Places a copy of the file content into the windows clipboard.
`

func main() {
	if IsInteractive() {
		fmt.Println(HELP)
		return
	}

	var input = ConsumeStdin()
	SetClipboardString(input)
}
