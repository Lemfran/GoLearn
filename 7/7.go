/*
请编写程序，输入 n (n≤20)，计算并输出从 0 到 n 的阶乘之和。

0!+1!+2!+3+⋯+n!

输入格式:
n

输出格式:
阶乘之和

输入样例:
在这里给出一组输入。例如：

4
输出样例:
在这里给出相应的输出。例如：

34
*/
package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	a :=make([]rune,n)
	b :=make([]rune,n)

	for i:=0;i<n;i++{
		fmt.Scanf("%c %c\n",&a[i],&b[i])
	}

	for i:=0;i<n;i++{
		if a[i]==b[i] {
			fmt.Printf("TIE\n")
		} else if (a[i]=='R'&&b[i]=='S')||(a[i]=='S'&&b[i]=='P')||(a[i]=='P'&&b[i]=='R'){
			fmt.Printf("Player1\n")
		} else {
			fmt.Printf("Player2\n")
		}
	}
}