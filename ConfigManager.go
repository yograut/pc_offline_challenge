package main

// Used this module to capture config values
// As per requirement we have to configure Expiration time of catch and retry count
// for api invoking if it fails

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConf() map[string]int {
	//We have used json file to handle configuration as it is easy to handle and human readable

	//Initialized configuration map with default values
	confMap := map[string]int{"CacheExpirationMin": 1, "RetryReq": 1}

	ConfFile, err := os.Open("conf.json")

	//if json file opened successfully then go further else return initialized map object
	if err == nil {

		//Reading json file
		jsonStr, _ := ioutil.ReadAll(ConfFile)

		err1 := json.Unmarshal([]byte(jsonStr), &confMap)

		//Closing of file
		defer ConfFile.Close()

		//if error is there then we are sending object with 0 values
		if err1 != nil {
			fmt.Println("Error while reading config file, sending object with 0 values", err)
		}
	}

	return confMap
}
