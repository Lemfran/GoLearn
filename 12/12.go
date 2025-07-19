package main 

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)
	a:=s[:]
	length:=len(a)
	for i:=0;i<length;i++ {
		fmt.Printf("%c",a[length-i-1])
	}
	fmt.Printf("\n")
}