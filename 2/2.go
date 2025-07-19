package main

import "fmt"

type account struct{
	name string
	balance float64
}

func (acc *account)Deposit (amount float64) {
	defer fmt.Printf("存款：%s + %.2f\n", acc.name, amount)
	acc.balance+=amount
} 

func (acc *account)Withdraw (amount float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("账户余额不足，无法取款")
		}
	}()
	fmt.Printf("取款：%s - %.2f\n", acc.name, amount)
	if amount > acc.balance {
		panic("账户余额不足")
	}
	acc.balance-=amount
}

func (acc *account) PrintBalance() {
	fmt.Printf("账户余额：%.2f\n", acc.balance)
}

func main(){
	account := account{
		name:    "Alice",
		balance: 0.0,
	}

	var depositAmount, withdrawAmount float64
	//fmt.Print("请输入存款金额：")
	fmt.Scanln(&depositAmount)
	//fmt.Print("请输入取款金额：")
	fmt.Scanln(&withdrawAmount)

	account.Deposit(depositAmount)
	account.Withdraw(withdrawAmount)
	account.PrintBalance()
}