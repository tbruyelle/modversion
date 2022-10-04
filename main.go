package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/blang/semver"
	"github.com/go-git/go-git/v5"
	"golang.org/x/mod/module"
)

func main() {
	var buf bytes.Buffer
	cmd := exec.Command("git", "describe", "--tags", "--exclude", "v0.0*")
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	ver, err := semver.ParseTolerant(buf.String())
	if err != nil {
		panic(err)
	}
	repo, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		panic(err)
	}

	var (
		major = fmt.Sprintf("v%d", ver.Major)
		older = fmt.Sprintf("v%d.%d.%d", ver.Major, ver.Minor, ver.Patch)
		rev   = ref.Hash().String()[:12]
		date  = commit.Author.When
	)
	v := module.PseudoVersion(major, older, date, rev)
	fmt.Println(v)
}
