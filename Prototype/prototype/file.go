package prototype

import (
	"fmt"
)

//File is struct
type File struct {
	Name string
}

//Print is function
func (f *File) Print(s string) {
	fmt.Println(s + f.Name)
}

//Clone is function
func (f *File) Clone() INode {
	return &File{Name: f.Name + " _Clone"}
}