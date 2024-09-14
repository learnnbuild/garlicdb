package index

import (
	"github.com/learnnbuild/garlicdb/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Unit tests for btree methods

func TestBTree_Put(t *testing.T) {
	//nil key, false
	newTree := NewBTree()
	res1 := newTree.Put(nil, &data.LogRecordPos{
		Fid:    1,
		Offset: 2,
	})

	assert.False(t, res1)
	// non-nil key, true

	res2 := newTree.Put([]byte("key1"), &data.LogRecordPos{
		Fid:    2,
		Offset: 3,
	})
	assert.True(t, res2)

	//non nil key and key already exists
	res3 := newTree.Put([]byte("key1"), &data.LogRecordPos{
		Fid:    2,
		Offset: 4,
	})
	assert.True(t, res3)
}

func TestBTree_Get(t *testing.T) {
	newTree := NewBTree()

	// nil key
	res1 := newTree.Get(nil)
	assert.Nil(t, res1)

	//key doesn't exist
	res2 := newTree.Get([]byte("key1"))
	assert.Nil(t, res2)

	//key exists
	_ = newTree.Put([]byte("key1"), &data.LogRecordPos{
		Fid:    1,
		Offset: 2,
	})

	res3 := newTree.Get([]byte("key1"))
	assert.Equal(t, uint32(1), res3.Fid)
	assert.Equal(t, int64(2), res3.Offset)

	//key exists and value updated
	_ = newTree.Put([]byte("key1"), &data.LogRecordPos{
		Fid:    1,
		Offset: 3,
	})
	res4 := newTree.Get([]byte("key1"))
	assert.Equal(t, uint32(1), res4.Fid)
	assert.Equal(t, int64(3), res4.Offset)

}

func TestBTree_Delete(t *testing.T) {
	// nil key

	newTree := NewBTree()

	res1 := newTree.Delete(nil)
	assert.False(t, res1)

	//non-nil key but key doesn't exist
	//
	res2 := newTree.Delete([]byte("key1"))
	assert.False(t, res2)
	//
	////key exists
	_ = newTree.Put([]byte("key1"), &data.LogRecordPos{
		Fid:    1,
		Offset: 2,
	})

	res3 := newTree.Delete([]byte("key1"))
	res4 := newTree.Get([]byte("key1"))
	assert.True(t, res3)
	assert.Nil(t, res4)
	//
	//// delete twice
	_ = newTree.Put([]byte("key2"), &data.LogRecordPos{
		Fid:    1,
		Offset: 2,
	})

	res5 := newTree.Delete([]byte("key2"))
	assert.True(t, res5)

	res6 := newTree.Delete([]byte("key2"))
	assert.False(t, res6)
	//
}
