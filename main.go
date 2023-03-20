package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	r, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}
	if status.IsClean() {
		fmt.Println("Clean")
	} else {
		fmt.Println("Dirty")
	}
}
