/*
将给定的若干本图书的信息（书号、书名、定价）按输入的先后顺序加入到一个单链表中。然后遍历单链表，寻找并输出价格最高的图书信息。若存在相同的定价，则按原始顺序全部输出。

输入格式:
首先输入一个正整数T，表示测试数据的组数，然后是T组测试数据。每组测试的第一行输入正整数n，表示有n本不同的书。接下来n行分别输入一本图书的信息。其中，书号由长度等于6的纯数字构成；而书名则由长度不超过50且不含空格的字符串组成，价格包含2位小数。

输出格式:
对于每组测试，输出价格最高的图书信息（书号、书名、定价），数据之间用一个空格隔开，定价的输出保留2位小数。

输入样例:
在这里给出一组输入。例如：

1
4
023689 DataStructure 26.50
123456 FundamentalsOfDataStructure 76.00
157618 FundamentalsOfC++Language 24.10
057618 OpereationSystem 76.00
输出样例:
在这里给出相应的输出。例如：

123456 FundamentalsOfDataStructure 76.00
057618 OpereationSystem 76.00
*/

package main 

import "fmt"

type book struct{
	name string
	id string
	price float64
}

func main(){
	var n int
	fmt.Scanf("%d\n",&n)
	for i:=0;i<n;i++ {
		var m int
		fmt.Scanf("%d\n",&m)
		var max float64
		max=0
		books := make([]book,m)
		for j:=0;j<m;j++ {
			fmt.Scanf("%s %s %f\n",&books[j].id,&books[j].name,&books[j].price)
			if(books[j].price>max){
				max = books[j].price
			}
		}

		for j:=0;j<m;j++ {
			if max == books[j].price {
				fmt.Printf("%s %s %.2f\n",books[j].id,books[j].name,books[j].price)
			}
		}
	}
}