package main

import (
	"sync"
	"time"

	"golang.org/x/text/language"
)

type CatchKey struct {
	FromLanguage language.Tag
	ToLanguage   language.Tag
	Data         string
}

type CacheMap struct {
	cm map[*CatchKey]CatchValue
	cl sync.Mutex
}

type CatchValue struct {
	value     string
	createdAt int64
}

func CreateCatch(ExpInMin int64) *CacheMap {
	cm := &CacheMap{cm: make(map[*CatchKey]CatchValue)}
	go func() {
		for now := range time.Tick(time.Second) {
			cm.cl.Lock()
			for k, v := range cm.cm {
				if now.Unix()-v.createdAt > ExpInMin {
					delete(cm.cm, k)
				}
			}
			cm.cl.Unlock()
		}
	}()
	return cm
}

func UpdateCache(cm *CacheMap, pKey *CatchKey, pData CatchValue) {
	cm.cl.Lock()
	cm.cm[pKey] = pData

	cm.cl.Unlock()
}

func GetCache(cm *CacheMap, pKey *CatchKey) (CatchValue, bool) {
	cm.cl.Lock()
	CachedData, found := cm.cm[pKey]
	cm.cl.Unlock()
	return CachedData, found
}
