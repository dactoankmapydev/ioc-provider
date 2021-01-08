package ioc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	errUnexpectedResponse = "unexpected response: %s"
)

type HTTPClient struct{}

var (
	httpClient = HTTPClient{}
)

func (c HTTPClient) getVirustotal(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("X-Apikey", os.Getenv("VIRUSTOTAL_API_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", api, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, resp.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c HTTPClient) getOtx(api string) ([]byte, error) {
	req, _ := http.NewRequest("GET", api, nil)
	req.Header.Set("X-OTX-API-KEY", os.Getenv("OTX_API_KEY"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	c.info(fmt.Sprintf("GET %s -> %d", api, resp.StatusCode))
	if resp.StatusCode != 200 {
		respErr := fmt.Errorf(errUnexpectedResponse, resp.Status)
		fmt.Sprintf("request failed: %v", respErr)
		return nil, respErr
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c HTTPClient) info(msg string)  {
	log.Printf("[JSONClient] %s\n", msg)
}
