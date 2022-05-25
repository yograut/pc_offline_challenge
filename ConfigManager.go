package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	CacheExpirationMin int
	RetryReq           int
}

func ReadConf() Conf {
	ConfFile, _ := os.Open("conf.json")
	jsonDecoder := json.NewDecoder(ConfFile)
	configuration := Conf{}
	err := jsonDecoder.Decode(&configuration)

	if err == nil {
		fmt.Println("Error in config", err)
	}
	ConfFile.Close()
	return configuration
}
