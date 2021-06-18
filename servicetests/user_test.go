package servicetests

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_User1(t *testing.T) {
	helper := testHttpHelper.get("/", url.Values{})

	response := struct {
	}{}

	p := &response
	helper.json(p)

	fmt.Println(p)
}
