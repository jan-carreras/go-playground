package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	FileType = "file"
	DirType  = "dir"

	DiskSize   = 70000000
	UpdateSize = 30000000
)

type File struct {
	name   string
	length int
	ftype  string // file / dir
	files  []*File
	parent *File
}

func (f *File) FullPath() string {
	if f.IsRoot() {
		return "/"
	}
	path := ""
	n := f
	for n.parent != nil {
		path = n.name + "/" + path
		n = n.parent
	}
	return "/" + path[:len(path)-1]
}

func (f *File) AddDir(name string) {
	if f.Find(name) != nil {
		return
	}

	newFile := &File{
		ftype:  DirType,
		name:   name,
		files:  make([]*File, 0),
		parent: f,
	}

	f.files = append(f.files, newFile)
}

func (f *File) AddFile(name string, size int) {
	found := f.Find(name)
	if found != nil {
		found.length = size // Update the length of the file
		return
	}

	newFile := &File{
		ftype:  FileType,
		name:   name,
		length: size,
		files:  make([]*File, 0),
		parent: f,
	}

	f.files = append(f.files, newFile)
}

func (f *File) Find(name string) *File {
	for _, file := range f.files {
		if file.name == name {
			return file
		}
	}

	return nil
}

func (f *File) IsRoot() bool {
	return f.parent == nil
}

func (f *File) IsDirectory() bool {
	return f.ftype == DirType
}

type Fs struct {
	currentDirectory *File
	root             *File
}

func (fs *Fs) ChangeDirectory(s string) error {
	if s == "/" {
		fs.currentDirectory = fs.root
	} else if s == ".." {
		if !fs.currentDirectory.IsRoot() {
			// We can not get "upper" than that
			fs.currentDirectory = fs.currentDirectory.parent
		}
	} else {
		f := fs.currentDirectory.Find(s)
		if f == nil {
			return fmt.Errorf("unknown directory %q", s)
		}

		fs.currentDirectory = f
	}

	return nil
}

func NewFs() Fs {
	f := &File{name: "/", length: 0, ftype: DirType}
	fs := Fs{currentDirectory: f, root: f}
	return fs
}

func parseInput(input io.Reader) (Fs, error) {
	fs := NewFs()

	const (
		readingCommand    = iota
		changingDirectory = iota
		parsingListOutput = iota
	)

	st := readingCommand

	b := bufio.NewScanner(input)
	for b.Scan() {
		t := b.Text()

		parseLineAgain := true
		for parseLineAgain {
			parseLineAgain = false

		PARSE:
			switch st {
			case readingCommand:
				if !strings.HasPrefix(t, "$ ") {
					return fs, fmt.Errorf("expected a command, %q found", t)
				}

				t = t[2:] // Ignore "$ "

				if strings.HasPrefix(t, "cd") {
					st = changingDirectory
					parseLineAgain = true
					break PARSE
				} else if strings.HasPrefix(t, "ls") {
					st = parsingListOutput
				} else {
					return fs, fmt.Errorf("unknown command: %s", t)
				}
			case changingDirectory:
				t = t[2:] // Ignore "cd"
				directory := strings.Trim(t, " ")
				if directory == "" {
					st = readingCommand
					continue // Do nothing: "$ cd" is a no-op
				}
				if err := fs.ChangeDirectory(directory); err != nil {
					return fs, fmt.Errorf("cannot change directory: %w", err)
				}
				st = readingCommand
			case parsingListOutput:
				if strings.HasPrefix(t, "$ ") {
					st = readingCommand
					parseLineAgain = true
					break PARSE
				} else if strings.HasPrefix(t, "dir ") {
					name := t[4:]
					fs.currentDirectory.AddDir(name)
				} else { // Parse file
					size, name := 0, ""
					if _, err := fmt.Sscanf(t, "%d %s", &size, &name); err != nil {
						return fs, fmt.Errorf("unable to parse file: %q", t)
					}
					fs.currentDirectory.AddFile(name, size)
				}
			}
		}
	}

	return fs, nil
}

func findSizeLimit(f *File, maxsize int, output map[string]int) int {
	if !f.IsDirectory() {
		return f.length
	}

	sum := 0
	for _, subFiles := range f.files {
		sum += findSizeLimit(subFiles, maxsize, output)
	}

	if sum <= maxsize {
		output[f.FullPath()] = sum
	}

	return sum
}

func WalkDirectories(f *File, fnx func(*File, int)) int {
	if !f.IsDirectory() {
		return f.length
	}

	sum := 0
	for _, subFiles := range f.files {
		sum += WalkDirectories(subFiles, fnx)
	}

	fnx(f, sum)

	return sum
}

func usedSize(f *File) int {
	if !f.IsDirectory() {
		return f.length
	}

	sum := 0
	for _, subFiles := range f.files {
		sum += usedSize(subFiles)
	}

	return sum
}
