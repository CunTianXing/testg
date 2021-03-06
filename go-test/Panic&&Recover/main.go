package main

import (
	"fmt"
)

func main() {
	fmt.Println("start test")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
	//   start test
	// panic: runtime error: last name cannot be nil
	//
	// goroutine 1 [running]:
	// main.fullName(0xc42003bf40, 0x0)
	// 	/data/go/src/github.com/CunTianXing/go_app/go-test/Panic&&Recover/main.go:19 +0x166
	// main.main()
	// 	/data/go/src/github.com/CunTianXing/go_app/go-test/Panic&&Recover/main.go:10 +0x97
	// exit status 2

}

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
