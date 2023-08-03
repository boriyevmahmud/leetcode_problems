package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	doHttp()

}

func doHttp() {
	url := "https://io.bellissimo.uz/api/verify"

	r, _ := json.Marshal("")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(r))
	if err != nil {
		fmt.Println("err", err)
	}

	c := http.Client{}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("err1", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err2", err)
	}
	fmt.Println("body", string(body))
}
