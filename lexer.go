package ankaboot

type itemType int

const (
	itemError itemType = iota // Error
	itemEOF                   // End Of File
	itemLTag                  // Left Tag <
	itemRTag                  // Right Tag >
	itemHTML
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

type item struct {
	typ itemType
	val string
}

func lexItem() {

}
