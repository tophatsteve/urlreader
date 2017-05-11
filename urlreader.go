// The ulrreader package implements the io.Reader interface for HTTP resources.
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

// NewReader create a new instance of urlreader, passing in the url to be read. The returned urlreader can than be treated like any other io.Reader.
func NewReader(url string) *Reader {
	return &Reader{url, false}
}

// Read reads data into p and returns the number of bytes read into p. The entire resource is read at once, so len(p) will
// be the length of the resource. After one read, the end of the resource will have been reached so any subsequent reads
// will return a count of zero and an error of type io.EOF.
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
