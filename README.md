# Kiwi
A package for memory editing in go.


## Installation
`go get github.com/Andoryuuta/kiwi`

## Usage
```
package main

import(
	"github.com/Andoryuuta/kiwi"
	"fmt"
)

func main(){
	//The memory address of variable inside of target process.
	externVarAddr := uintptr(0x001A51E8)

	//Find the process from the executable name.
	proc, err := kiwi.GetProcessByFileName("example.exe")
	if err != nil{
		fmt.Println("Error while trying to find process.")
		panic(err)
	}

	//Read from the target process
	externVar, err := proc.ReadUint32(externVarAddr)
	if err != nil{
		fmt.Println("Error while trying to read from target process.")
		panic(err)
	}
	
	//Output the variable we read.
	fmt.Println("The value is", externVar)
	
	//Write a new value of 177357172 to the variable
	err = proc.WriteUint32(externVarAddr, 177357172)
	if err != nil{
		panic(err)
	}
}
```
## Current Features
* Finding process given executable file name
* Reading and Writing with support for [uint & int 8, 16, 32, 64] [float 32, 64] data types
* Support for Windows and Linux(assuming /proc/ directory exists.) 

## _Future_ plans
* Pattern scanning for bytecode
* Call remote functions via injected assembly
* Hooking functions via injected assembly
* Setting breakpoints via windows debugging api
* Mono runtime features (if hooking and remote functions are possible)


