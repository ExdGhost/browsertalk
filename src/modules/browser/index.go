package browser

import (
	"errors"
	"strings"
)

var chrome *Chrome
var safari *Safari

// Start ...
func Start(query *Query) error {

	if strings.ToLower(query.Browser) == "chrome" {
		chrome.Start()
	} else if strings.ToLower(query.Browser) == "safari" {
		safari.Start()
	} else {
		return errors.New("Invalid browser")
	}

	return nil
}
