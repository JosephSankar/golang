package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func main() {
	p := Pair{"/usr", "0xfdfe"}

	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}

	fmt.Println(p)
	fmt.Println(fn)
	fmt.Println(fn.Filename())
}
