package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// log to brower console
	fmt.Println("hellowasm console - Hello World!")

	// open alert window
	alert := js.Global().Get("alert")
	alert.Invoke("hellowasm console - Hello Alert!")
}
