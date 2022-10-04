package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"golang.org/x/mod/module"
)

func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	tag, err := repo.TagObject(ref.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println(tag)

	module.PseudoVersion("0", "v0.24.0", time.Now(), "")
}
