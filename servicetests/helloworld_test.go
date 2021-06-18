package servicetests

import (
	"net/url"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	helper := testHttpHelper.get("/", url.Values{})
	value, err := helper.bodyString()

	if err != nil {
		t.Error(err)
	}

	expected := "Hi there, you are visiting !"
	if value != expected {
		t.Errorf("expected '%s', got '%s'", expected, value)
	}
}
