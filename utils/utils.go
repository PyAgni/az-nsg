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
	Ip string
}

const IpApiUrl = "https://api.ipify.org/?format=json"

func GetPublicIP() string {
	res, err := http.Get(IpApiUrl)
	if err != nil {
		fmt.Printf("Couldn't get your public IP. Error: %s\n", err)
		return ""
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
		return ""
	}
	var publicIP IP
	err = json.Unmarshal(body, &publicIP)
	if err != nil {
		fmt.Printf("Couldn't unmarshall response body. Error: %s\n", err)
		return ""
	}

	return publicIP.Ip
}
