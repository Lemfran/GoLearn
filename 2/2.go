/*
7-2 defer使用
请编写一个程序，模拟一个简单的银行账户系统。该系统可以记录用户的账户余额，并提供存款和取款的功能。 要求如下： 程序需要定义一个 Account 结构体，包含以下字段： name：表示账户名的字符串 
balance：表示账户余额的浮点数 Account 结构体需要提供以下方法： Deposit(amount float64)：存款方法，将传入的金额加到账户余额中。 Withdraw(amount float64)：取款方法，从账户余额
中扣除传入的金额。如果余额不足，则无法取款。 PrintBalance()：打印当前账户余额。 使用 defer 关键字，在存款和取款方法中添加日志记录功能。在存款时打印格式为 "存款：账户名 + 存款金额" 的
日志，取款时打印格式为 "取款：账户名 - 取款金额" 的日志。 主函数中创建一个账户实例，并演示存款和取款的操作，以及打印账户余额。 请根据以上要求编写程序，并保证正确处理边界情况。 输入输出示
例： （假设账户名为 "Alice"） 注意： 使用 fmt.Printf 函数打印日志和账户余额，保留两位小数，带回车。 需要使用 defer 关键字来确保日志的打印顺序。 在取款时，如果账户余额不足，需要打印错
误信息："账户余额不足，无法取款"。

输入格式:
输入：
 存款金额：100.0
 取款金额：50.0

输出格式:
存款：Alice + 100.0
取款：Alice - 50.0
账户余额：50.0

输入样例:
在这里给出一组输入。例如：

100
50
输出样例:
在这里给出相应的输出。例如：

存款：Alice + 100.00
取款：Alice - 50.00
账户余额：50.00
*/

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