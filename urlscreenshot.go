package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	//GET
	apiFlashEndpoint := "https://api.apiflash.com/v1/urltoimage"
	request, _ := http.NewRequest("GET", apiFlashEndpoint, nil)
	query := request.URL.Query()
	query.Add("access_key", "514f1487de894b47bb90616b63f40059")

	//Receive the URL
	var url string
	fmt.Println("Enter the url: ")
	fmt.Scanln(&url)
	query.Add("url", url)
	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, _ := client.Do(request)
	defer response.Body.Close()

	image, _ := os.Create("screenshot.jpeg")
	io.Copy(image, response.Body)
	image.Close()
}
