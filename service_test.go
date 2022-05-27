package main

import (
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestReadConf(t *testing.T) {
	ExpOutput_CacheExpirationMin := 30
	ExpOutput_RetryReq := 3

	var testConf = ReadConf()

	if testConf["CacheExpirationMin"] != ExpOutput_CacheExpirationMin {
		t.Errorf("Test 1 Fail: Actual value of cache expiration is not matching with expected value.(%d != %d) ",
			testConf["CacheExpirationMin"], ExpOutput_CacheExpirationMin)
	} else {
		t.Errorf("Test 1 Pass: Actual value of cache expiration is matching with expected value.(%d == %d) ",
			testConf["CacheExpirationMin"], ExpOutput_CacheExpirationMin)
	}

	if testConf["RetryReq"] != ExpOutput_RetryReq {
		t.Errorf("Test 2 Fail: Actual value of retry of service invocation is not matching with expected value.(%d != %d) ",
			testConf["RetryReq"], ExpOutput_RetryReq)
	} else {
		t.Errorf("Test 2 Pass: Actual value of retry of service invocation is matching with expected value.(%d == %d) ", testConf["RetryReq"], ExpOutput_RetryReq)
	}

}

func TestGenerateKey(t *testing.T) {
	// var from language.Tag
	// var to language.Tag
	// var data string

	var pKeyTest = CatchKey{FromLanguage: "en",
		ToLanguage: "ja",
		Data:       "test"}

	pKey := GenerateKey(language.English, language.Japanese, "test")

	if pKeyTest == pKey {
		t.Error("Test 3 Pass: Expected value of generated key and test key is same ",
			pKeyTest, pKey)
	} else {
		t.Error("Test 3 Fail: Expected value of generated key and test key is not same ",
			pKeyTest, pKey)
	}

}

func TestUpdateCache(t *testing.T) {
	var pKeyTest = CatchKey{FromLanguage: "en",
		ToLanguage: "ja",
		Data:       "test"}

	var pValue = CatchValue{value: "1234456675", createdAt: time.Now().Unix()}

	pCacheMap := *CreateCatch()

	Updated := UpdateCache(&pCacheMap, pKeyTest, pValue)

	if Updated == true {
		t.Error("Test 4 Pass: Value is added/updated in catche successfully.",
			pKeyTest, pValue)
	} else {
		t.Error("Test 4 Fail: Error while handling catche.",
			pKeyTest, pValue)
	}

}

func TestGetCache(t *testing.T) {
	var pKeyTest = CatchKey{FromLanguage: "en",
		ToLanguage: "ja",
		Data:       "test"}

	var pValue = CatchValue{value: "1234456675", createdAt: time.Now().Unix()}

	pCacheMap := *CreateCatch()

	UpdateCache(&pCacheMap, pKeyTest, pValue)

	v, found := GetCache(&pCacheMap, pKeyTest)

	if found == true {
		t.Error("Test 5 Pass: Provided key is available in catche. ", pKeyTest)
	} else {
		t.Error("Test 5 Fail: Provided key is available in catche. ", pKeyTest)
	}

	if found == true {
		if pValue == v {
			t.Error("Test 6 Pass: Key value is matching with expected value. ", pValue, v)
		} else {
			t.Error("Test 6 Fail: Key is available in catch but value doesn't matching with expected value. ", pValue, v)
		}
	}
}

func TestValidateCatch(t *testing.T) {

	var pKeyTest = CatchKey{FromLanguage: "en",
		ToLanguage: "ja",
		Data:       "test"}

	var pValue = CatchValue{value: "1234456675", createdAt: time.Now().Unix()}

	pCacheMap := *CreateCatch() //Created Catch map

	UpdateCache(&pCacheMap, pKeyTest, pValue) //Added a key to map
	NoKey := len(pCacheMap.cm)
	if NoKey > 0 {
		t.Error("Test 7 Pass: Total numbers of keys available in catche. : ", NoKey)
	} else {
		t.Error("Test 7 Fail: Key is not added in catche.")
	}
	wgrp.Add(1)
	//Added sleep for 1 min to check if validator deletes catch or not
	//but it causes time out exception
	//Usually this test will always fail, because we have provided 1 minute
	//expiration period for catch key

	//time.Sleep(time.Minute)
	ValidateCatch(&pCacheMap, 1) //Called ValidateCatch function to check validity of keys
	wgrp.Wait()

	NoKey = len(pCacheMap.cm)
	if NoKey > 0 {
		t.Error("Test 8 Fail: Key is not validated. : ", NoKey)
	} else {
		t.Error("Test 8 Pass: Key is validated by validate function.", NoKey)
	}
}
