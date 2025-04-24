package main

import (
	"bufio
	"fmt"
	"os"
)

func main () {
reader := bufio.NewReader(os.Stdin)

fmt.print("Enter your full name")
name, _ := reader.ReadString ('/n') // includes newline character
fmt.printin("Your name is ", name)
/ / fmt.printin(reflect.Type of (name))
fmt.printin ("Enter your faviourte number")
favNum, _ := reader.ReadString('/n')
fmt.printin ("Your name is ",")
fmt.print ("Enter your fun fact")
funfact, _ := reader.ReadString ('/n')
fmt.printin ("Your fun fact is ", funfact)