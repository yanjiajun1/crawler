package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {

		return nil,
			fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	uf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err1 := ioutil.ReadAll(uf8Reader)

	if err1 != nil {
		panic(err1)
	}
//	fmt.Printf("%s", all)
	return all, nil
}
func determineEncoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
