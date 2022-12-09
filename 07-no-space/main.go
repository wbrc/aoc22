package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	root := newDir(nil)
	currentDir := root

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if strings.HasPrefix(s.Text(), "$ cd ") {
			newDir := strings.TrimPrefix(s.Text(), "$ cd ")
			if newDir == "/" {
				currentDir = root
			} else if newDir == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.dirs[newDir]
			}
		} else if s.Text() == "$ ls" {
			continue
		} else if strings.HasPrefix(s.Text(), "dir ") {
			var dirname string
			_, err := fmt.Sscanf(s.Text(), "dir %s", &dirname)
			if err != nil {
				log.Fatal("parse error")
			}
			currentDir.dirs[dirname] = newDir(currentDir)
		} else {
			var (
				size int
				name string
			)
			_, err := fmt.Sscanf(s.Text(), "%d %s", &size, &name)
			if err != nil {
				log.Fatal("parse error")
			}

			currentDir.files[name] = &file{
				size: size,
			}

			fix(currentDir)
		}
	}

	totalSize := 70000000
	needSize := 30000000
	actualSize := root.size
	freeUp := needSize - totalSize + actualSize

	var dirSizes []int

	walk(root, func(d *dir) {
		if d.size >= freeUp {
			dirSizes = append(dirSizes, d.size)
		}
	})

	sort.Ints(dirSizes)

	fmt.Println(dirSizes[0])
}

type file struct {
	size int
}

type dir struct {
	parent *dir
	size   int
	dirs   map[string]*dir
	files  map[string]*file
}

func newDir(parent *dir) *dir {
	return &dir{
		parent: parent,
		dirs:   make(map[string]*dir),
		files:  make(map[string]*file),
	}
}

func fix(currentDir *dir) {
	size := 0

	for _, dir := range currentDir.dirs {
		size += dir.size
	}

	for _, file := range currentDir.files {
		size += file.size
	}

	currentDir.size = size

	if currentDir.parent != nil {
		fix(currentDir.parent)
	}
}

func walk(d *dir, f func(*dir)) {
	f(d)
	for _, nestedDir := range d.dirs {
		walk(nestedDir, f)
	}
}
