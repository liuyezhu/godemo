package infra

import (
	"io/ioutil"
	"net/http"
)

//空接口：可以承载任何类型的变量，也可以说任何类型都实现(满足)了空接口
type Retriever struct{}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
