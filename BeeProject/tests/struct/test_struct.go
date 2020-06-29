package main

import "fmt"

type student struct {
	id int
	name  string
}

func main(){
	var s = student{
		id: 1,
		name: "aaa",
	}

	fmt.Println(s)
}
