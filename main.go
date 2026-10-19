package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var rx = regexp.MustCompile(` *\[remote "([^"]+)"] *`)

func main() {
	content, _ := os.ReadFile(".git/config")
	matches := rx.FindAllStringSubmatch(string(content), -1)
	for i := range matches {
		cmd := exec.Command("git", "push", matches[i][1])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			fmt.Println("Retrying once...")
			cmd = exec.Command("git", "push", matches[i][1])
			retry := cmd.Run()
			if retry != nil {
				_, _ = fmt.Fprintln(os.Stderr, retry)
			}
		}
	}
}
