package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	// log to brower console
	fmt.Println("hellowasm console - Hello World!")

	// open alert window
	alert := js.Global().Get("alert")
	alert.Invoke("hellowasm console - Hello Alert!")

	// call javascript code
	js.Global().Get("console").Call("log", "Hello world Go/wasm!")
	js.Global().Get("document").Call("getElementById", "timestamp").Set("innerText", time.Now().String())
}
