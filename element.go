package ankaboot

type Element struct {
	Name       string
	Properties map[string][]string // element["a"] -> Property["href"] -> ["http://www.google.com"]
	Next       *Element
	Prev       *Element
}

func (e *Element) NextElement() *Element {
	/// Go to next element in the tree
	return e.Next
}

func (e *Element) PreviousElement() *Element {
	// go to previous Element in tree
	return nil
}

func (e *Element) Parent() *Element {
	// return parent of the Element
	return nil
}

func (e *Element) Children() []*Element {
	// return slice of Elements that are the children of the Element
	return nil
}
