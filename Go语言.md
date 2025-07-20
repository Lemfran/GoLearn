# **Go语言**

## 一、Hello World!

```go
package main //1

import "fmt" //2

func main(){
    fmt.Printf("Hello World!") //3
}
```

1.首先`package main`是必要的，而且必须在`import`之前，独立文件必须是`package main`，项目文件有不同的包。

2.其次`import “fmt”`包含go语言中最基本的库，包括`Printf，Println ，Scanf` 等。还可以像下面这样用。

```go
import{
	"fmt"
	"math"
}
```

3.`fmt.Println(...)`将字符串输出到控制台，并在最后自动增加换行字符`\n`。用`fmt.Print("Hello, World!\n")`可以得到相同的结果。

4.程序的运行

步骤如下：

打开编辑器如Sublime2，将以上代码添加到编辑器中。

将以上代码保存为 hello.go

打开命令行，并进入程序文件保存的目录中。

输入命令 go run hello.go 并按回车执行代码。

如果操作正确你将在屏幕上看到 "Hello World!" 字样的输出。

（如果用的是vscode，这里如果有用户的输入，推荐这样使用。但是vscode可以直接编译运行，因为有插件，但是前面那种情况不允许）

5.注释在程序执行时将被忽略。//单行注释，/..../多行注释也叫块注释，不可以嵌套使用，一般用于包的文档描述或注释成块的代码片段。

6.{}中"{"不可以单独放一行。



## 二、基本数据类型

Go语言中有丰富的数据类型，除了基本的整形、浮点型、布尔型、字符串外，还有数组、切片、结构体、函数、map、通道（channel）等。Go语言的基本数据类型和其它语言大同小异。

### 1. 整型

按长度分为：int8,int16,int32,int64

对应的无符号整型：uint8,uint16,uint32,uint64

### 2. 浮点型

go语言中有两种浮点型数：float32和float64

实例：

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    fmt.Printf("%f\n", math.Pi)
    fmt.Printf("%.2f\n", math.Pi)
    asd := 3.1415926
    fmt.Printf("%T\n", asd)  
    asf := float32(3.14159)
    fmt.Printf("%T\n", asf) 
    fmt.Println(asf) 
}
```

### 3. 复数

 两种数据类型complate64和complate128

```go
package main
import (
    "fmt"
    "math"
)
func main() {
    var c1 complex64
    c1 = 1 + 2i
    var c2 complex128
    c2 = 2 + 3i
    fmt.Println(c1)
    fmt.Println(c2)
}
```

### 4. 布尔型

Go语言中以bool类型进行声明布尔类型数据，布尔类型数据只有true(真)和falde(假)两个值。

ps.

布尔类型变量的默认值为false。

Go语言中不允许将整形强制转换为布尔型。

布尔型无法参予数值运算，也无法与其它类型进行转换。

### 5. 字符串

**字符串默认不可以修改**

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64等）一样。Go语言里的字符串的内部实现使用UTF-8编码。字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

### 6. 字符串转义符

Go语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表格所示。

| 转义符 | 含义                               |
| ------ | ---------------------------------- |
| \r     | 回车符（返回行首）                 |
| \n     | 换行符（直接跳到下一行的同列位置） |
| \t     | 制表符                             |
| \\"    | 双引号                             |
| \\'    | 单引号                             |
| \\\    | 反斜杠                             |

**注意：使用反引号输出字符串（反引号用于原始字符串，不会转义）就是这个``**

课程测试：
我们要打印一个Windows平台下的一个文件路径：
输出
"c:\User\AppDate\go\"
c:\User\AppDate\go\

```go
package main

import "fmt"

func main(){
    str:="\"c:\\User\\AppDate\\go\\\""
    fmt.Println(str)
	fmt.Println(`c:\User\AppDate\go\`)
}
```



### 7. byte和rune类型

组成每个字符串的元素叫做“字符“，可以通过遍历或者单个获取字符串元素获得字符。字符用单引号（''）包裹起来，像：

var a = '汉' 
var b = 'x'
**Go语言的字符有以下两种：**

unit8类型，或者叫byte型，代表了ASCII码的一个字符。
rune类型，代表一个utf-8字符。

### 8. 类型转换

Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

强制类型转换的基本语法如下：
T(表达式)

其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等。

比如计算直角三角形的斜边长时**使用math包的Sqrt()函数，该函数接收的是float64类型的参数**，而变量a和b都是int类型的，这个时候就需要将a和b强制转换为float64类型:

```go
func sqrtDemo() {
    var a, b = 3, 4
    var c int
    c = int(math.Sqrt(float64(a*a + b*b)))  //强制转换；
    fmt.Println(c)
}
```



## 三、变量与常量

### 1. 变量的类型

变量（Variable）的功能是存储数据。常见的变量的数据类型有：整形、浮点型、布尔型等。
Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用，这个和c还有java很像。

### 2. 变量的声明

Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。

#### 2.1 标准声明

Go语言的变量声明格式为：

`var 变量名 变量类型`

变量声明以关键字var开头，变量类型放在变量的后面，行尾无需分号。

`var name string`
`var age int`
`var isOK bool`

#### 2.2 批量声明

每声明一个变量就需要写var关键字会比较繁琐，go语言中还支持批量变量的声明

```go
var(
    a string
    b int
    c bool
    d float32
)
```



### 3. 变量的初始化

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如：整形和浮点类型变量的默认值为0。字符串变量的默认值为空字符串。布尔类型莫认为false。切片、函数、指针变量默认为nil。

也可以在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

`var 变量名 类型 = 表达式`
例:
`var username string = "cuit"`
`var age int = 18`

在函数内部，可以使用更简略的:=方式声明并初始化变量。 无需制定类型:

 `a:=17`



### 总结：

变量声明的三种方式：

var name1 string = "Hashflag"
var name2 = "HashFlag"
函数内部专属：name3 := "HashFlag"



### 4. 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。常量的声明和变量的声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

`const pi = 3.14159275453`
`const e = 2.7182`
声明了pi e 这两个常量之后，在整个程序运行期间他们的值都不能再发生变化了。多个常量也可以一起声明：

```go
const(
    pi = 3.14159275453
    e = 2.7182
)
```

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。例如：

```go
const(
    n1 = 100
    n2
    n3
)
```



### 5.fmt格式化

如下：

```go
package main 

import "fmt"

// fmt占位符
func main(){
    var n = 100
    // 查看类型
    fmt.Printf("%T\n", n)  // 查看数据类型
    fmt.Printf("%v\n", n)  // 查看变量值
    fmt.Printf("%b\n", n)  // 二进制
    fmt.Printf("%d\n", n)  // 十进制
    fmt.Printf("%o\n", n)  // 八进制
    fmt.Printf("%x\n", n)  // 十六进制

    var s = "word"
    fmt.Printf("%s\n", s)  // 字符串 word  不会体现类型
    fmt.Printf("%v\n", s)  // 字符串 word  不会体现类型
    fmt.Printf("%#v\n", s)  // 字符串 "word"  体现出了具类型
}
```

在 Go 语言中，`fmt.Printf` 是一个格式化输出函数，`%s`、`%v` 和 `%#v` 是不同的格式化占位符，它们在输出时的行为有所不同：

#### 5.1 `%s` 格式化占位符
- `%s` 用于格式化字符串类型（`string`）。
- 当使用 `%s` 输出字符串时，只会输出字符串的内容，不会显示其类型信息。
- 例如：
  ```go
  var s = "word"
  fmt.Printf("%s\n", s)
  ```
  输出结果为：
  ```go
  word
  ```

#### 5.2 `%v` 格式化占位符
- `%v` 是一个通用的格式化占位符，用于输出变量的值。
- 它会根据变量的类型自动选择合适的格式进行输出。
- 对于字符串类型，`%v` 的输出效果与 `%s` 相同，不会显示类型信息。
- 例如：
  ```go
  var s = "word"
  fmt.Printf("%v\n", s)
  ```
  输出结果为：
  ```go
  word
  ```

#### 5.3 `%#v` 格式化占位符
- `%#v` 是 `%v` 的一个变体，用于输出变量的值，并且会显示其具体的类型信息。
- 对于字符串类型，`%#v` 会将字符串用双引号括起来，同时显示其类型为 `string`。
- 例如：
  ```go
  var s = "word"
  fmt.Printf("%#v\n", s)
  ```
  输出结果为：
  ```go
  "word"
  ```

#### 5.4 总结
- `%s` 和 `%v` 在输出字符串时，都不会显示类型信息，只会输出字符串的内容。
- `%#v` 会显示变量的具体类型信息，对于字符串，它会用双引号将字符串括起来，表示这是一个字符串类型的值。



好的！以下是完整的笔记内容，补充了每个部分的详细信息：

---

## 四、运算符

运算符用于在程序运行时执行数学或逻辑运算。Go 语言内置的运算符有：

1. 算术运算符
2. 关系运算符
3. 逻辑运算符
4. 位运算符
5. 赋值运算符

---

### 1. 算术运算符

算术运算符用于执行基本的数学运算。以下是 Go 语言中的算术运算符：

| 运算符 | 描述               | 示例    |
| ------ | ------------------ | ------- |
| `+`    | 加法               | `a + b` |
| `-`    | 减法               | `a - b` |
| `*`    | 乘法               | `a * b` |
| `/`    | 除法               | `a / b` |
| `%`    | 取模（求余数）     | `a % b` |
| `++`   | 自增（变量值加 1） | `a++`   |
| `--`   | 自减（变量值减 1） | `a--`   |

**示例代码：**
```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println("a + b =", a+b) // 输出 13
    fmt.Println("a - b =", a-b) // 输出 7
    fmt.Println("a * b =", a*b) // 输出 30
    fmt.Println("a / b =", a/b) // 输出 3
    fmt.Println("a % b =", a%b) // 输出 1

    a++
    fmt.Println("a++ =", a)     // 输出 11
    a--
    fmt.Println("a-- =", a)     // 输出 10
}
```

---

### 2. 关系运算符

关系运算符用于比较两个值的大小关系，结果为布尔值（`true` 或 `false`）。以下是 Go 语言中的关系运算符：

| 运算符 | 描述     | 示例     |
| ------ | -------- | -------- |
| `==`   | 等于     | `a == b` |
| `!=`   | 不等于   | `a != b` |
| `>`    | 大于     | `a > b`  |
| `<`    | 小于     | `a < b`  |
| `>=`   | 大于等于 | `a >= b` |
| `<=`   | 小于等于 | `a <= b` |

**示例代码：**
```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println("a == b =", a == b) // 输出 false
    fmt.Println("a != b =", a != b) // 输出 true
    fmt.Println("a > b =", a > b)   // 输出 true
    fmt.Println("a < b =", a < b)   // 输出 false
    fmt.Println("a >= b =", a >= b) // 输出 true
    fmt.Println("a <= b =", a <= b) // 输出 false
}
```

---

### 3. 逻辑运算符

逻辑运算符用于组合多个布尔表达式，结果为布尔值（`true` 或 `false`）。以下是 Go 语言中的逻辑运算符：

| 运算符 | 描述          | 示例     |
| ------ | ------------- | -------- |
| `&&`   | 逻辑与（AND） | `a && b` |
| `||`   | 逻辑或（OR）  | `a || b` |
| `!`    | 逻辑非（NOT） | `!a`     |

**示例代码：**
```go
package main

import "fmt"

func main() {
    a := true
    b := false

    fmt.Println("a && b =", a && b) // 输出 false
    fmt.Println("a || b =", a || b) // 输出 true
    fmt.Println("!a =", !a)         // 输出 false
}
```

---

### 4. 位运算符

位运算符用于对整数的二进制位进行操作。以下是 Go 语言中的位运算符：

| 运算符 | 描述          | 示例     |
| ------ | ------------- | -------- |
| `&`    | 位与（AND）   | `a & b`  |
| `|`    | 位或（OR）    | `a | b`  |
| `^`    | 位异或（XOR） | `a ^ b`  |
| `<<`   | 左移          | `a << n` |
| `>>`   | 右移          | `a >> n` |
| `^`    | 位取反（NOT） | `^a`     |

**示例代码：**
```go
package main

import "fmt"

func main() {
    a := 60  // 二进制 111100
    b := 13  // 二进制 001101

    fmt.Println("a & b =", a&b) // 输出 12 (二进制 001100)
    fmt.Println("a | b =", a|b) // 输出 61 (二进制 111101)
    fmt.Println("a ^ b =", a^b) // 输出 49 (二进制 011001)
    fmt.Println("a << 2 =", a<<2) // 输出 240 (二进制 11110000)
    fmt.Println("a >> 2 =", a>>2) // 输出 15 (二进制 00001111)
    fmt.Println("^a =", ^a)       // 输出 -61 (二进制 11111111 11111001)
}
```

---

### 5. 赋值运算符

赋值运算符用于将值赋给变量。以下是 Go 语言中的赋值运算符：

| 运算符 | 描述       | 示例                             |
| ------ | ---------- | -------------------------------- |
| `=`    | 简单赋值   | `a = b`                          |
| `+=`   | 加法赋值   | `a += b`（等同于 `a = a + b`）   |
| `-=`   | 减法赋值   | `a -= b`（等同于 `a = a - b`）   |
| `*=`   | 乘法赋值   | `a *= b`（等同于 `a = a * b`）   |
| `/=`   | 除法赋值   | `a /= b`（等同于 `a = a / b`）   |
| `%=`   | 取模赋值   | `a %= b`（等同于 `a = a % b`）   |
| `&=`   | 位与赋值   | `a &= b`（等同于 `a = a & b`）   |
| `|=`   | 位或赋值   | `a |= b`（等同于 `a = a | b`）   |
| `^=`   | 位异或赋值 | `a ^= b`（等同于 `a = a ^ b`）   |
| `<<=`  | 左移赋值   | `a <<= n`（等同于 `a = a << n`） |
| `>>=`  | 右移赋值   | `a >>= n`（等同于 `a = a >> n`） |

**示例代码：**
```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    a += b
    fmt.Println("a += b =", a) // 输出 13

    a -= b
    fmt.Println("a -= b =", a) // 输出 10

    a *= b
    fmt.Println("a *= b =", a) // 输出 30

    a /= b
    fmt.Println("a /= b =", a) // 输出 10

    a %= b
    fmt.Println("a %= b =", a) // 输出 1
}
```



## 五、流程控制

流程控制语句用于控制程序的执行流程，包括条件语句、循环语句和跳转语句。Go 语言提供了以下几种流程控制语句：

1. 条件语句（`if`、`switch`）
2. 循环语句（`for`、`range`）
3. 跳转语句（`break`、`continue`、`goto`、`return`）

---

### 1. 条件语句

#### 1.1 `if` 语句
`if` 语句用于根据条件执行代码块。

**语法：**
```go
if 条件 {
    // 条件为 true 时执行的代码
} else if 条件2 {
    // 条件2 为 true 时执行的代码
} else {
    // 其他情况执行的代码
}
```

**示例：**
```go
package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("成年人")
    } else {
        fmt.Println("未成年人")
    }
}
```

**特点：**
- Go 的 `if` 语句中，条件表达式不需要括号。
- 可以在 `if` 前面声明变量，变量的作用域仅限于 `if` 语句块。

**示例：**
```go
package main

import "fmt"

func main() {
    if num := 10; num > 0 {
        fmt.Println("正数")
    } else {
        fmt.Println("非正数")
    }
}
```

#### 1.2 `switch` 语句
`switch` 语句用于根据变量的值选择执行不同的代码块。

**语法：**
```go
switch 表达式 {
case 值1:
    // 表达式等于值1时执行的代码
case 值2:
    // 表达式等于值2时执行的代码
default:
    // 默认执行的代码
}
```

**示例：**
```go
package main

import "fmt"

func main() {
    grade := 'B'

    switch grade {
    case 'A':
        fmt.Println("优秀")
    case 'B':
        fmt.Println("良好")
    case 'C':
        fmt.Println("及格")
    default:
        fmt.Println("不及格")
    }
}
```

**特点：**
- Go 的 `switch` 语句不需要显式地写 `break`，每个 `case` 会自动退出。
- 可以使用类型断言、表达式等作为 `switch` 的条件。

**示例：**
```go
package main

import "fmt"

func main() {
    num := 10

    switch {
    case num > 0:
        fmt.Println("正数")
    case num < 0:
        fmt.Println("负数")
    default:
        fmt.Println("零")
    }
}
```

---

### 2. 循环语句

#### 2.1 `for` 循环
`for` 循环是 Go 语言中唯一的循环语句，用于重复执行代码块。

**语法：**
```go
for 初始化语句; 条件表达式; 后置语句 {
    // 循环体
}
```

**示例：**
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

**特点：**
- `for` 循环的初始化语句、条件表达式和后置语句都可以省略。
- 如果省略所有内容，`for` 循环将变成一个无限循环。

**无限循环示例：**
```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("无限循环")
        break // 防止程序陷入死循环
    }
}
```

#### 2.2 `range` 循环
`range` 用于遍历数组、切片、字符串、映射等集合类型。

**语法：**
```go
for 索引, 值 := range 集合 {
    // 循环体
}
```

**示例：**
```go
package main

import "fmt"

func main() {
    // 遍历数组
    arr := [3]int{1, 2, 3}
    for i, v := range arr {
        fmt.Printf("索引：%d, 值：%d\n", i, v)
    }

    // 遍历字符串
    str := "hello"
    for i, c := range str {
        fmt.Printf("索引：%d, 字符：%c\n", i, c)
    }

    // 遍历映射
    m := map[string]int{"a": 1, "b": 2}
    for k, v := range m {
        fmt.Printf("键：%s, 值：%d\n", k, v)
    }
}
```

**特点：**
- 在遍历字符串时，`range` 返回的是字符的索引和 Unicode 码点。
- 在遍历映射时，`range` 返回的是键和值。

---

### 3. 跳转语句

#### 3.1 `break` 语句
`break` 用于跳出当前的循环。

**示例：**
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
}
```

#### 3.2 `continue` 语句
`continue` 用于跳过当前循环的剩余部分，直接进入下一次循环。

**示例：**
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}
```

#### 3.3 `goto` 语句
`goto` 用于跳转到指定的标签处。

**语法：**
```go
goto 标签
...
标签:
```

**示例：**
```go
package main

import "fmt"

func main() {
    i := 0
    loop:
    if i > 5 {
        goto end
    }
    fmt.Println(i)
    i++
    goto loop

    end:
    fmt.Println("结束")
}
```

**特点：**
- `goto` 通常不推荐使用，因为它会使代码难以理解和维护。

#### 3.4 `return` 语句
`return` 用于从函数返回，可以返回一个或多个值。

**示例：**
```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println(result)
}
```

---

### 总结
- **条件语句**：`if` 和 `switch` 用于根据条件选择执行代码。
- **循环语句**：`for` 和 `range` 用于重复执行代码。
- **跳转语句**：`break`、`continue`、`goto` 和 `return` 用于控制程序的流程。

## 六、数组

数组是同种数据类型元素的集合。
在Go语言中，数组从声明时就确定，**使用时可以修改数组成员，但是数组大小不可变化**。

### 1.数组定义

var 数组变量名 [元素数量]T
注：数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。
[5]int和[10]int是不同类型。

 var a [3]int
 var b [4]int

 a = b // 不可以这样做，因为此时的a和b是不同的类型

 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，访问越界（下标在合法范围之外），则触发访问越界，会panic。

//在Go语言中，panic是一种运行时错误处理机制，用于表示程序在运行过程中遇到了无法恢复的错误。当panic被触发时，程序会立即中断当前的执行流程，终止当前的函数调用，并开始“回溯”（unwind）调用栈，直到程序完全终止。这种行为类似于其他语言中的“异常”（exception），但Go语言的panic通常用于更严重、无法恢复的错误场景。

### 2.数组的初始化

数组的初始化方式有很多种：

方式一

初始化数组时可以使用初始化列表来设置数组元素的值。

```go
func main() {
    var testArray [3]int                        //数组会初始化为int类型的零值
    var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
    var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
    fmt.Println(testArray)                      //[0 0 0]
    fmt.Println(numArray)                       //[1 2 0]
    fmt.Println(cityArray)                      //[北京 上海 深圳]
}
```

方式二

除了每次确保提供初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

```go
func main(){
    var testArray [3]int
    var numArray = [...]int{1, 2}  // [2]int
    var cityArray = [...]string{"北京", "上海", "深圳"} // [3]string
    fmt.Println(testArray)                          //[0 0 0]
    fmt.Println(numArray)                           //[1 2]
    fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
    fmt.Println(cityArray)                          //[北京 上海 深圳]
    fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}
```

方式三

我们还可以使用指定索引值的方式来初始化数组，例如：

```go
func main(){
    a := [...]int{1: 1, 6: 5}  // [7]int
    fmt.Println(a)                  // [0 1 0 5]
    fmt.Printf("type of a:%T\n", a) //type of a:[4]int 
}
```



### 3.数组的遍历

遍历数组a有以下两种方式：

方式一：for循环遍历

```go
func main() {
    var a = [...]string{"北京", "上海", "深圳"}
    // 方法1：for循环遍历
    for i := 0; i < len(a); i++ {
        fmt.Println(a[i])
    }
}
```


方式二：for range遍历

```go
func main() {
    var a = [...]string{"北京", "上海", "深圳"}
    // 方法2：for range遍历
    for index, value := range a {
        fmt.Println(index, value)
    }
}
```



### 4.多维数组

Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。

**二维数组的定义**

```go
func main(){
    a := [3][2]string{
        {"北京", "上海"},
        {"广州", "深圳"},
        {"成都", "重庆"},
    }
    fmt.Println(a)  
    fmt.Println(a[2][1])  
}
```

**二维数组的遍历**

```go
func main(){
    a := [3][2]string{
        {"北京", "上海"},
        {"广州", "深圳"},
        {"成都", "重庆"},
    }
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s\t", v2)
        }
        fmt.Println()
    }
}
```

输出：

北京  上海  
广州  深圳  
成都  重庆

注意：**多维数组只有第一层可以使用...来让编译器推到数组长度**。例如：

```go
// 支持的写法
a := [...][2]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
// 不支持多维数组的内层使用...
b := [3][...]string{
    {"北京", "上海"},
    {"广州", "深圳"},
    {"成都", "重庆"},
}
```



### 5.数组是值类型

**镜像拷贝；实参 不会改变**

**数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。**

```go
func modifyArray(x *[3]int) {
    x[0] = 100
}
func modifyArray2(x *[3][2]int) {
    x[2][0] = 100
}
func main() {
    a := [3]int{10, 20, 30}
    modifyArray(&a) //在modify中修改的是a的副本x
    fmt.Println(a) //[100 20 30]
    b := [3][2]int{
        {1, 1},
        {1, 1},
        {1, 1},
    }
    modifyArray2(&b) //在modify中修改的是b的副本x
    fmt.Println(b)  //[[1 1] [1 1] [100 1]]
}
```

**go语言的指针和c语言的指针的区别？？见其他**

## 七、切片

在 Go 语言中，**切片（Slice）** 是一种灵活的、动态的数组类型，它提供了对数组的动态访问和操作能力。切片是 Go 中非常强大的数据结构，广泛用于处理可变长度的序列数据。

### 1. **切片的定义**
切片是一个引用类型，它包含三个重要的部分：
- **指向数组的指针**：切片底层是一个数组，切片通过指针指向数组的某一部分。
- **长度（Length）**：切片中实际包含的元素数量，可以通过 `len(slice)` 获取。
- **容量（Capacity）**：切片可以扩展的最大长度，即它所引用的数组部分的总长度，可以通过 `cap(slice)` 获取。

### 2. **切片的创建**
切片可以通过多种方式创建：

**（1）使用 `make` 函数**

`make` 是创建切片的常用方法，它会分配一个底层数组，并返回一个切片。

```go
slice := make([]int, length, capacity)
```

- `length`：切片的初始长度。
- `capacity`：切片的容量（可选）。如果不指定容量，则默认等于长度。

```go
s1 := make([]int, 5)        // 长度为 5，容量为 5
s2 := make([]int, 3, 10)    // 长度为 3，容量为 10
```

**（2）使用数组的切片操作**

可以从一个数组或另一个切片中创建切片，通过指定起始索引和结束索引：

```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[1:4]         // 创建切片 [2, 3, 4]
```

**（3）使用字面量**

可以直接通过字面量创建切片，而无需显式指定长度和容量：

```go
slice := []int{1, 2, 3, 4, 5}
```

 **切片定义的基本语法： var a []int**

 **数组：指定长度 var a [3]int{1,2,3}**

### 3. **切片的特性**

**（1）动态性**

切片的长度可以动态变化，通过 `append` 函数可以向切片中添加元素。

```go
slice := []int{1, 2, 3}
slice = append(slice, 4)    // 添加元素 4，结果为 [1, 2, 3, 4]
slice = append(slice, 5, 6) // 添加多个元素，结果为 [1, 2, 3, 4, 5, 6]
```

当切片的容量不足以容纳新元素时，Go 会自动分配一个新的更大的底层数组，并将数据复制过去。

**（2）共享底层数组**

切片之间可能共享底层数组，修改一个切片可能会影响其他切片。例如：

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[1:3]               // s2 = [2, 3]
s2[0] = 100                 // 修改 s2，s1 也会受到影响
fmt.Println(s1)             // 输出 [1, 100, 3, 4, 5]
```

 **arr[1:3]:代表从arr数组的第2个(第一个是0),到第三个元素,但是不取第三个元素**

为了避免共享底层数组，可以使用 `copy` 函数创建一个独立的切片副本：

```go
s2 := make([]int, len(s1[1:3]))
copy(s2, s1[1:3])           // s2 是 s1 的一个独立副本
```

因此：切片是引用变量以及内存分配

  slice是对数组的一个引用，也就是说slice是直接对所引用的数组直接进行操作的，
  只要slice中的数据变了，数组中的数据也会跟着变

  slice从底层来说，其实就是一个struct结构体类型

  对于上面的slice我们可以理解为下面这个结构体

```go
  type slice struct {
    ptr *[2]int  // ptr是intArr的地址，也是intArr[0]的地址
    len int      // len是指目前slice占用的内存长度
    cap int      // cap是内存预计给slice分配的容量，可以动态变化
   }
```

**（3）切片的零值**

切片的零值是 `nil`，表示它没有指向任何数组。对 `nil` 切片调用 `len` 和 `cap` 会返回 0。

```go
var slice []int
fmt.Println(slice == nil)   // true
fmt.Println(len(slice))     // 0
fmt.Println(cap(slice))     // 0
```

### 4. **切片的用途**
切片在 Go 中非常灵活，广泛用于以下场景：
- **动态数组**：切片可以动态扩展，适合处理不确定长度的数据。
- **函数参数**：切片可以作为函数参数，方便传递和操作可变长度的数据。
- **并发**：切片可以用于在 Goroutine 之间传递数据（需注意并发安全）。
- **数据处理**：切片可以方便地对数据进行切分、拼接和操作。

### 5. **切片的注意事项**
- **容量管理**：虽然切片会自动扩容，但频繁扩容会影响性能。可以通过预分配足够容量来优化。
- **共享底层数组**：修改切片可能会影响其他切片，需要谨慎处理。
- **内存泄漏**：如果切片引用了大数组，但实际只使用了部分数据，可能会导致内存浪费。可以通过 `copy` 创建独立副本解决。

### 6.切片的长度与容量的解释

在 Go 语言中，切片的“长度”和“容量”是两个非常重要的概念，它们共同决定了切片的行为和性能。理解这两个概念对于正确使用切片非常重要。

**长度（Length）**

切片的长度是指切片中当前实际包含的元素数量。它可以通过 `len(slice)` 函数获取。

例如：
```go
slice := []int{1, 2, 3}
length := len(slice) // length 的值为 3
```

在这个例子中，切片 `slice` 包含 3 个元素 `[1, 2, 3]`，因此它的长度为 3。

**容量（Capacity）**

切片的容量是指切片底层数组可以容纳的最大元素数量。它可以通过 `cap(slice)` 函数获取。

容量的大小决定了切片在动态扩展时的性能。如果切片的容量足够大，可以直接在底层数组中添加新元素；如果容量不足，则需要重新分配一个更大的数组，并将数据复制过去。

例如：
```go
slice := make([]int, 3, 10) // 长度为 3，容量为 10
```

- **长度为 3**：表示切片当前包含 3 个元素。
- **容量为 10**：表示切片底层的数组可以容纳最多 10 个元素。

**长度和容量的关系**

1. **长度 ≤ 容量**：切片的长度永远不会超过它的容量。
2. **容量的作用**：容量决定了切片在动态扩展时的性能。如果容量足够大，切片可以直接在底层数组中添加新元素，而不需要重新分配数组。如果容量不足，切片会自动扩容，但这会涉及内存分配和数据复制，性能开销较大。

**例子**

```go
package main

import "fmt"

func main() {
    // 创建一个长度为 3，容量为 10 的切片
    slice := make([]int, 3, 10)
    fmt.Println("Initial slice:", slice)
    fmt.Println("Length:", len(slice)) // 长度为 3
    fmt.Println("Capacity:", cap(slice)) // 容量为 10

    // 向切片中添加元素
    slice = append(slice, 4, 5, 6)
    fmt.Println("After appending:", slice)
    fmt.Println("Length:", len(slice)) // 长度变为 6
    fmt.Println("Capacity:", cap(slice)) // 容量仍为 10

    // 继续添加元素，直到容量不足
    slice = append(slice, 7, 8, 9, 10, 11)
    fmt.Println("After more appending:", slice)
    fmt.Println("Length:", len(slice)) // 长度变为 11
    fmt.Println("Capacity:", cap(slice)) // 容量变为 20（自动扩容）
}
```

**输出**

```
Initial slice: [0 0 0]
Length: 3
Capacity: 10
After appending: [0 0 0 4 5 6]
Length: 6
Capacity: 10
After more appending: [0 0 0 4 5 6 7 8 9 10 11]
Length: 11
Capacity: 20
```

**解释**

1. **初始状态**：切片的长度为 3，容量为 10。这意味着切片当前有 3 个元素，但底层数组可以容纳 10 个元素。
2. **添加元素**：通过 `append` 向切片中添加了 3 个元素后，长度变为 6，但容量仍然是 10，因为底层数组足够大。
3. **自动扩容**：继续添加元素后，切片的长度超过了容量（10），因此 Go 自动扩容，将容量调整为 20，并将数据复制到新的底层数组中。

**总结**

- **长度**：表示切片当前实际包含的元素数量。
- **容量**：表示切片底层数组可以容纳的最大元素数量。
- **容量的作用**：容量决定了切片在动态扩展时是否需要重新分配内存。合理的容量可以减少内存分配的次数，提高性能。

通过合理管理切片的长度和容量，可以优化代码的性能和内存使用。

### 7.切片的细节

①切片初始化时， `var slice = arr[startIndex : endIndex]`

从`arr`数组下标为`startIndex`，取到下标为`endIndex`的元素，不含`arr[endIndex]`

②切片初始化时，仍然不能越界，范围在`[0 - len(arr)]`之间，但是可以动态增长

③几种简写方法：

```go
var slice = arr[0:end]  可简写为  var slice = arr[:end]
var slice = arr[start:len(arr)]  可简写为  var slice = arr[start:]
var slice = arr[0:len(arr)]  可简写为  var slice = arr[:]
```


④cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素

⑤切片定义完后，还不能使用，因为本身是一个空的，
需要让其引用到一个数组，或者make一个空间供切片来使用   

⑥切片可以继续切片。 

⑦内置函数append可以对切片进行动态的追加

```go
func main() {
    var slice []int = []int{1, 2, 3}
    fmt.Println("1. The value of slice is:", slice)
    slice1 := append(slice, 4, 5)
    fmt.Println("2. The value of slice1 is:", slice1) 
    slice = append(slice, 4, 5, 6)
    fmt.Println("3. The value of slice is:", slice) 
    slice = append(slice, slice...)   
    fmt.Println("4. The value of slice is:", slice) 
}
```

对切片append,其实就是对数组扩容，Golang会创建一个new的数组newArr，
安装扩容后的大小，然后将slice原来引用的数组拷贝到新的数组newArr，
再将等号左边的(上面是slice或者slice1)引用到newArr

⑧slice的copy 仅仅是 值传递

```go
func main()
{
   var slice1 []int = []int{1,2,3,4,5 }
   var slice2 []int = make([]int,10)
   fmt.Println("The value of slice2",slice2) //
   copy(slice1,slice2)
   fmt.Println("The value of slice1 is ",slice1) //
   fmt.Println("The value of slice2 is ",slice2) // 
}
```

好的！以下是关于 Go 语言指针的详细笔记内容，涵盖了指针的基本概念、操作以及常见用法。

---

## 八、指针

指针是 Go 语言中的一个重要概念，它存储变量的内存地址。通过指针，可以直接操作变量的内存，实现高效的数据处理和内存管理。

### 1. 指针的基本概念

指针是一个变量，它存储另一个变量的内存地址。使用指针可以间接访问和修改变量的值。

#### 1.1 指针的声明
指针的声明需要指定指针的类型，格式如下：
```go
var 指针变量名 *类型
```

**示例：**
```go
package main

import "fmt"

func main() {
    var a int = 10
    var p *int // 声明一个指向 int 类型的指针
    p = &a     // 将指针 p 指向变量 a 的地址

    fmt.Println("a 的值:", a)       // 输出 10
    fmt.Println("a 的地址:", &a)    // 输出 a 的内存地址
    fmt.Println("指针 p 的值:", p)  // 输出 p 的值（即 a 的地址）
    fmt.Println("指针 p 指向的值:", *p) // 输出 p 指向的值（即 a 的值）
}
```

#### 1.2 指针的解引用
通过指针访问变量的值，称为解引用（Dereferencing）。使用 `*` 符号可以解引用指针。

**示例：**
```go
package main

import "fmt"

func main() {
    var a int = 10
    var p *int = &a

    fmt.Println("a 的值:", a)       // 输出 10
    fmt.Println("指针 p 指向的值:", *p) // 输出 10

    *p = 20 // 修改指针 p 指向的值
    fmt.Println("修改后 a 的值:", a) // 输出 20
}
```

### 2. 指针的使用场景

#### 2.1 函数参数传递
在 Go 语言中，函数参数默认是值传递，即传递的是变量的副本。如果需要修改原始变量，可以使用指针传递。

**示例：**
```go
package main

import "fmt"

func modifyValue(value *int) {
    *value = 100 // 修改指针指向的值
}

func main() {
    a := 10
    fmt.Println("修改前 a 的值:", a) // 输出 10

    modifyValue(&a) // 传递 a 的地址
    fmt.Println("修改后 a 的值:", a) // 输出 100
}
```

#### 2.2 结构体指针
结构体指针用于直接操作结构体的字段，避免不必要的拷贝。

**示例：**
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println("修改前:", p) // 输出 {Alice 25}

    modifyPerson(&p) // 传递结构体的指针
    fmt.Println("修改后:", p) // 输出 {Bob 26}
}

func modifyPerson(person *Person) {
    person.Name = "Bob"
    person.Age++
}
```

### 3. 指针的注意事项

#### 3.1 空指针
空指针（`nil`）表示指针不指向任何有效的内存地址。尝试解引用空指针会导致运行时错误。

**示例：**
```go
package main

import "fmt"

func main() {
    var p *int = nil
    if p == nil {
        fmt.Println("p 是空指针")
    }

    // *p = 10 // 这会导致运行时错误
}
```

#### 3.2 指针的比较
指针可以进行比较，比较的是它们的内存地址。

**示例：**
```go
package main

import "fmt"

func main() {
    a := 10
    b := 10
    p1 := &a
    p2 := &b

    fmt.Println("p1 == p2:", p1 == p2) // 输出 false，因为 a 和 b 的地址不同
}
```

#### 3.3 指针的生命周期
指针的生命周期取决于它所指向的变量的生命周期。如果变量被销毁，指针将变成野指针（指向无效内存地址的指针）。

**示例：**
```go
package main

import "fmt"

func main() {
    var a int = 10
    p := &a

    fmt.Println("a 的值:", a)       // 输出 10
    fmt.Println("指针 p 指向的值:", *p) // 输出 10

    a = 20
    fmt.Println("修改后 a 的值:", a) // 输出 20
    fmt.Println("修改后指针 p 指向的值:", *p) // 输出 20
}
```

### 4. 指针的高级用法

#### 4.1 指针的指针
可以声明指针的指针，即指向指针的指针。

**示例：**
```go
package main

import "fmt"

func main() {
    a := 10
    p1 := &a
    p2 := &p1

    fmt.Println("a 的值:", a)           // 输出 10
    fmt.Println("指针 p1 的值:", p1)     // 输出 a 的地址
    fmt.Println("指针 p2 的值:", p2)     // 输出 p1 的地址
    fmt.Println("指针 p2 指向的值:", *p2) // 输出 p1 的地址
    fmt.Println("指针 p2 指向的值的值:", **p2) // 输出 a 的值
}
```

#### 4.2 指针与切片
切片本身是一个指针，指向底层数组的一部分。通过指针操作切片可以实现高效的内存管理。

**示例：**
```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    p := &slice[1]

    fmt.Println("slice:", slice) // 输出 [1 2 3]
    fmt.Println("指针 p 指向的值:", *p) // 输出 2

    *p = 100
    fmt.Println("修改后 slice:", slice) // 输出 [1 100 3]
}
```

### 5. 指针的总结

- **指针**：存储变量的内存地址。
- **声明**：`var 指针变量名 *类型`。
- **解引用**：通过 `*` 符号访问指针指向的值。
- **使用场景**：
  - 函数参数传递（避免拷贝）。
  - 结构体指针（直接操作结构体字段）。
- **注意事项**：
  - 空指针（`nil`）不能解引用。
  - 指针的生命周期依赖于所指向的变量。
- **高级用法**：
  - 指针的指针。
  - 指针与切片。

## 九、Map

### 1.map定义

Go语言中提供的映射关系容器为`map`，其内部使用散列表（hash）实现。
·是一种无序的基于`key-value`的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

Go语言中 `map`的定义语法如下：

```go
map[KeyType]ValueType
```

其中，

`KeyType`:表示键的类型，`ValueType`:表示键对应的值的类型。

**`map`类型的变量默认初始值为nil，需要使用make()函数来分配内存。**语法为：

```go
make(map[KeyType]ValueType, [cap])
```

其中`cap`表示`map`的容量，该参数虽然不是必须的，但是我们应该在初始化`map`的时候就为其指定一个合适的容量。

### 2.map基本使用

**map中的数据都是成对出现的**，map的基本使用示例代码如下：

```go
func main() {
    scoreMap := make(map[string]int, 8)//分配空间
    scoreMap["张三"] = 90
    scoreMap["小明"] = 100
    fmt.Println(scoreMap)
    fmt.Println(scoreMap["小明"])
    fmt.Printf("type of a:%T\n", scoreMap)
}
```

**写出输出结果：**

```
map[小明:100 张三:90]
100
type of a:map[string]int
```

**map也支持在声明的时候填充元素**，例如：

```go
 func main() {
    userinfo := map[string]string{
        "username": "cuit",
        "password": "123456",
    }     
    fmt.Println(userinfo)
    fmt.Println(userinfo["username"])
    fmt.Println(userinfo["password"])
}
```

**写出输出结果**

```
map[password:123456 username:cuit]
cuit
123456
```



### 3.判断某个键是否存在

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

举个例子：

```go
func main() {
    userinfo := make(map[string]string, 8)
    userinfo["jack"] = "123456"
    userinfo["rose"] = "A123456"
    v, ok := userinfo["peter"]
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("无此用户")
    }
}
```

**输出结果**

```
无此用户
```

### 4.map的遍历  

**Go语言中使用for range遍历map**

```go
func main() {
    userinfo := make(map[string]string, 8)
    userinfo["jack"] = "123456"
    userinfo["rose"] = "A123456"
    userinfo["peter"] = "B123456"
    for k,_ := range userinfo {
        fmt.Println(k)
    }     
}
```

**输出结果**

```
peter
jack
rose
```

这里`for k, _ := range userinfo` 只遍历了 `map` 的键（key），而忽略了值（value）。如果希望同时输出键和值，可以直接在 `range` 循环中同时获取键和值。

以下的代码，用于输出 `map` 中的所有键值对：

```go
package main

import "fmt"

func main() {
    userinfo := make(map[string]string, 8)
    userinfo["jack"] = "123456"
    userinfo["rose"] = "A123456"
    userinfo["peter"] = "B123456"

    // 遍历 map 并输出键和值
    for k, v := range userinfo {
        fmt.Printf("Key: %s, Value: %s\n", k, v)
    }
}
```

**输出结果**

```
Key: jack, Value: 123456
Key: rose, Value: A123456
Key: peter, Value: B123456
```

**代码解释**

**`for k, v := range userinfo`**：

- `range` 关键字用于遍历 `map`。
- `k` 是键（key），`v` 是值（value）。
- 在每次迭代中，`range` 会返回键和值，你可以直接使用它们。

### 5.使用delete()函数删除键值对

使用`delete()`内建函数从`map`中删除一组键值对，`delete()`函数的格式如下：

```
delete(map, key)
```

其中，
`map`:表示要删除键值对的map
`key`:表示要删除的键值对的键

示例代码如下：

```go
func main() {
    userinfo := make(map[string]string, 8)
    userinfo["jack"] = "123456"
    userinfo["rose"] = "A123456"
    fmt.Println(userinfo)
    delete(userinfo, "rose")
    delete(userinfo, "jack")
    if len(userinfo)==0{ //进行判断；
       fmt.Println("空的map")
    }else {
      fmt.Println(userinfo)
    }
}


```

### 6.按照指定顺序遍历map

```go
func main() {
    rand.Seed(time.Now().UnixNano()) //初始化随机数种子
    var scoreMap = make(map[string]int, 200)
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
        value := rand.Intn(100)          //生成0~99的随机整数
        scoreMap[key] = value
    }
    //取出map中的所有key存入切片keys
    var keys = make([]string, 0, 200)
    for key := range scoreMap {
        keys = append(keys, key)
    }
    //对切片进行排序
    sort.Strings(keys)
    //按照排序后的key遍历map
    for _, key := range keys {
        fmt.Println(key, scoreMap[key])
    }
}
```

### 7.元素为map类型的切片

```go
func main() {

var mapSlice = make([]map[string]string, 3)
for index, value := range mapSlice {
    fmt.Printf("index:%d value:%v\n", index, value)
}
fmt.Println("after init")
// 对切片中的map元素进行初始化
mapSlice[0] = make(map[string]string, 10)
mapSlice[0]["name"] = "小王子"
mapSlice[0]["password"] = "123456"
mapSlice[0]["address"] = "沙河"
for index, value := range mapSlice {
    fmt.Printf("index:%d value:%v\n", index, value)
}

}
```

**输出结果**

```
index:0 value:map[]
index:1 value:map[]
index:2 value:map[]
after init
index:0 value:map[address:沙河 name:小王子 password:123456]
index:1 value:map[]
index:2 value:map[]
```



### 8.值为切片类型的map

```go
func main() {
    var sliceMap = make(map[string][]string, 3)
    fmt.Println(sliceMap)
    fmt.Println("after init")
    key := "中国"
    value, ok := sliceMap[key]
    if !ok {
        value = make([]string, 0, 2)
    }
    value = append(value, "北京", "上海")
    sliceMap[key] = value
    fmt.Println(sliceMap)
}
```

**输出结果**

```
map[]
after init
map[中国:[北京 上海]]
```

**两者区分：**

`[]map[string]string`：`map`类型的切片 切片：`[]类型` ->这里类型是`map[key]value`  

`map[string][]string`：切片类型的`map` `map`：`map[key]value`  ->这里value是`[]类型`



## 十、函数

### 1 函数定义

Go语言中定义函数使用`func`关键字，具体格式如下：

```go
func 函数名(参数)(返回值){
    函数体
}
```

其中：

函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
参数：**参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。**
返回值：**返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。**

我们先来定义一个求两个数之和的函数：

```go
func intSum(x int, y int) int {
    return x + y
}

func intSum(x int, y int,z int ) int {
    return x + y+ z
}
```

函数的参数和返回值都是可选的，例如我们可以实现一个既不需要参数也没有返回值的函数：

```go
func sayHello() {
    fmt.Println("Hello cuit")
}
```



### 2 函数的调用

定义了函数之后，我们可以通过函数名()的方式调用函数。 例如我们调用上面定义的两个函数，代码如下：

```go
func main() {
    sayHello()
    ret := intSum(10, 20)
    fmt.Println(ret)
}
```

**注意，调用有返回值的函数时，可以不接收其返回值。**



### 3 参数

**3.1类型简写**

**函数的参数中如果相邻变量的类型相同，则可以省略类型**，例如：

```go
func intSum(x, y int) int {
    return x + y
}
```

上面的代码中，intSum函数有两个参数，这两个参数的类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型。

**3.2 可变参数**

**可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。**

**注意：可变参数通常要作为函数的最后一个参数。**

举个例子：

```go
func intSum2(x ...int) int {
    fmt.Println(x) //x是一个切片
    sum := 0
    for _, v := range x {
        sum = sum + v
    }
    return sum
}
```

调用上面的函数：

```go
ret1 := intSum2()
ret2 := intSum2(10)
ret3 := intSum2(10, 20)
ret4 := intSum2(10, 20, 30)
fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
```

函数名相同，参数不同；重载；//哪个语言可以实现重载？

固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：

```go
func intSum3(x int, y ...int) int {
    fmt.Println(x, y)
    sum := x
    for _, v := range y {
        sum = sum + v
    }
    return sum
}
```


调用上述函数：

```go
ret5 := intSum3(100)  //100
ret6 := intSum3(100, 10)  //110
ret7 := intSum3(100, 10, 20)//130
ret8 := intSum3(100, 10, 20, 30)// 160
fmt.Println(ret5, ret6, ret7, ret8) 
```

**重点：本质上，函数的可变参数是通过切片来实现的**



### 4 返回值

  Go语言中通过return关键字向外输出返回值。

**4.1 多返回值**

Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。

例子：

```go
func calc(x, y int) (  int,   int) {
    sum := x + y
    sub := x - y
    return sum, sub
}
```

调用示例

```go
i,j:=calc(x,y)  //go语言的好处，  变量名称1,...,:=函数调用 
```

**4.2 返回值命名**

 **函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。**

例如：

```go
func calc(x, y int) (sum, sub int) {
    sum = x + y
    sub = x - y
    return
}
```


4.3 返回值补充

**当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice**，没必要显示返回一个长度为0的切片。

```go
func someFunc(x string) []int {
    if x == "" {
        return nil // 没必要返回[]int{}
    }
    ...
}
```



### 5 变量作用域

5.1 全局变量

   全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
    fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
    testGlobalVar()  // num=10
}
```



5.2 局部变量


局部变量又分为两种： 函数内定义的变量无法在该函数外使用，
例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：

```go
func testLocalVar() {
    //定义一个函数局部变量x,仅在该函数内生效
    var x int64 = 100
    fmt.Printf("x=%d\n", x)  //x=100
}

func main() {
    testLocalVar()
    fmt.Println(x)  //无法访问 函数内部定义的变量
}
```


**如果局部变量和全局变量重名，优先访问局部变量。**

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
    num = 100
    fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}
func main() {
    testNum()   //num=100
}


```


for(int i=0;i<10;i++)

1)通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

```go
func testLocalVar2(x, y int) {
    fmt.Println(x, y) //函数的参数也是只在本函数中生效
    if x > 0 {
        z := 100  
        fmt.Println(z)
    }
    fmt.Println(z) // 有无问题 ???  有问题
}
```



2)for循环语句中定义的变量，也是只在for语句块中生效：

```go
func testLocalVar3() {
    for i := 0; i < 10; i++ {
        fmt.Println(i) //变量i只在当前for语句块中生效
    }
    fmt.Println(i) //有无问题 ???  有问题
}
```



### 6 函数类型与变量

**6.1 定义函数类型**   

 我们可以使用type关键字来定义一个函数类型，具体格式如下：

```go
type calculation func(int, int) int
```

上面语句定义了一个`calculation`类型，它是一种**函数类型**，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

简单来说，**凡是满足这个条件的函数都是`calculation`类型的函数**，例如下面的`add`和`sub`是`calculation`类型。

```go
func add(x, y int) int {
    return x + y
}

func sub(x, y int) int {
    return x - y
}

//add和sub都能赋值给calculation类型的变量。

var c calculation
c = add
```

**6.2 函数类型变量**

我们可以声明函数类型的变量并且为该变量赋值：

```go
type calculation func(int, int) int

func main() {
    var c calculation                
    c = add                          
    fmt.Printf("type of c:%T\n", c)  
    fmt.Println(c(1, 2))             
    f := add                         
    fmt.Printf("type of f:%T\n", f)  
    fmt.Println(f(10, 20))           
}
```




### 7 高阶函数

定义:**高阶函数分为函数作为参数和函数作为返回值两部分**。

**7.1函数可以作为参数**

```go
func add(x, y int) int {
    return x + y
}
func calc(x, y int, op func(int, int) int) int {
    return op(x, y)
}
func main() {
    ret2 := calc(10, 20, add)
    fmt.Println(ret2) //30
}
```

**7.2 函数作为返回值**

```go
func do(s string) (func(int, int) int, error) {
    switch s {
    case "+":
        return add, nil
    case "-":
        return sub, nil
    default:
        err := errors.New("无法识别的操作符")
        return nil, err
    }
}
```


//课中尝试 在main中实现调用；

```go
xx,error1:=do("*")    
fmt.Println(xx(3,4))
fmt.Println(error1)
```

### 8 匿名函数和闭包

**8.1 匿名函数**

   函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。
   匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
  func(参数)(返回值){
    函数体
  }
```

  匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
    // 将匿名函数保存到变量
    add := func(x, y int) {
        fmt.Println(x + y)
    }
add(10, 20) // 通过变量调用匿名函数

func(x, y int) {
    fmt.Println(x + y)
}(10, 20)//自执行函数：匿名函数定义完加()直接执行
}
```

**匿名函数多用于实现回调函数和闭包。**

**8.2 闭包**

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 

例子：

```go
func adder() func(int) int {
    var x int
    return func(y int) int {
        x += y
        return x
    }
}

func main() {
    var f = adder()
    fmt.Println(f(10))     
    fmt.Println(f(20))     
    fmt.Println(f(30))     
    f1 := adder()
    fmt.Println(f1(40))    
    fmt.Println(f1(50))   
}

```




变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 **在f的生命周期内，变量x也一直有效。** x//静态变量；


重点解释：

闭包就是一个函数和一个函数外的变量的封装，而且这个变量就对应着类中的静态变量。 
这样就可以将这个程序的输出结果解释的通了。



闭包进阶示例1：

```go
func adder2(x int) func(int) int {
    return func(y int) int {
        x += y
        return x
    }
}
func main() {
    var f = adder2(10)
    fmt.Println(f(10))  //20
    fmt.Println(f(20))  //40
    fmt.Println(f(30))  //70

    f1 := adder2(20)
    fmt.Println(f1(40))   //60
    fmt.Println(f1(50))   //110

}
```

这里`f`变成了`adder2`中的匿名函数，因为匿名函数中有变量`x`（是在`adder2`中的），但是其本身并没有，就相当于全局变量一样了，所以保留了`x`。

闭包进阶示例2：

```go
func makeSuffixFunc(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}

func main() {
    jpgFunc := makeSuffixFunc(".jpg")
    txtFunc := makeSuffixFunc(".txt")
    fmt.Println(jpgFunc("test")) //test.jpg
    fmt.Println(txtFunc("test")) //test.txt
}
```



### 9 defer语句

**Go语言中的defer语句会将其后面跟随的语句进行延迟处理。**
**在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，**
**也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。  //defer执行 如堆栈。**

例子：

```go
func main() {

    fmt.Println("start")
    defer fmt.Println(1)
    defer fmt.Println(2)
    defer fmt.Println(3)
    fmt.Println("end")

}
```

输出：

```
start
end
3
2
1
```

**由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。**

比如：资源清理、文件关闭、解锁及记录时间等。



**在Go语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和`RET`指令两步**。
而defer语句执行的时机就在返回值赋值操作后，`RET`指令执行前。 

由于defer语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。

比如：资源清理、文件关闭、解锁及记录时间等。

在Go语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和`RET`指令两步。
而`defer`语句执行的时机就在返回值赋值操作后，`RET`指令执行前。 

阅读下面的代码，写出最后的打印结果。

```go
func f1() int {
    x := 5
    defer func() {
        x++
    }()
    return x
} 
func f2() (x int) {
    defer func() {
        x++
    }()
    return 5
} 
func f3() (y int) {
    x := 5
    defer func() {
        x++
    }()
    return x
}
func f4() (x int) {
    defer func(x int) {
        x++
    }(x)
    return 5
} 
func main() {

    fmt.Println(f1())
    fmt.Println(f2())
    fmt.Println(f3())
    fmt.Println(f4())

}
```

①`defer`执行之前，将x赋值给了返回值（这是一个值拷贝），然后修改x的值，对返回值是无影响的，所以返回的是5

②返回值的名称就是x，此时defer执行前把x赋值为5，然后`defer`修改x的值， x被增加，故返回的是6

③返回值名称是y，`defer`执行前，y被赋值为5，`defer`执行修改x对y无影响，返回也是5

④返回值名称虽然是x，但是`defer`执行的`func`是一个带参数的函数，此时传入的参数x是一个值拷贝，作用域是内部，对于外部的x无影响，所以返回的也是5

原理：
`return`会将返回值先保存起来。
对于无名返回值来说：返回值保存在临时对象中，defer无法处理这个临时对象
对于有名返回值来说：就保存在已命名的变量中，defer可以处理这个临时对象



defer面试题

```go
func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}
func main() {
    x := 1
    y := 2
    defer calc("AA", x, calc("A", x, y))
    x = 10
    defer calc("BB", x, calc("B", x, y))
    y = 20    
}
```

问，上面代码的输出结果是？ 

```go
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
```

## 十一、结构体

共用体： 成员共享内存空间，结构体成员所占的空间，以最大成员所占空间。
结构体： 各个成员 所占空间之和。

### 1.go结构体描述

 Go语言通过自定义的方式形成新的类型，结构体是类型中都有成员的复合类型。
 // type 
 Go语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性。

结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”。

字段有以下特性：

- 字段拥有自己的类型和值。

- 字段名必须唯一。
- 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

结构体的存储空间是连续的，字段按照声明时的顺序存放 
<提示> Go语言没有“类”的概念，也不支持“类”的继承等面向对象的概念。



Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
Go语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。

### 2.结构体定义


结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

```go
variable_name := structure_variable_type {value1, value2...valuen}
```

或

```go
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
```

实例如下：

```go
 package main
 import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
    // 忽略的字段为 0 或 空

   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```

### 3.访问结构体成员

如果要访问结构体成员，需要使用点号 . 操作符，格式为：

`结构体.成员名`

### 4.结构体作为函数参数

你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：

```go
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook(Book1)

   /* 打印 Book2 信息 */
   printBook(Book2)
}

func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
```

### 5.结构体指针

 可以定义指向结构体的指针类似于其他指针变量，格式如下：

```go
var struct_pointer *Books
```

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

```go
struct_pointer = &Book1
```

使用结构体指针访问结构体成员，使用 "." 操作符：  //->

```go
struct_pointer.title // ??
```

接下来让我们使用结构体指针重写以上实例，代码如下：

```go
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
```

### 6.构造函数

```go
type person struct {
    name string
    city string
    age int8
}
```


Go语言的结构体没有构造函数，但可以实现。

示例：

```go
func newPerson(name, city string, age int8) *person {
    return &person{
        name: name,
        city: city,
        age:  age,
    }
}
```

调用构造函数示例：

```go
p9 := newPerson("eric", "成都"， 18)
fmt.Printf("%#v\n", p9) //&main.person{name:"eric", "成都", age:18}
```

这段代码是用 Go 语言编写的，`*` 和 `&` 在这里分别表示指针类型和取地址操作，它们的作用和含义不同，以下是具体解释：

**1\. `*`（指针类型）**

在函数返回值类型 `*person` 中，`*` 表示返回的是一个指向 `person` 类型的指针。

- **指针类型的作用**：在 Go 中，指针用于存储变量的内存地址。通过指针，可以间接访问和修改变量的值。返回一个指针类型，意味着调用者将获得一个指向 `person` 结构体实例的指针，而不是结构体的副本。这种方式可以节省内存（尤其是对于较大的结构体），并且允许调用者通过指针直接修改结构体的内容。
- **示例**：假设有一个 `person` 结构体定义如下：
  ```go
  type person struct {
      name string
      city string
      age  int8
  }
  ```
  函数返回 `*person` 指针后，调用者可以通过指针访问和修改 `person` 的字段，例如：
  ```go
  p := newPerson("Alice", "Beijing", 30)
  fmt.Println(p.name) // 输出 Alice
  p.age = 31          // 通过指针修改 age 字段
  ```

**2\. `&`（取地址操作符）**

在 `return &person{...}` 中，`&` 是取地址操作符，用于获取变量的内存地址。

- **作用**：`&` 用于将一个普通变量的地址取出来，生成一个指向该变量的指针。在这里，`person{...}` 是一个 `person` 类型的结构体字面量，它会创建一个 `person` 类型的实例。`&person{...}` 则会获取这个实例的内存地址，并将其转换为一个指向 `person` 的指针。
- **示例**：
  ```go
  p := person{name: "Alice", city: "Beijing", age: 30}
  ptr := &p // ptr 是一个指向 p 的指针
  fmt.Println(ptr) // 输出类似 &{Alice Beijing 30}
  ```

**为什么一个是 `*`，一个是 `&`**

- **`*` 是指针类型声明**：在函数返回值类型中，`*person` 表示函数返回的是一个指向 `person` 的指针类型。这是对返回值的类型声明，告诉调用者函数返回的是一个指针，而不是一个普通的 `person` 实例。***将变量变成指针**
- **`&` 是取地址操作**：在函数体中，`&person{...}` 是对一个 `person` 实例取地址的操作。`person{...}` 创建了一个 `person` 实例，`&` 将这个实例的地址取出，生成一个指向它的指针，并将这个指针返回。**&将实例变成指针**

**总结**

- `*` 用于声明指针类型，表示函数返回的是一个指针。
- `&` 用于取地址，将一个普通变量的地址取出，生成一个指针。
- 在这段代码中，`*` 和 `&` 的使用是为了实现函数返回一个指向 `person` 实例的指针，而不是直接返回一个 `person` 实例的副本。

### 7.方法和接收者

Go语言中的方法（Method）是一种用于特定类型变量的函数。
这种特定类型变量叫接收者（Receiver）。
接收者的概念类似于其它语言中的this或者self。

定义方式：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

接收者变量：接收者中的参数变量在命名时，官方建议使用接收者类型名称首字母的小写，而不是self或this等命名。例：Dog 命名为 d
接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：具体格式与函数定义相同。

## 十二、接口

引入了面向对象的概念

### 1.什么是接口

把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

```go
/* 定义接口 */
type 接口名 interface {
   方法1 [返回类型]
   方法2 [返回类型]
   方法3 [返回类型]
   ...
   方法N [返回类型]
}

/* 定义结构体 */
type 结构体名 struct {
   /* 结构体成员 */
}

/* 实现接口方法 */
func (结构体实例 结构体名) 方法名1() [return_type] {
   /* 方法实现 */
}
...
func (结构体实例 结构体名) 方法名2() [return_type] {
   /* 方法实现*/
}
```

### 2.接口定义流程说明

定义接口
定义一个或多个结构体
为每个结构体写一些方法
写一个接口，将这些有共性的方法定义在一起
使用接口
实例化结构体
将结构体的实例传入接口，并实现方法

```go
// 定义接口
type Shape interface {
    Area() float64
    Perimeter() float64
}


// 定义结构体类型Circle和Rectangle，并实现Shape接口的方法
type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}


func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}
    shapes := []Shape{circle, rectangle}

    for _, shape := range shapes {
        fmt.Printf("Area: %.2f\n", shape.Area())
        fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
        fmt.Println("----------------------")
    }

}
```

案例：动物发声
    需求描述：

定义一个接口 Animal，包含一个方法 Speak()，用于表示动物发声。
然后定义两个结构体 Dog 和 Cat，分别实现 Animal 接口，最后通过接口调用不同动物的发声方法。
实现步骤：
1.定义接口 Animal，包含方法 Speak() string。
2.定义结构体 Dog 和 Cat，分别实现 Animal 接口。
3.在 main 函数中，通过接口调用不同动物的发声方法。

接口的定义：
接口是一组方法的集合。
通过定义接口，可以规定多个不同类型需要实现的方法。
接口的实现：
结构体通过实现接口中定义的所有方法来满足接口。
在本例中，Dog 和 Cat 都实现了 Animal 接口的 Speak 方法。
接口的使用：
在函数参数中使用接口类型，可以接受任何满足该接口的类型。
通过接口调用方法时，会根据实际类型的实现来执行对应的方法。

```go
package main

import "fmt"

// 定义接口 Animal
type Animal interface {
    Speak() string
}

// 定义结构体 Dog
type Dog struct {
    Name string
}

// 实现 Animal 接口的 Speak 方法
func (d Dog) Speak() string {
    return "汪汪"
}

// 定义结构体 Cat
type Cat struct {
    Name string
}

// 实现 Animal 接口的 Speak 方法
func (c Cat) Speak() string {
    return "喵喵"
}

func main() {
    // 创建 Dog 和 Cat 实例
    myDog := Dog{Name: "旺财"}
    myCat := Cat{Name: "妙妙"}

    // 定义一个函数，接受 Animal 接口作为参数
    makeSound := func(a Animal) {
        fmt.Println(a.Speak())
    }
    
    // 调用函数，传入不同的动物
    makeSound(myDog) // 输出：汪汪
    makeSound(myCat) // 输出：喵喵

}
```



## 其他

### 1.`len( ... )` 求长度

在 Go 语言中，`len` 是一个内置函数，用于获取各种数据结构的长度。它的用法非常广泛，适用于切片、字符串、数组、通道等类型。以下是 `len` 的具体用法和一些常见场景：

**1.1 获取字符串的长度**

`len` 可以获取字符串的字节长度。在 Go 中，字符串是 UTF-8 编码的，因此一个字符可能占用多个字节。

```go
s := "Hello, 世界"
length := len(s)
fmt.Println(length) // 输出 13，因为 "Hello, " 占用 7 个字节，"世界" 占用 6 个字节
```

**注意**：如果需要获取字符串中字符的数量（而不是字节长度），可以使用 `rune` 类型进行转换，因为一个 `rune` 表示一个 Unicode 字符。

```go
s := "Hello, 世界"
runeCount := len([]rune(s))
fmt.Println(runeCount) // 输出 9
```

**1.2 获取数组的长度**

`len` 可以获取数组的长度，即数组中元素的数量。

```go
arr := [5]int{1, 2, 3, 4, 5}
length := len(arr)
fmt.Println(length) // 输出 5
```

数组的长度是固定的，因此 `len` 的结果不会改变。

**1.3 获取切片的长度**

`len` 用于获取切片的当前长度，即切片中实际包含的元素数量。

```go
slice := []int{1, 2, 3, 4, 5}
length := len(slice)
fmt.Println(length) // 输出 5
```

切片的长度可以通过操作（如添加或删除元素）动态变化，`len` 会返回最新的长度。

**1.4  获取通道的长度**

`len` 可以获取通道中当前的元素数量，但只能用于带缓冲的通道。

```go
ch := make(chan int, 3) // 带缓冲的通道
ch <- 1
ch <- 2

length := len(ch)
fmt.Println(length) // 输出 2
```

对于无缓冲的通道，`len` 的结果总是 0。

**1.5 获取字典（map）的长度**

`len` 可以获取字典中键值对的数量。

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
length := len(m)
fmt.Println(length) // 输出 3
```

**1.6 使用场景**

- **循环遍历**：在遍历数组、切片或字符串时，`len` 常用于控制循环的范围。
- **动态操作**：在操作切片时，`len` 可以帮助判断当前切片的大小。
- **通道状态检查**：在并发编程中，`len` 可以用于检查通道中是否有数据。



### 2.fmt.Printf("%+V")格式化占位符

在 Go 语言中，`fmt.Printf` 是一个格式化输出函数，用于将数据按照指定的格式打印到标准输出（通常是终端或控制台）。`%+v` 是 `fmt.Printf` 中的一个格式化占位符，用于控制输出的格式。

**解释 `%+v` 的含义：**

- `%v` 是一个通用的格式化占位符，表示以默认格式输出变量的值。
- 加号 `+` 是一个修饰符，当与 `%v` 结合时，它会以更详细的方式输出变量的值，特别是对于结构体类型，会显示字段名和字段值。

**具体作用：**

1. **对于普通变量**（如整数、字符串、布尔值等）：
   - `%v` 和 `%+v` 的输出效果相同，都是以默认格式输出变量的值。
   - 例如：
     ```go
     a := 10
     fmt.Printf("%v\n", a)    // 输出：10
     fmt.Printf("%+v\n", a)   // 输出：10
     ```

2. **对于结构体**：
   - `%v` 只会输出结构体的字段值，而 `%+v` 会同时输出字段名和字段值，更便于调试。
   - 例如：
     ```go
     type Person struct {
         Name string
         Age  int
     }
     
     p := Person{Name: "Alice", Age: 30}
     fmt.Printf("%v\n", p)    // 输出：{Alice 30}
     fmt.Printf("%+v\n", p)   // 输出：{Name:Alice Age:30}
     ```

**总结：**

`fmt.Printf("%+v\n")` 是一种格式化输出的方式，主要用于调试时更清晰地展示变量的结构和内容，尤其是对于结构体类型，能够帮助开发者快速了解字段的名称和对应的值。

3.

`fmt.Sprintf` 是 Go 语言标准库中的一个函数，用于根据格式化占位符生成格式化的字符串，而不是直接将结果打印到标准输出。它的功能类似于 `fmt.Printf`，但返回的是一个字符串，而不是直接输出到终端。

### 3.函数签名
```go
func Sprintf(format string, a ...interface{}) string
```

- **`format`**：格式化字符串，包含占位符（如 `%d`、`%s`、`%v` 等）。
- **`a ...interface{}`**：可变参数，表示要格式化的值。
- **返回值**：返回一个格式化后的字符串。

**使用示例**

**示例 1：格式化基本类型**

```go
package main

import (
	"fmt"
)

func main() {
	age := 25
	name := "Alice"
	message := fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
	fmt.Println(message)
}
```
**输出：**

```
Hello, my name is Alice and I am 25 years old.
```

**示例 2：格式化浮点数**

```go
package main

import (
	"fmt"
)

func main() {
	pi := 3.141592653589793
	formattedPi := fmt.Sprintf("Pi is approximately %.2f", pi)
	fmt.Println(formattedPi)
}
```
**输出：**
```
Pi is approximately 3.14
```
- `%.2f` 表示保留两位小数。

**示例 3：格式化结构体**

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Bob", Age: 30}
	info := fmt.Sprintf("%+v", p)
	fmt.Println(info)
}
```
**输出：**
```
{Name:Bob Age:30}
```

**示例 4：格式化日期和时间**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	formattedTime := fmt.Sprintf("Current time is %02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
	fmt.Println(formattedTime)
}
```
**输出：**
```
Current time is 14:35:47
```
- `%02d` 表示以两位数字格式化整数，不足两位时前面补零。

**常用格式化占位符**

- `%d`：十进制整数。
- `%f`：浮点数。
- `%s`：字符串。
- `%v`：默认格式输出值。
- `%+v`：输出值并显示字段名（结构体）。
- `%T`：输出值的类型。
- `%t`：布尔值（`true` 或 `false`）。
- `%q`：双引号括起来的字符串。
- `%x`：十六进制数字。
- `%02d`：两位数字，不足两位时前面补零。

### 4.go语言的指针和c语言的指针的区别

C语言：指针是C语言的核心特性之一，广泛用于数组操作、动态内存分配、结构体操作等。指针的灵活性使得C语言在系统编程和底层开发中非常强大。

int* p = malloc(sizeof(int)); // 动态分配内存
*p = 10;
free(p); // 释放内存

Go语言：虽然Go语言也有指针，但它的使用场景相对较少。Go语言提供了高级数据结构（如切片、map）和垃圾回收机制，减少了对指针的直接操作需求。
指针在Go语言中主要用于：
按引用传递（避免拷贝大对象）。
操作结构体字段（尤其是方法接收器）。
实现接口时传递可变状态。

### 5.Go 语言中数组和切片的特性，以及它们与 C 语言指针的比较

**1. Go 语言中的数组和切片**

**数组（Array）**

- **值类型**：数组是值类型，这意味着在赋值或传递参数时，会复制整个数组的内容。
- **固定大小**：数组的大小是固定的，一旦定义，不能改变。
- **行为**：由于是值类型，修改副本不会影响原始数组。

在你的代码中：

go复制

```go
func modifyArray(x *[3]int) {
    x[0] = 100
}
```

这里传入的是数组的指针（`*`），因此函数内部对数组的修改会直接影响原始数组。如果你传入的是数组本身（而不是指针），则不会影响原始数组。

**切片（Slice）**

- **引用类型**：切片是一个动态的、灵活的数组，它引用底层数组的一部分。
- **动态大小**：切片的大小可以动态变化，通过 `append` 等操作可以扩展切片。
- **行为**：切片之间可能共享底层数组，因此修改一个切片可能会影响其他切片。

**2. Go 语言的切片是否代替了 C 语言的指针？**

**答案：不完全正确。**

**相似之处**

1. **动态性**：切片类似于 C 语言中的动态数组（通过指针操作）。它们都可以动态地扩展大小。
2. **引用特性**：切片是引用类型，类似于 C 语言中的指针，修改切片可能会影响其他引用同一底层数组的部分。

**不同之处**

1. **内存管理**：
   - **C 语言**：指针需要手动管理内存，容易出错（如内存泄漏、野指针等）。
   - **Go 语言**：切片的内存管理由运行时自动完成，无需手动分配或释放内存。
2. **安全性**：
   - **C 语言**：指针操作容易引发安全问题（如越界访问、指针悬垂等）。
   - **Go 语言**：切片操作是安全的，运行时会检查越界，并且切片的扩容机制会自动处理底层数组的重新分配。
3. **语义清晰**：
   - **C 语言**：指针的使用需要明确地处理指针的解引用、地址运算等操作，语义较为复杂。
   - **Go 语言**：切片的使用更加直观，隐藏了底层的复杂性。例如，`append` 函数可以方便地扩展切片，而无需手动管理内存。

**3. Go 语言的切片和数组的使用场景**

**数组**

- **固定大小**：适用于大小固定且不会改变的场景。
- **值类型**：传递数组时会复制整个数组，适合小数组或需要值语义的场景。

**切片**

- **动态大小**：适用于大小不确定或需要动态扩展的场景。
- **引用类型**：适合处理大量数据，因为切片的底层数组可以共享，减少内存拷贝。

**4. 你的代码分析**

在你的代码中：

```go
func modifyArray(x *[3]int) {
    x[0] = 100
}
```

这里传入的是数组的指针，因此函数内部对数组的修改会直接影响原始数组。

```go
func modifyArray2(x *[3][2]int) {
    x[2][0] = 100
}
```

这里传入的是二维数组的指针，同样会直接影响原始数组。

如果你希望避免修改原始数组，可以传递数组的值（而不是指针）：

```go
func modifyArray(x [3]int) {
    x[0] = 100 // 这里修改的是副本，不会影响原始数组
}
```

C语言的数组本质是指针，因此如果同是上述这样的话，则会改变。

**总结**

- **切片**：是 Go 语言中处理动态数组的首选方式，它提供了灵活性和安全性。
- **数组**：适用于固定大小的场景，是值类型，传递时会复制整个数组。
- **与 C 语言指针的关系**：切片在某些方面类似于 C 语言的指针（动态性和引用特性），但它更安全、更易用，且隐藏了底层的复杂性。

因此，不能简单地说切片完全代替了 C 语言的指针，但切片确实解决了指针在动态数组操作中的很多问题。
