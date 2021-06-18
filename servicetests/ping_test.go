package servicetests

import (
	"net/url"
	"testing"
)

type PingResponse struct {
	Up    bool `json:"up"`
	Build struct {
		Commit string `json:"commit"`
		Branch string `json:"branch"`
		CiPlat string `json:"ci_plat"`
	} `json:"build"`
}

func Test_Ping(t *testing.T) {
	helper := testHttpHelper.get("/ping", url.Values{})

	t.Logf("ping response: %s", *helper.bodyString())

	response := PingResponse{}

	p := &response
	helper.json(p)

	expected := PingResponse{
		Up: true,
	}

	if response.Up != true {
		t.Errorf("Service is not up: Got %#v, wanted %#v", response, expected)
	}

	if response.Build.Commit == "" {
		t.Errorf("Commit is missing: Got %#v, wanted %#v", response, expected)
	}

	if response.Build.CiPlat == "" {
		t.Errorf("CI Platform is missing: Got %#v, wanted %#v", response, expected)
	}

	if response.Build.Branch == "" {
		t.Errorf("Branch is missing: Got %#v, wanted %#v", response, expected)
	}
}
