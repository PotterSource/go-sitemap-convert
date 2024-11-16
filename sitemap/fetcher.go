package sitemap

import (
	"io"
	"net/http"
)

func Fetch(url string) ([]byte, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}

	return data, res.Header.Get("Content-Type"), nil
}
