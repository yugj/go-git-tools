package main

import (
	"go-git-tools/cmd/autosummary"
	"go-git-tools/cmd/git"
	"log"
	"os"
	"path/filepath"
)

func main()  {

	log.Println(os.Args[0])
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("failed to dir")
	}

	log.Println("base dir :{}", dir)

	autosummary.Generate(dir)

	git.Execute()
}