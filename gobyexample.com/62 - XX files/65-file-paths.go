package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir//", "filename"))
	fmt.Println(filepath.Join("dir/../dir", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println("is abs", filepath.IsAbs("dir/file"))
	fmt.Println("is abs", filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println("ext of", filename, "is", ext)

	fmt.Println("trim suffix of", filename, "=", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil { panic(err) }
	fmt.Println("relative path a/b and a/b/t/file", rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil { panic(err) }
	fmt.Println("relative path a/b and a/c/t/file", rel)
}