package main

import (
	"flag"
	"fmt"
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path to list files from, recursively")
	flag.Parse()

	if path == "" {
		fmt.Println("You must specify a path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	q, err := scanFiles(path)
	if err != nil {
		log.Fatalln(err)
	}

	printFiles(q)
}

func printFiles(q *adt.Queue) {
	for q.Len() != 0 {
		v, _ := q.Dequeue()
		fmt.Println(v)
	}
}

func scanFiles(path string) (*adt.Queue, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path = filepath.Join(pwd, path)

	q := &adt.Queue{}

	walk := func(path string, info fs.FileInfo, err error) error {
		base := strings.ReplaceAll(filepath.Dir(path), pwd, "")
		c := strings.Count(base, "/")
		f := strings.Repeat("\t", c) + filepath.Base(path)
		q.Enqueue(f)
		return nil
	}

	if err = filepath.Walk(path, walk); err != nil {
		return nil, err
	}
	return q, nil
}
