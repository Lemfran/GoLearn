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