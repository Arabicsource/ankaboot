package ankaboot

// Greatly inspired by Rob Pike's talk on Lexical Analysis in Go
// and borrowed a lot of the code from slide, to get it to work.
// Will require to modify the code to be more suitable for my own
// use case.
//
type itemType int

// Define all of the 'items' the lexer will need to lex.
const (
	itemError itemType = iota // Error
	itemEOF                   // End Of File
	itemLTag                  // Left Tag <
	itemRTag                  // Right Tag >
	itemHTML                  // HTML tag
	itemHead
	itemLink
	itemTitle
	itemStyle
	itemBody
	itemScript
	itemDiv
	itemP
	itemA
	itemImg
	itemUl
	itemLi
	itemOl
	itemText
	itemId
	itemClass
	itemHref
)

// item struct that will be passed through lexer items channel to the
// parser.
type item struct {
	typ itemType
	val string
}

// lexer struct
type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan item
}

// lex initialises the lexer
func lex(name, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}

	go l.run()
	return l, l.items
}

// run lexer
func (l *lexer) run() {
	// here will run
	close(l.items)
}

// emit tokens
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos

}

//close open channel
func close(items chan item) {
	// close channel
}
