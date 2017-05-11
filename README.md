# urlreader

The ulrreader package implements the io.Reader interface for HTTP resources.

``` shell
go get -u github.com/tophatsteve/urlreader
```

Create a new instance of urlreader by calling the NewReader function passing in the url to be read. The returned urlreader can than be treated like any other io.Reader, for example:

``` go
r := NewReader("http://example.com/example.txt")
result, _ := ioutil.ReadAll(r)
fmt.Println(result)
```

