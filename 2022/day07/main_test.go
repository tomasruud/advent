package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const inputFixture = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_parse(t *testing.T) {
	got, err := parse(inputFixture)

	e := newDir(
		"e",
		newFile("i", 584),
	)

	a := newDir(
		"a",
		newFile("f", 29116),
		newFile("g", 2557),
		newFile("h.lst", 62596),
	)

	e.parent = a
	a.dirs = newDirs(e)

	d := newDir(
		"d",
		newFile("j", 4060174),
		newFile("d.log", 8033020),
		newFile("d.ext", 5626152),
		newFile("k", 7214296),
	)

	root := newDir(
		"/",
		newFile("b.txt", 14848514),
		newFile("c.dat", 8504156),
	)

	a.parent = root
	d.parent = root
	root.dirs = newDirs(a, d)

	want := &fs{
		root: root,
	}

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_fs_totalSizeOfDirsWithMaxSize(t *testing.T) {
	got, _ := parse(inputFixture)

	assert.Equal(t, 95437, got.totalSizeOfDirsWithMaxSize(100000))
}

func Test_fs_selectFolderForFreeing(t *testing.T) {
	f, _ := parse(inputFixture)
	f.capacity = 70000000
	got, err := f.selectFolderForFreeing(30000000)

	assert.NoError(t, err)
	assert.Equal(t, 24933642, got.size())
}
