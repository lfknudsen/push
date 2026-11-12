package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var rx = regexp.MustCompile(` *\[remote "([^"]+)"] *`)

var tag = flag.String("tag", "", "Tag to push to each remote repository.")
var help = flag.Bool("h", false, "Print this message and exit.")

func printHelp() {
	fmt.Println("Usage: push [OPTIONS]\n\nOptions:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help {
		printHelp()
		os.Exit(0)
	}

	content, _ := os.ReadFile(".git/config")
	matches := rx.FindAllStringSubmatch(string(content), -1)
	for i := range matches {
		fmt.Printf("\033[34mPushing to %s:\033[0m\n", matches[i][1])
		cmd := exec.Command("git", "push", matches[i][1])
		if *tag != "" {
			cmd.Args = append(cmd.Args, "tag", *tag)
		}
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
