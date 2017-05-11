package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type GeoIP struct {
    Ip          string  `json:"ip"`
    CountryCode string  `json:"country_code"`
    CountryName string  `json:"country_name"`
    RegionCode  string  `json:"region_code"`
    RegionName  string  `json:"region_name"`
    City        string  `json:"city"`
    Zipcode     string  `json:"zipcode"`
    Lat         float32 `json:"latitude"`
    Lon         float32 `json:"longitude"`
    MetroCode   int     `json:"metro_code"`
    AreaCode    int     `json:"area_code"`
}

var (
    address   string
    err       error
    geo     GeoIP
    response  *http.Response
    body      []byte
)

func main(){
    address = "www.sensetime.com"
    response, err = http.Get("https://freegeoip.net/json/" + address)
    if err != nil {
        fmt.Println(err)
    }
    defer response.Body.Close()

    body, err = ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
    }
    err = json.Unmarshal(body,&geo)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("\n==== IP Geolocation Info ====\n")
	fmt.Println("IP address:\t", geo.Ip)
	fmt.Println("Country Code:\t", geo.CountryCode)
	fmt.Println("Country Name:\t", geo.CountryName)
	fmt.Println("Zip Code:\t", geo.Zipcode)
	fmt.Println("Latitude:\t", geo.Lat)
	fmt.Println("Longitude:\t", geo.Lon)
	fmt.Println("Metro Code:\t", geo.MetroCode)
	fmt.Println("Area Code:\t", geo.AreaCode)
}