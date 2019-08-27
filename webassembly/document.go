package main

import (
	"syscall/js"
)

// Document is struct for representation document
type Document struct {
}

// GetElementByID is call document.getElementById
func(document Document) GetElementByID(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

// GetElementsByClassName is call document.getElementsByClassName
func(document Document) GetElementsByClassName(class string) js.Value {
	return js.Global().Get("document").Call("getElementsByClassName", class)
}

// GetElementsByName is call document.getElementsByName
func(document Document) GetElementsByName(name string) js.Value {
	return js.Global().Get("document").Call("getElementsByName", name)
}

// getElementsByTagName is call document.getElementsByTagName
func(document Document) getElementsByTagName(name string) js.Value {
	return js.Global().Get("document").Call("getElementsByTagName", name)
}

// GetElementsByTagNameNS is call document.getElementsByTagNameNS
func(document Document) GetElementsByTagNameNS(name string) js.Value {
	return js.Global().Get("document").Call("getElementsByTagNameNS", name)
}

// GetSelection is call document.getSelection
func(document Document) GetSelection(name string) js.Value {
	return js.Global().Get("document").Call("getSelection", name)
}

func(document Document) test(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int() + i[1].Int())
}