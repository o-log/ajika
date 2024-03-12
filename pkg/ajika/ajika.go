package ajika

import (
	"ajika/pkg/wraperr"
	"bytes"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"time"
)

// struct used to create service instance for injections
type Ajika struct {
	AllowedDomains []string // protect against SSRF
}

func (ajika Ajika) Request(salt Salt, timeoutSec int, method string, sourceUrl string, requestBody string, headers map[string]string) (string, error) {
	urlParts, err := url.Parse(sourceUrl)
	if err != nil {
		return "", wraperr.Wrap(err)
	}
	if !slices.Contains(ajika.AllowedDomains, urlParts.Hostname()) {
		return "", wraperr.Err("AllowedDomains does not contain " + urlParts.Hostname())
	}

	request, err := http.NewRequest(method, sourceUrl, bytes.NewReader([]byte(requestBody)))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	for header, value := range headers {
		request.Header.Set(header, value)
	}

	salt.Debugf("requesting %v, body: %v", sourceUrl, requestBody)

	client := http.Client{Timeout: time.Duration(timeoutSec) * time.Second}
	result, err := client.Do(request)
	if err != nil {
		return "", err
	}
	salt.Debugf("response status code: %v", result.StatusCode)

	// close the body to free resources
	defer result.Body.Close()

	responseBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	salt.Debugf("response: %v", string(responseBytes))

	// check http code after reading body, because body may contain error details
	successHttpCodes := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent,
	}
	if !slices.Contains(successHttpCodes, result.StatusCode) {
		return string(responseBytes), wraperr.Err("response code not 200/201/204: " + strconv.Itoa(result.StatusCode))
	}

	return string(responseBytes), nil
}
