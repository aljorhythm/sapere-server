package servicetests

import (
	"net/url"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	helper := testHttpHelper.get("/", url.Values{})

	value := helper.bodyString()

	expected := "Hi there, you are visiting !"

	if *value != expected {
		t.Errorf("expected '%s', got '%s'", expected, *value)
	}
}
