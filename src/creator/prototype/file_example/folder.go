package main

import "fmt"

// Folder 具体原型
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var temChildren []Inode
	for _, i := range f.children {
		copyI := i.clone()
		temChildren = append(temChildren, copyI)
	}
	cloneFolder.children = temChildren
	return cloneFolder
}
