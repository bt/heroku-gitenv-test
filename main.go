package main

import (
	"fmt"
	"github.com/bt/heroku-gitenv-test/appver"
)

func main() {
	fmt.Printf("Application version %s, commit %s\n", appver.Version, appver.Commit)
	fmt.Printf("%s branch, built on %s\n", appver.Branch, appver.BuildTime)
}