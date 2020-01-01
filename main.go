package main

import (
	"fmt"

	"github.com/Kourin1996/re-parser-in-go/compiler"
	"github.com/Kourin1996/re-parser-in-go/re"
	"github.com/Kourin1996/re-parser-in-go/vm"
)

func main() {
	fmt.Println("Hello, World")
	/*
		codes := []operator.Operator{
			&operator.OChar{'a'},
			&operator.OSplit{0, 2},
			&operator.OChar{'b'},
			&operator.OMatch{},
		}
		vm := vm.NewVM(codes, "aab")
		vm.RunCode()
	*/
	r := &re.REGroup{"1", &re.REZeroOrMany{&re.REConcat{&re.REChar{'a'}, &re.REChar{'b'}}}}
	fmt.Printf("[RE] %#v\n", r)
	codes := compiler.Compile(r)
	fmt.Printf("Codes %#v\n", codes)
	for i, code := range codes {
		fmt.Printf("%d %s\n", i, code.String())
	}

	vm := vm.NewVM(codes, "abab")
	vm.RunCode()

}
