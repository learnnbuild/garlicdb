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

// Implement the Put function in the Indexer interface, lock the write opertions when perform write
func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
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
	oneItem := &Item{
		key: key,
	}
	bt.lock.Lock()
	oldRecord := bt.tree.Delete(oneItem)
	if oldRecord == nil {
		return false
	}
	bt.lock.Unlock()
	return true
}

// Implement the Get function in the Indexer interface, Get method in Google BTree is thread safe
func (bt *BTree) Get(key []byte) *data.LogRecordPos {
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
