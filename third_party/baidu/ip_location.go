package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"test_go_mod/utils/helper"
	"time"
)

type IpLocationData struct {
	Address string `json:"address"`
	Content struct {
		Address       string `json:"address"`
		AddressDetail struct {
			Adcode       string `json:"adcode"`
			City         string `json:"city"`
			CityCode     int    `json:"city_code"`
			District     string `json:"district"`
			Province     string `json:"province"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_detail"`
	} `json:"content"`
}

const url = "https://api.map.baidu.com/location/ip?"

var client = http.Client{
	Timeout: 10 * time.Second,
}

//获取Ak
func getAk() string {
	aks := [...]string{"cs6hgmXjKyokaq1WCN7dnpXVqKnGf1KK", "TEV04btYOdHr9BK58ClDNV7s8R9fx1VS", "fxZ83jhePibGTNoAyP8NIqkGSQjhUM9O"}
	rand.Seed(time.Now().UnixNano())
	//注意  这句一定要加~！！！！！！获取随机数，不加随机种子，每次遍历获取都是重复的一些随机数据
	key := rand.Intn(3)
	ak := aks[key]
	return ak
}

func IpLocation(ip string) string {
	ak := getAk()
	param := make(map[string]string)
	param["ak"] = ak
	param["ip"] = ip
	param["coor"] = "bd09l"
	url := url + helper.HttpBuiltQuery(param)
	fmt.Println(url)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, _ := client.Get(url)
	result, _ := ioutil.ReadAll(resp.Body)
	var data IpLocationData
	fmt.Println(string(result))
	e := json.Unmarshal([]byte(string(result)), &data)
	if e != nil {
		panic("error occurred")
	}
	fmt.Println(data.Content.AddressDetail.City)
	return "ss"
}

func HttpGetRequest(url string, result interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	fmt.Println(decoder)

	err = decoder.Decode(&result)

	return err
}

func main() {
	getAk()
	IpLocation("116.21.14.124")
}
