package main

import (
	"context"
	"errors"
	"time"

	"golang.org/x/text/language"
)

var gConf Conf
var gCacheMap CacheMap

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
	gConf = ReadConf()
	gCacheMap = *CreateCatch(int64(gConf.CacheExpirationMin))

	return &Service{
		translatorClient: t,
	}
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	var pKey = &CatchKey{FromLanguage: from,
		ToLanguage: to,
		Data:       data}

	v, found := GetCache(&gCacheMap, pKey)

	if found {
		return v.value, errors.New("")
	} else {
		//Commented original statement
		//return s.translatorClient.Translate(ctx, from, to, data)

		str, err := s.translatorClient.Translate(ctx, from, to, data)
		var pValue = &CatchValue{value: str, createdAt: time.Now().Unix()}
		UpdateCache(&gCacheMap, pKey, *pValue)
		return str, err
	}

}
