package urlreader

import (
	"fmt"
	"io/ioutil"
)

func ExampleNewReader() {
	r := NewReader("http://example.com/example.txt")
	result, _ := ioutil.ReadAll(r)
	fmt.Println(result)
}
