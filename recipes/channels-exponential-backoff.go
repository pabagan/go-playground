package recipes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var urls = []string{
	"http://pulsoconf.co/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
	"https://dev.klikkie.com/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		default:
			fmt.Printf(".")
			time.Sleep(5e7)
		}
	}
	return responses
}

type ResponseData struct {
	Project string `json:"project"`
}

type ErrorResponse struct {
	Message string `json:"error_message"`
}

func GetAccount(url string) (*ResponseData, *ErrorResponse) {
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(body)
	// if res.StatusCode > 204 {
	// 	error := new(ErrorResponse)
	// 	errJSON := json.Unmarshal(body, &error)

	// 	if errJSON != nil {
	// 		panic(errJSON)
	// 	}
	// 	return nil, error
	// }

	account := new(ResponseData)
	// errJSON := json.Unmarshal(body, &account)

	fmt.Println(account)

	// if errJSON != nil {
	// 	panic(errJSON)
	// }

	return account, nil
}

func RequestBackoff(url string) *HttpResponse {
	ch := make(chan *HttpResponse)

	go func(url string) {
		fmt.Printf("Fetching %s \n", url)
		resp, err := http.Get(url)
		ch <- &HttpResponse{url, resp, err}
	}(url)

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			fmt.Printf("%s was fetched\n", r.response)
			// response = r
			return r
		default:
			fmt.Printf(".")
			time.Sleep(5e7)
		}
	}
}

func RequestExponentialBackofff() {
	results := asyncHttpGets(urls)
	for _, result := range results {
		fmt.Printf("%s status: %s\n", result.url, result.response.Status)
	}
}
