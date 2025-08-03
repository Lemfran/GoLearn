/*
输入三角形的三边长a、b、c(边长可以是小数)，求三角形面积area，并输出。如果输入的三边构不成三角形，应给出“data error”的信息提示。注：根据“海伦－秦九韶”公式，area＝√p(p-a)(p-b)(p-c)，其中p＝(a+b+c)/2。编程可用素材：printf("Output:\ndata error\n")...

输入格式:
3 4 5

输出格式:
6.00

输入样例:
在这里给出一组输入。例如：

4 5 8
输出样例:
在这里给出相应的输出。例如：

8.18
*/
package main

import(
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scan(&a,&b,&c)
	if a+b<c||a+c<b||b+c<a {
		fmt.Printf("data error")
	} else {
		p := (a+b+c)/2
		area := math.Sqrt(p*(p-a)*(p-b)*(p-c))
		fmt.Printf("%.2f",area)
	}
}