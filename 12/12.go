/*
将字符串反转

输入格式:
输入一行字符串，不包含空格

输出格式:
输出该字符串的反转

输入样例:
在这里给出一组输入。例如：

abc
输出样例:
在这里给出相应的输出。例如：

cba
*/
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