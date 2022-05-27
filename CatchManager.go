package main

//This module is for caching purpose.

import (
	"sync"
	"time"
)

// First structure for from language , To language and data
// We are making composite key of FL, TL and data
type CatchKey struct {
	FromLanguage string
	ToLanguage   string
	Data         string
}

// Second structure to create map. (It contains map structure and mutex.)
// Here we are using map functionality to handle catche functionality
// As I am thinking that map is best in memory structure instead of using
// third party catche technique.
type CacheMap struct {
	cm map[CatchKey]CatchValue
	cl sync.Mutex
}

// Third structure to handle value of catche
// Here we are also capturing creation time of catch which is required to delete catche
type CatchValue struct {
	value     string
	createdAt int64
}

func ValidateCatch(cm *CacheMap, ExpInMin int64) {

	//while deleting catch, we are locking catch object to prevent unnecessary access by other thread.
	cm.cl.Lock()
	for k, v := range cm.cm {
		//If any key is expired then deleting that key from catched map
		//Converting minutes to seconds
		if int64(time.Now().Unix())-v.createdAt > (ExpInMin * 60) {

			delete(cm.cm, k)
		}
	}

	//We are relesing lock after fulfillment of operation.
	cm.cl.Unlock()
	defer wgrp.Done() //To decrease waitgroup by one.

}
func CreateCatch() *CacheMap {

	//Initializing catch map object
	cm := &CacheMap{cm: make(map[CatchKey]CatchValue)}
	return cm
}

//Inserting new key with values in catch if it is not available
func UpdateCache(cm *CacheMap, pKey CatchKey, pData CatchValue) bool {
	Updated := false
	cm.cl.Lock()
	cm.cm[pKey] = pData
	cm.cl.Unlock()
	Updated = true
	return Updated
}

//Reading catched value of a particular key
func GetCache(cm *CacheMap, pKey CatchKey) (CatchValue, bool) {
	cm.cl.Lock()
	CachedData, found := cm.cm[pKey]
	cm.cl.Unlock()
	return CachedData, found
}
