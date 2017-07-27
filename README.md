# Kiwi
[![GoDoc](https://godoc.org/github.com/Andoryuuta/kiwi?status.svg)](https://godoc.org/github.com/Andoryuuta/kiwi)

A package for memory editing in go.

## Current Features
* Reading and Writing with support for [uint & int 8, 16, 32, 64] [float 32, 64] data types
* Support for Windows and Linux(assuming /proc/ directory exists.) 

## _Future_ plans
* Pattern scanning for bytecode
* Call remote functions via injected assembly
* Hooking functions via injected assembly
* Setting breakpoints via windows debugging api
* Mono runtime features (if hooking and remote functions are possible)

## Installation
`go get github.com/Andoryuuta/kiwi`

## Usage
```Go
package main

import (
	"log"

	"github.com/Andoryuuta/kiwi"
)

func main() {
	// The memory address of variable inside of target process.
	externVarAddr := uintptr(0x001A51E8)

	// Find the process from the executable name.
	proc, err := kiwi.GetProcessByFileName("example.exe")
	if err != nil {
		log.Fatalln("Error while trying to find process.")
	}

	// Read from the target process.
	externVar, err := proc.ReadUint32(externVarAddr)
	if err != nil {
		log.Fatalln("Error while trying to read from target process.")
	}

	// Output the variable we read.
	log.Println("Read", externVar)

	// Write a new value of 1000 to the variable
	err = proc.WriteUint32(externVarAddr, 1000)
	if err != nil {
		log.Fatal(err)
	}
}
```



