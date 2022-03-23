package action

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Country string `json:"country"`
	CountryCode string `json:"countryCode"`
	RegionName string `json:"regionName"`
	City string `json:"city"`
	Timezone string `json:"timezone"`
	ISP string `json:"isp"`
}

func TrackIP(ip string){
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Something went wrong bro. Try again later")
		return
	}

	body, errs := io.ReadAll(resp.Body)

	if errs != nil {
		fmt.Println("Something went wrong bro, Try again later")
		return
	}

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)

	if responseBody.Status == "fail" {
		fmt.Println("Something went wrong bro, Try again later")
		return
	}

	fmt.Println("Country:", responseBody.Country)
	fmt.Println("City:", responseBody.City)
	fmt.Println("Region", responseBody.RegionName)
	fmt.Println("Timezone:", responseBody.Timezone)
	fmt.Println("ISP:", responseBody.ISP)
}