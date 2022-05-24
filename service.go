package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"golang.org/x/text/language"
)

// Service is a Translator user.
type Service struct {
	translatorClient TranslatorAPI
}

type Conf struct {
	CacheExpirationMin int
	RetryReq           int
}

func NewService() *Service {
	t := newTranslatorStub(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)

	return &Service{
		translatorClient: t,
	}
}

func ReadConf() Conf {
	ConfFile, _ := os.Open("conf.json")
	//defer file.Close()
	jsonDecoder := json.NewDecoder(ConfFile)
	configuration := Conf{}
	err := jsonDecoder.Decode(&configuration)

	if err == nil {
		fmt.Println("Error in config", err)
	}
	return configuration
}

func (s *Service) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {

	// ConfigValue := ReadConf()
	// if ConfigValue != nil {

	// }
	return s.translatorClient.Translate(ctx, from, to, data)
}
