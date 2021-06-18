package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestHttpContext(t *testing.T)  {

	req, err := http.NewRequest("GET", "http://127.0.0.1:8090/hello", nil)
	if err != nil {
		t.Fatal(fmt.Errorf("http.NewRequest Error: %s", err.Error()))
	}
	timeout, _ := context.WithTimeout(context.Background(), time.Second*3)
	req = req.WithContext(timeout)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
