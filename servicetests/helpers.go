package servicetests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var testHttpHelper = newHelper()

type HttpHelper struct {
	host string
}

type HttpResponseHelper struct {
	response *http.Response
	bytes    *[]byte
}

func (h *HttpResponseHelper) bodyString() *string {
	bodyBytes, err := h.bodyBytes()
	if err != nil {
		return nil
	}
	str := string(bodyBytes)
	return &str
}

func (h HttpResponseHelper) status() int {
	return h.response.StatusCode
}

func (h *HttpResponseHelper) bodyBytes() ([]byte, error) {
	if h.bytes != nil {
		return *h.bytes, nil
	}

	bodyBytes, err := ioutil.ReadAll(h.response.Body)
	if err != nil {
		return []byte{}, err
	}

	h.bytes = &bodyBytes
	return *h.bytes, err
}

func (h *HttpResponseHelper) json(obj interface{}) error {
	bytes, err := h.bodyBytes()

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, obj)

	return err
}

func (h HttpHelper) get(path string, params url.Values) HttpResponseHelper {
	base, err := url.Parse(h.host)
	if err != nil {
		panic(err)
	}

	// Path params
	base.Path += path

	// Query params
	base.RawQuery = params.Encode()

	url := base.String()
	log.Printf("Get: %s", url)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	return HttpResponseHelper{
		response: resp,
	}
}

func newHelper() HttpHelper {
	host := os.Getenv("SERVICE_HOST")
	host = strings.Trim(host, "/")
	if host == "" {
		panic("host should not be empty")
	}
	return HttpHelper{host: host}
}
