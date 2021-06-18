package servicetests

import (
	"net/url"
	"testing"
)

type PingResponse struct {
	Up bool `json:"up"`
}

func Test_Ping(t *testing.T) {
	helper := testHttpHelper.get("/ping", url.Values{})

	response := PingResponse{}

	p := &response
	helper.json(p)

	expected := PingResponse{
		Up: true,
	}

	if response != expected {
		t.Errorf("Got %#v, wanted %#v", response, expected)
	}
}
