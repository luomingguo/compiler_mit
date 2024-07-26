package bptree

import (
	"github.com/timtadh/data-structures/errors"
	"github.com/timtadh/data-structures/types"
)

/* A BpMap is a B+Tree with support for duplicate keys disabled. This makes it
 * behave like a regular Map rather than a MultiMap.
 */
type BpMap BpTree

func NewBpMap(node_size int) *BpMap {
	return &BpMap{
		root: NewLeaf(node_size, true),
		size: 0,
	}
}

func (self *BpMap) Size() int {
	return (*BpTree)(self).Size()
}

func (self *BpMap) Has(key types.Hashable) bool {
	return (*BpTree)(self).Has(key)
}

func (self *BpMap) Put(key types.Hashable, value interface{}) (err error) {
	had := self.Has(key)
	new_root, err := self.root.put(key, value)
	if err != nil {
		return err
	}
	self.root = new_root
	if !had {
		self.size += 1
	}
	return nil
}

func (self *BpMap) Get(key types.Hashable) (value interface{}, err error) {
	j, l := self.root.get_start(key)
	if l.keys[j].Equals(key) {
		return l.values[j], nil
	}
	return nil, errors.NotFound(key)
}

func (self *BpMap) Remove(key types.Hashable) (value interface{}, err error) {
	value, err = self.Get(key)
	if err != nil {
		return nil, err
	}
	ns := self.root.NodeSize()
	new_root, err := self.root.remove(key, func(value interface{}) bool { return true })
	if err != nil {
		return nil, err
	}
	if new_root == nil {
		self.root = NewLeaf(ns, true)
	} else {
		self.root = new_root
	}
	self.size--
	return value, nil
}

func (self *BpMap) Keys() (ki types.KIterator) {
	return (*BpTree)(self).Keys()
}

func (self *BpMap) Values() (vi types.Iterator) {
	return (*BpTree)(self).Values()
}

func (self *BpMap) Iterate() (kvi types.KVIterator) {
	return (*BpTree)(self).Iterate()
}
