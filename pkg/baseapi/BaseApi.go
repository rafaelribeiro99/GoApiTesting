package baseapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func CheckHealth(client *resty.Client, url string) (*resty.Response, error) {

	resp, err := client.R().
		Get(url)

	if err != nil {
		fmt.Println("Error on PageObject function CheckHealth:", err)
		return resp, err
	}

	return resp, err
}

func GetList(client *resty.Client, headers, queryParams map[string]string, url string) (*resty.Response, error) {

	resp, err := client.R().
		SetHeaders(headers).
		SetQueryParams(queryParams).
		Get(url)

	if err != nil {
		fmt.Println("Error on PageObject function GetList:", err)
		return resp, err
	}

	return resp, err
}

func GetByID(client *resty.Client, headers, pathParams map[string]string, url string) (*resty.Response, error) {

	resp, err := client.R().
		SetHeaders(headers).
		SetPathParams(pathParams).
		Get(url)

	if err != nil {
		fmt.Println("Error on PageObject function GetByID:", err)
		return resp, err
	}

	return resp, err
}

func DeleteRequest(client *resty.Client, headers, pathParams map[string]string, url string) (*resty.Response, error) {

	resp, err := client.R().
		SetHeaders(headers).
		SetPathParams(pathParams).
		Delete(url)

	if err != nil {
		fmt.Println("Error on PageObject function DeleteRequest:", err)
		return resp, err
	}

	return resp, err
}

func PostRequest(client *resty.Client, headers map[string]string, body string, url string) (*resty.Response, error) {
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		Post(url)

	if err != nil {
		fmt.Println("Error on PageObject function PostRequest:", err)
		return resp, err
	}
	return resp, err
}

func PutRequest(client *resty.Client, headers, pathParams map[string]string, body string, url string) (*resty.Response, error) {
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		SetPathParams(pathParams).
		Put(url)

	if err != nil {
		fmt.Println("Error on PageObject function PutRequest:", err)
		return resp, err
	}
	return resp, err
}

func PatchRequest(client *resty.Client, headers, pathParams map[string]string, body string, url string) (*resty.Response, error) {
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(body).
		SetPathParams(pathParams).
		Patch(url)

	if err != nil {
		fmt.Println("Error on PageObject function PatchRequest:", err)
		return resp, err
	}
	return resp, err
}
