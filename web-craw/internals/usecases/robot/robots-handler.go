package robots

import (
	"net/http"

	"github.com/temoto/robotstxt"
)

// IsAllowed checks if the URL can be crawled based on robots.txt.
func IsAllowed(baseURL, userAgent, urlPath string) (bool, error) {
	resp, err := http.Get(baseURL + "/robots.txt")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		return false, err
	}

	group := data.FindGroup(userAgent)
	return group.Test(urlPath), nil
}
