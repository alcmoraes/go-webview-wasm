package main

import (
	"fmt"
	"syscall/js"
)

var (
	no int
)

func main() {
	js.Global().Set("PrintMessage", js.FuncOf(PrintMessage))
	select {}
	fmt.Println("Bye Wasm !")
}

func PrintMessage(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	no++
	fmt.Printf("Message no %d: %s\n", no, message)
	return message
}
