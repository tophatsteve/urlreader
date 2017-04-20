package urlreader

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Reader struct {
	url string
	eof bool
}

func NewReader(url string) *Reader {
	return &Reader{url, false}
}

func (r *Reader) Read(p []byte) (int, error) {
	if r.eof == true {
		return 0, io.EOF
	}

	resp, err := http.Get(r.url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	r.eof = true
	copy(p, content)
	return len(content), nil
}
