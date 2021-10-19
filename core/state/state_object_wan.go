// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package state

import (

	"github.com/ethereum/go-ethereum/common"
)


type StorageByteArray map[common.Hash][]byte


// stateObject represents an Ethereum account which is being modified.
//
// The usage pattern is as follows:
// First you need to obtain a state object.
// Account values can be accessed and modified through the object.
// Finally, call CommitTrie to write the modified storage trie into a database.
type stateObject struct {
	address  common.Address
	addrHash common.Hash // hash of ethereum address of the account
	data     Account
	db       *StateDB

	// DB error.
	// State objects are used by the consensus core and VM which are
	// unable to deal with database-level errors. Any error that occurs
	// during a database read is memoized here and will eventually be returned
	// by StateDB.Commit.
	dbErr error

	// Write caches.
	trie Trie // storage trie, which becomes non-nil on first access
	code Code // contract bytecode, which gets set when code is loaded

	originStorage  Storage // Storage cache of original entries to dedup rewrites, reset for every transaction
	pendingStorage Storage // Storage entries that need to be flushed to disk, at the end of an entire block
	dirtyStorage   Storage // Storage entries that have been modified in the current transaction execution
	fakeStorage    Storage // Fake storage which constructed by caller for debugging purpose.

	// Cache flags.
	// When an object is marked suicided it will be delete from the trie
	// during the "update" phase of the state transition.
	dirtyCode bool // true if the code was updated
	suicided  bool
	deleted   bool

	// new
	cachedStorageByteArray StorageByteArray
	dirtyStorageByteArray  StorageByteArray
	cachedStorage Storage // Storage entry cache to avoid duplicate reads

	touched   bool
	onDirty   func(addr common.Address) // Callback method to mark a state object newly dirty
}


func (self *stateObject) GetStateByteArray(db Database, key common.Hash) []byte {
	value, exists := self.cachedStorageByteArray[key]
	if exists {
		return value
	}
	// Load from DB in case it is missing.
	value, err := self.getTrie(db).TryGet(key[:])
	if err == nil && len(value) != 0 {
		self.cachedStorageByteArray[key] = value
	}
	return value
}

func (self *stateObject) SetStateByteArray(db Database, key common.Hash, value []byte) {

	self.db.journal.append(storageByteArrayChange{
		account:  &self.address,
		key:      key,
		prevalue: self.GetStateByteArray(db, key),
	})
	self.setStateByteArray(key, value)

}


func (self *stateObject) setStateByteArray(key common.Hash, value []byte) {
	self.cachedStorageByteArray[key] = value
	self.dirtyStorageByteArray[key] = value

	if self.onDirty != nil {
		self.onDirty(self.Address())
		self.onDirty = nil
	}
}


func (self StorageByteArray) Copy() StorageByteArray {
	cpy := make(StorageByteArray)
	for key, value := range self {
		cpy[key] = value
	}

	return cpy
}