package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Visitor struct {
	Client   string
	IP       string
	Location GeoIP
}

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

func main() {
	lambda.Start(GetVisitoryLambdaHandler)

}

func GetVisitorInfoHttpHandler(w http.ResponseWriter, req *http.Request) {
	visitor := Visitor{
		IP:     req.RemoteAddr,
		Client: req.Header.Get("User-Agent"),
	}

	err := visitor.Location.GeoLocateIPHandler(visitor.IP)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
	fmt.Fprintf(w, "%v\n", visitor)
}

func GetVisitoryLambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	v := Visitor{
		IP:     req.Headers["client-ip"],
		Client: req.Headers["client-user-agent"],
	}

	err := v.Location.GeoLocateIPHandler(v.IP)
	if err != nil {
		return createAPIGatewayProxyResponseError(err), err
	}
	visitor, err := json.Marshal(v)
	if err != nil {
		return createAPIGatewayProxyResponseError(err), err
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(visitor),
	}, nil
}

func createAPIGatewayProxyResponseError(err error) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: 503,
		Body:       "Something went wrong: " + err.Error(),
	}
}

func (g *GeoIP) GeoLocateIPHandler(address string) (err error) {
	accessKey := os.Getenv("GEOIP_ACCESS_KEY")
	if strings.HasPrefix(address, "127.0.0.1") {
		address = "github.com"
	}
	log.Println("address: ", address)
	resp, err := http.Get("https://api.ipstack.com/" + address + "?access_key=" + accessKey)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &g)
	if err != nil {
		fmt.Println("Failed to unmarshel JSON")
		return err
	}

	return nil
}
