package index

import (
	"bytes"
	"github.com/google/btree"
	"github.com/learnnbuild/garlicdb/data"
)

type Indexer interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}

// Item: implement the Item Interface in Google BTree in order to use the methods provided by Google BTree
type Item struct {
	key []byte
	pos *data.LogRecordPos
}

// Implement the method in the btree.Item interface, -1 means the key of operand one is strictly less than the key of operand two
func (operandOne *Item) Less(operandTwo btree.Item) bool {
	return bytes.Compare(operandOne.key, operandTwo.(*Item).key) == -1
}
