package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// func main() {

// 	fmt.Println("sendWithSmsCentre:=>", "+998900265560")
// 	client := http.Client{}

// 	parms := url.Values{}
// 	parms.Add("login", "mangasushi")
// 	parms.Add("psw", "H&H@Vj3qq4")
// 	parms.Add("list", "+998900265560:test")

// 	request, err := http.NewRequest("POST", "https://smsc.kz/sys/send.php", strings.NewReader(parms.Encode()))
// 	if err != nil {
// 		fmt.Println("errrr", err)
// 	}

// 	// request.Form.Add("login", "login")
// 	// request.Header.Add("Content-Type", "application/json")
// 	defer request.Body.Close()

// 	res, err := client.Do(request)
// 	fmt.Println("res", res)
// 	if err != nil {
// 		fmt.Println("error while client.Do(request):=> ", err)

// 	}
// 	byte1, _ := ioutil.ReadAll(res.Body)

// 	fmt.Println("rbyte1", string(byte1))

// 	if res.StatusCode != 200 {
// 		byte1, _ := ioutil.ReadAll(res.Body)
// 		fmt.Println("RES Body from sendWithSmsCentre:=> ", string(byte1))

// 	}

// }
