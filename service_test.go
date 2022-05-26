package main

import "testing"

func TestReadConf(t *testing.T) {
	ExpOutput_CacheExpirationMin := 30
	ExpOutput_RetryReq := 3

	var testConf = ReadConf()

	if testConf["CacheExpirationMin"] != ExpOutput_CacheExpirationMin {
		t.Errorf("Test Fail: Actual value of cache expiration is not matching with expected value.(%d != %d) ",
			testConf["CacheExpirationMin"], ExpOutput_CacheExpirationMin)
	} else {
		t.Errorf("Test Pass: Actual value of cache expiration is matching with expected value.(%d == %d) ",
			testConf["CacheExpirationMin"], ExpOutput_CacheExpirationMin)
	}

	if testConf["RetryReq"] != ExpOutput_RetryReq {
		t.Errorf("Test Fail: Actual value of retry of service invocation is not matching with expected value.(%d != %d) ",
			testConf["RetryReq"], ExpOutput_RetryReq)
	} else {
		t.Errorf("Test Pass: Actual value of retry of service invocation is matching with expected value.(%d == %d) ", testConf["RetryReq"], ExpOutput_RetryReq)
	}
}
