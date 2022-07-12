package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//add public ip logic

// IP Struct for parsing response of public IP from https://api.ipify.org/?format=json
type IP struct {
	ip string
}

const IpApiUrl = "https://api.ipify.org/?format=json"

func GetPublicIP() IP {
	res, err := http.Get(IpApiUrl)
	if err != nil {
		fmt.Printf("Couldn't get your public IP. Error: %s\n", err)
		return IP{}
	}
	//Close the response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Couldn't get your public IP. Error: %s\n", err)
		}
	}(res.Body)

	//Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Couldn't read IP from response. Error: %s\n", err)
		return IP{}
	}
	var ip IP
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return IP{}
	}

	return ip
}
