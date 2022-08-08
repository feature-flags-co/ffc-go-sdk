package datamodel

import "sync"

var dataStorageMap = make(map[Category]map[string]Item)
var lock sync.RWMutex

// Initialization Overwrites the storage with a set of items for each collection, if the new version > the old one
// @Param allData map of Category and their data set Item
// @Param version the version of dataset, Ordinarily it's a timestamp.
func Initialization(allData map[Category]map[string]Item, version int64) {

	lock.Lock()

	// TODO

	lock.Unlock()
}

// Get Retrieves an item from the specified collection, if available.
// @Param category specifies which collection to use
// @Param key the unique key of the item in the collection
// @Return a versioned item that contains the stored data or null if item is deleted or unknown
func Get(category Category, key string) Item {
	return Item{}
}

// GetAll Retrieves all items from the specified collection.
// @Param category specifies which collection to use
// @Return a map of ids and their versioned items
func GetAll(category Category) map[string]Item {

	return nil
}

// Upsert Updates or inserts an item in the specified collection. For updates, the object will only be
// updated if the existing version is less than the new version; for inserts, if the version > the existing one, it will replace
// the existing one.
// The SDK may pass an Item that contains a archived object,
// In that case, assuming the version is greater than any existing version of that item, the store should retain
// a placeholder rather than simply not storing anything.
// @Param category  specifies which collection to use
// @Param key the unique key of the item in the collection
// @Param item the item to insert or update
// @Param the version of item
// @Return true if success
func Upsert(category Category, key string, item Item, version int64) bool {
	return false
}

// IsInitialized Checks whether this store has been initialized with any data yet.
// @Return true if the storage contains data
func IsInitialized() bool {

	return false
}

// GetVersion return the latest version of storage
// @Return a long value
func GetVersion() int64 {

	return 0
}
