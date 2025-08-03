/*
已知素数序列为2、3、5、7、11、13、17、19、23、29……，即素数的第一个是2，第二个是3，第三个是5……那么，随便挑一个数，若是素数，能确定是第几个素数吗？如果不是素数，则输出0。

输入格式:
测试数据有多组，处理到文件尾。每组测试输入一正整数N（1≤N≤1000000）。

输出格式:
对于每组测试，输出占一行，如果输入的正整数是素数，则输出其排位，否则输出0。

输入样例:
在这里给出一组输入。例如：

6
2
6
4
5
13
991703
输出样例:
在这里给出相应的输出。例如：

1
0
0
3
6
77901
*/

package main

import(
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
	fmt.Scanf("%d\n",&n)
	arr :=make([]int,n)

	for i:=0;i<n;i++ {
		fmt.Scanf("%d\n",&arr[i])
	}

	for _,num :=range arr {
		if issushu(num) {
			count:=1
			for i:=3;i<=num;i++{
				if issushu(i){
					count++
				}
			}
			fmt.Printf("%d\n",count)
		} else {
			fmt.Printf("0\n")
		}
	}
}