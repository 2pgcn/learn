package main

import "fmt"

func main()  {
	s := "rat"
	t := "car"
	var c1, c2 [26]int

	for _,ch := range s{
		c1[ch-'a']++
	}
	for _,ch := range t{
		c2[ch-'a']++
	}
	fmt.Println(c1,c2)
}
