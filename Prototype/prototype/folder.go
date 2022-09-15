package prototype

import (
	"fmt"
)

//Folder is struct
type Folder struct {
	Childrens []INode
	Name	  string
}

//Print is function
func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, child := range f.Childrens {
		child.Print(s + s)
	}
}

//Clone is function
func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildrens []INode
	for _, child := range f.Childrens {
		tempChildrens = append(tempChildrens, child.Clone())
	}
	cloneFolder.Childrens = tempChildrens
	return cloneFolder
}