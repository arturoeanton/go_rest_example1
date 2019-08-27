package main

// GOARCH=wasm GOOS=js go build -o  lib.wasm .
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"syscall/js"
)

var (
	document Document
	done     chan struct{}
)

func makeRequest() {

	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result)
	fmt.Println(result["data"])
}

func test(this js.Value, i []js.Value) interface{} {
	if len(i) > 1 {
		return js.ValueOf(i[0].Int() + i[1].Int())
	}
	if len(i) == 1 && i[0].Type() == js.TypeString {
		return document.GetElementByID(i[0].String())
	}
	go makeRequest()
	fmt.Println("Get......")
	return nil
}

func registerCallbacks() {
	js.Global().Set("test", js.FuncOf(test))
}

func main() {
	document = Document{}
	done = make(chan struct{})
	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-done
}
