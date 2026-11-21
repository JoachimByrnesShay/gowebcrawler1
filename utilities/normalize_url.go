package utilities

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(url_string string) (string, error) {
	url, err := url.Parse(strings.ToLower(url_string))
	fmt.Println("its here: ", err)
	if err != nil {
		return "", fmt.Errorf(strings.ToLower("could not parse url: %w"), err)
	}
	host := url.Host
	path := strings.TrimSuffix(url.Path, "/")

	return host + path, nil
}
