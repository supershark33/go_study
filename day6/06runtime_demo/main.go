package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {

	f1()

}

func f1()  {

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Celler() failed")
		return
	}
	fmt.Println("===========")
	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(path.Base(file))
	fmt.Println(line)
}
