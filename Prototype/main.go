package main

import (
	"fmt"
	p "test/prototype"
)

func main() {
	file1 := &p.File{Name: "File1"}
	file2 := &p.File{Name: "File2"}
	file3 := &p.File{Name: "File3"}

	folder1 := &p.Folder{
		Childrens: []p.INode{file1},
		Name:      "Folder1",
	}
	folder2 := &p.Folder{
		Childrens: []p.INode{
			folder1,
			file2,
			file3,
		},
		Name: "Folder2",
	}

	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
