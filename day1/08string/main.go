package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "刘云克"
	desc := "大帅比"

	ss := name + desc
	fmt.Print(ss + "\n")

	bb := fmt.Sprintf("%s%s", name, desc)

	fmt.Println(bb)

	sps := strings.Split(name, "云")
	fmt.Println(sps)

	join := strings.Join(sps, "+")
	fmt.Println(join)

	contains := strings.Contains(ss, name)
	fmt.Println(contains)

	//前缀
	fmt.Println(strings.HasPrefix(ss, name))
	//后缀
	fmt.Println(strings.HasSuffix(ss, desc))

	fmt.Println(strings.Index(ss, desc))

}
