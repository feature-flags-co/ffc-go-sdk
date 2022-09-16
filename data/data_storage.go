package data

import (
	"log"
	"sync"
)

type DataStorage interface {

	// Initialize Overwrites the storage with a set of items for each collection, if the new version > the old one
	// @Param allData map of Category and their data set Item
	// @Param version the version of dataset, Ordinarily it's a timestamp.
	Initialize(allData map[Category]map[string]Item, version int64)

	// Get Retrieves an item from the specified collection, if available.
	// @Param category specifies which collection to use
	// @Param key the unique key of the item in the collection
	// @Return a versioned item that contains the stored data or null if item is deleted or unknown
	Get(category Category, key string) Item

	// GetAll Retrieves all items from the specified collection.
	// @Param category specifies which collection to use
	// @Return a map of ids and their versioned items
	GetAll(category Category) map[string]Item

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
	Upsert(category Category, key string, item Item, version int64) bool

	// IsInitialized Checks whether this store has been initialized with any data yet.
	// @Return true if the storage contains data
	IsInitialized() bool

	// GetVersion return the latest version of storage
	// @Return a long value
	GetVersion() int64
}

type InMemoryDataStorage struct {
	initialized    bool
	lock           sync.RWMutex
	version        int64
	dataStorageMap map[Category]map[string]Item
}

var dataStorage *InMemoryDataStorage

func GetDataStorage() *InMemoryDataStorage {
	// if dataStorage not init
	if dataStorage == nil {
		dataStorage = &InMemoryDataStorage{
			initialized:    false,
			version:        -1,
			dataStorageMap: make(map[Category]map[string]Item, 0),
		}
	}
	return dataStorage
}

func (im *InMemoryDataStorage) Initialize(allData map[Category]map[string]Item, version int64) {
	if version < 0 || im.version >= version || allData == nil || len(allData) == 0 {
		return
	}
	im.lock.Lock()
	defer im.lock.Unlock()
	im.dataStorageMap = allData
	im.version = version
	im.initialized = true
	log.Printf("Data storage initialized")
}

func (im *InMemoryDataStorage) Get(category Category, key string) Item {

	im.lock.Lock()
	defer im.lock.Unlock()
	items := im.dataStorageMap[category]
	if items == nil {
		return Item{}
	}
	item := items[key]
	if item == (Item{}) || item.Item.Archived() {
		return Item{}
	}
	return item

}

func (im *InMemoryDataStorage) GetAll(category Category) map[string]Item {

	im.lock.Lock()
	defer im.lock.Unlock()
	// create new map for item save
	itemsMap := make(map[string]Item, 0)
	items := im.dataStorageMap[category]
	if items == nil {
		return itemsMap
	}
	for k, v := range items {
		if !v.Item.Archived() {
			itemsMap[k] = v
		}
	}
	return itemsMap
}

func (im *InMemoryDataStorage) Upsert(category Category, key string, item Item, version int64) bool {

	im.lock.Lock()
	defer im.lock.Unlock()
	if version < 0 || im.version >= version || item == (Item{}) || item.Item == nil {
		return false
	}

	oldItems := im.dataStorageMap[category]
	if oldItems != nil {
		oldItem := oldItems[key]
		if oldItem != (Item{}) && oldItem.Item.GetTimestamp() >= item.Item.GetTimestamp() {
			return false
		} else {
			oldItems[key] = item
			im.dataStorageMap[category] = oldItems
		}
	} else {
		tempMap := make(map[string]Item)
		tempMap[key] = item
		im.dataStorageMap[category] = tempMap
	}

	if !im.initialized {
		im.initialized = true
	}
	im.version = version
	log.Printf("upsert item %s into storage", key)
	return false
}

func (im *InMemoryDataStorage) IsInitialized() bool {

	im.lock.Lock()
	defer im.lock.Unlock()
	return im.initialized

}

func (im *InMemoryDataStorage) GetVersion() int64 {
	im.lock.Lock()
	defer im.lock.Unlock()
	return im.version
}
