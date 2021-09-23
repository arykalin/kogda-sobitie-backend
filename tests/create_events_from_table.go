package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/event"
	method := "POST"

	payload := strings.NewReader(`{
    "date":"24.09.2021",
    "title":"Чтения на среднем",
    "duration":"3-5 часов",
    "link":"https://vk.com/4tenia",
    "who_manages":"Татка",
    "for_whom":"для всех",
    "where":"Дом Средний",
    "description":"четния по пятницам на среднем",
    "wanting_people":"15"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
