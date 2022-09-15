package prototype

//INode is the interface for the prototype pattern
type INode interface{
	Clone() INode
	Print(s string)
}