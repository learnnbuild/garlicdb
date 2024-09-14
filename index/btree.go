package index

import (
	"github.com/google/btree"
	"github.com/learnnbuild/garlicdb/data"
	"sync"
)

// BTree: encapsulate btree by Google, support concurrent write
// To use the methods provided by Google BTree, we need to implement Item interface in the Google BTree Library
// After we implement the Item interface, we need to implement the Indexer interface
type BTree struct {
	tree *btree.BTree  // tree: not write safe
	lock *sync.RWMutex //lock: to ensure the write operations are thread-safe
}

// Init a new BTree instance
func NewBTree() *BTree {
	return &BTree{
		//To be optimized: degree could be a parameter for this function
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

// Implement the Put function in the Indexer interface, lock the write opertions when perform write
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	if keyIsInvalid(key) {
		return false
	}
	oneItem := &Item{
		key: key,
		pos: pos,
	}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(oneItem)
	bt.lock.Unlock()
	return true
}

// Implement the Delete function in the Indexer interface
func (bt *BTree) Delete(key []byte) bool {
	if keyIsInvalid(key) {
		return false
	}
	oneItem := &Item{
		key: key,
	}
	bt.lock.Lock()
	oldRecord := bt.tree.Delete(oneItem)
	bt.lock.Unlock()
	if oldRecord == nil {
		return false
	}

	return true
}

// Implement the Get function in the Indexer interface, Get method in Google BTree is thread safe
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	if keyIsInvalid(key) {
		return nil
	}
	oneItem := &Item{
		key: key,
	}

	returned := bt.tree.Get(oneItem)
	if returned == nil {
		return nil
	}
	//convert the returned item to an Item we have defined ourselves
	return returned.(*Item).pos
}
