package servicetests

import (
	"net/url"
	"testing"
)

func Test_User1(t *testing.T) {
	helper := testHttpHelper.get("/user1", url.Values{})

	response := struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}{}

	helper.json(&response)

	if response.Id != 1 {
		t.Errorf("Expected id to be 1 %s", *helper.bodyString())
	}

	if response.Name != "John" {
		t.Errorf("Expected name to be John")
	}
}
