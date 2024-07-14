package something

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Wget() bool {
	resp, err := http.Get("http://httpbin.org/anything?hello=world")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("body = %s\n\n", all)
	return true
}
