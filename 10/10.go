/*
输入一个数字n(n是正整数)，将其分解为质因数的乘积。返回一个由质因数组成的切片

输入格式:
输入一个正整数

输出格式:
质因数组成的切片

输入样例:
在这里给出一组输入。例如：

24
输出样例:
在这里给出相应的输出。例如：

[2 2 2 3]
*/

package main

import (
	"fmt"
	"math"
)


func issushu(n int) bool{
	if n<=1 {
		return false
	}
	for i:=2;i<int(math.Sqrt(float64(n)));i++ {
		if n%i==0{
			return false
		}
	}
	return true
}

func main(){
	var n int
	fmt.Scanf("%d",&n)
	var a []int
	for {
		if n==1 {
			break
		}
		for i:=2;i<=n&&issushu(i);i++ {
			if n%i==0 {
				a=append(a, i)
				n=n/i
				break
			}
		}
	}
	fmt.Println(a)
}