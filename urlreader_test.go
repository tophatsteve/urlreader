package urlreader

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test Data")
	}))
	defer ts.Close()

	r := NewReader(ts.URL)
	result, err := ioutil.ReadAll(r)

	assert.Equal(t, nil, err, "err should be nil")
	assert.Equal(t, "Test Data\n", fmt.Sprintf("%s", result), "returned data should be 'Test Data\n")
}
