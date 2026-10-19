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
			// Codeberg.org has a known issue where it will occasionally reject
			// a push for no reason. Therefore, we retry once if it fails.
			fmt.Println("Retrying once...")
			cmd = exec.Command("git", "push", matches[i][1])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			retry := cmd.Run()
			if retry != nil {
				_, _ = fmt.Fprintln(os.Stderr, retry)
			}
		}
	}
}
