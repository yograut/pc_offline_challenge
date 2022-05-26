package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/text/language"
)

//Used global variable to store config values and catched map structure
var gConfMap = map[string]int{}

//var gConfMap
var gCacheMap CacheMap

//x := map[string]int{}
// Service is a Translator user.
type Service struct {
	translatorClient TranslatorAPI
}

func NewService() *Service {
	t := newTranslatorStub(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	//Initialize global variables
	gConfMap = ReadConf()
	gCacheMap = CreateCatch(int64(gConfMap["CacheExpirationMin"]))

	return &Service{
		translatorClient: t,
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	//Initialize key object with provided values
	var p1 string = fmt.Sprintf("%v", from)
	var p2 string = fmt.Sprintf("%v", to)

	var pKey = CatchKey{FromLanguage: p1,
		ToLanguage: p2,
		Data:       data}

	//Checked if provided key is already available in catche
	v, found := GetCache(&gCacheMap, pKey)

	//if found then send the values to caller function with values from catch
	//We are not invoking translation service
	if found {
		return v.value, errors.New("")
	} else {
		//Commented original statement
		//return s.translatorClient.Translate(ctx, from, to, data)

		//If key is not available in catch then invoke service
		strValue, err := "", errors.New("")

		//we are invoking service for maximum time which is mentioned in config if error occures
		for i := 0; i < int(gConfMap["RetryReq"]); i++ {
			strValue, err = s.translatorClient.Translate(ctx, from, to, data)
			if err == nil {
				//Service returned new value successfully
				var pValue = CatchValue{value: strValue, createdAt: time.Now().Unix()}

				//We have updated catche with new values
				UpdateCache(&gCacheMap, pKey, pValue)
				break
			}
		}
		return strValue, err
	}
}
