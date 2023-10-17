package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f, err := parse(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("total size for dirs", f.totalSizeOfDirsWithMaxSize(100000))

	f.capacity = 70000000
	tot, err := f.selectFolderForFreeing(30000000)
	if err != nil {
		panic(err)
	}
	fmt.Println("total size to free", tot.size())
}

func parse(in string) (*fs, error) {
	system := &fs{
		root: newDir("/"),
	}
	current := system.root

	var cmd string
	for _, ln := range strings.Split(in, "\n") {
		ln = strings.TrimSpace(ln)

		if strings.HasPrefix(ln, "$") {
			raw := strings.Split(ln, " ")
			if len(raw) < 1 {
				return system, fmt.Errorf("no command given: %s", ln)
			}

			cmd = raw[1]
			arg := ""
			if len(raw) > 2 {
				arg = raw[2]
			}

			switch cmd {
			case "cd":
				switch arg {
				case "":
					return system, fmt.Errorf("no arguments given to cd command: %s", ln)

				case "/":
					current = system.root

				case "..":
					if current.parent == nil {
						return system, fmt.Errorf("no parent dir to go to: %s", ln)
					}

					current = current.parent

				default:
					sub, ok := current.subdir(arg)
					if !ok {
						return system, fmt.Errorf("no sub dir to go to: %s", ln)
					}

					current = sub
				}
				continue

			case "ls":
				continue

			default:
				return system, fmt.Errorf("unknown command: %s", cmd)
			}
		} else {
			switch cmd {
			case "ls":
				res := strings.Split(ln, " ")
				if len(res) != 2 {
					return system, fmt.Errorf("unable to split output line: %s", ln)
				}

				if res[0] == "dir" {
					d := newDir(res[1])
					d.parent = current
					current.dirs = append(current.dirs, d)
					continue
				}

				size, err := strconv.Atoi(res[0])
				if err != nil {
					return system, fmt.Errorf("unable to parse file size: %s", ln)
				}

				current.files = append(current.files, newFile(res[1], size))
			}
		}
	}
	return system, nil
}

type fs struct {
	root     *directory
	capacity int
}

func (f *fs) totalSizeOfDirsWithMaxSize(size int) int {
	return f.root.dirsWithMaxSize(size).size()
}

func (f *fs) selectFolderForFreeing(want int) (*directory, error) {
	free := f.capacity - f.root.size()
	need := want - free

	if need < 0 {
		return nil, fmt.Errorf("space is already available")
	}

	var min *directory
	for _, d := range f.root.dirsWithMinSize(need) {
		if min == nil {
			min = d
			continue
		}

		if d.size() < min.size() {
			min = d
		}
	}

	if min == nil {
		return nil, fmt.Errorf("there are no suitable matches")
	}

	return min, nil
}

type directories []*directory

func (ds directories) size() int {
	var sum int
	for _, d := range ds {
		sum += d.size()
	}
	return sum
}

type directory struct {
	name   string
	parent *directory
	dirs   directories
	files  []*file
}

func (dir *directory) dirsWithMaxSize(size int) directories {
	var ds directories

	if dir.size() <= size {
		ds = append(ds, dir)
	}

	for _, d := range dir.dirs {
		ds = append(ds, d.dirsWithMaxSize(size)...)
	}

	return ds
}

func (dir *directory) dirsWithMinSize(size int) directories {
	var ds directories

	if dir.size() >= size {
		ds = append(ds, dir)
	}

	for _, d := range dir.dirs {
		ds = append(ds, d.dirsWithMinSize(size)...)
	}

	return ds
}

func (dir *directory) subdir(name string) (*directory, bool) {
	for _, d := range dir.dirs {
		if d.name == name {
			return d, true
		}
	}

	return nil, false
}

func (dir *directory) size() int {
	var sum int
	for _, f := range dir.files {
		sum += f.size
	}

	for _, d := range dir.dirs {
		sum += d.size()
	}

	return sum
}

type file struct {
	name string
	size int
}

func newDirs(dirs ...*directory) directories {
	return dirs
}

func newDir(name string, files ...*file) *directory {
	return &directory{
		name:  name,
		files: files,
	}
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}
