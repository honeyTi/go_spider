package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	transform "golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error:status code", resp.StatusCode)
		return
	} else {
		e := determineEncoding(resp.Body)
		// htmlcharset方式判断解析
		utf8Reader :=transform.NewReader(resp.Body, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", all)
		printCityList(all)
	}
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(content []byte)  {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(content, -1)
	for _, m := range matches {
		fmt.Printf("City: %s; Url: %s \n", m[2], m[1])
	}
	fmt.Printf("matches found: %d\n",len(matches))

}
