package core

import (
	"errors"
	"strings"

	//    "strings"

	"net/http"

	"github.com/Testausserveri/uptimes/configuration"
)

func VerifyHost(address string, config configuration.DomainRequirements) error {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	return verifyRequest(response, config)
}

func verifyRequest(response *http.Response, config configuration.DomainRequirements) error {
	if config.Status != response.StatusCode {
		return errors.New("status code doesn't match")
	}

	if !strings.Contains(response.Header.Get("Content-Type"), config.ContentType) {
		return errors.New("content type doesn't match")
	}

	return nil
}
