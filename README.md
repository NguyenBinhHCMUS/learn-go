# Go cơ bản

## 2.5 Kiểu Dữ Liệu trong Go Lang

### 2.5.1 Số nguyên

```go
package main

import "fmt"

func main() {
    var a int8 = -128
    var b uint8 = 255
    var c int16 = -32768
    var d uint16 = 65535
    var e int32 = -2147483648
    var f uint32 = 4294967295
    var g int64 = -9223372036854775808
    var h uint64 = 18446744073709551615

    fmt.Println(a, b, c, d, e, f, g, h)
}
```

### 2.5.2 Số Thực (Floating-Point Numbers)

Số thực là các số có phần thập phân. Trong Go, số thực được chia thành hai loại: float32 và float64. Số thực 64-bit (float64) có độ chính xác cao hơn và là kiểu mặc định nếu không chỉ định rõ ràng.

```go
package main

import "fmt"

func main() {
    var x float32 = 3.14
    var y float64 = 2.718281828459045

    fmt.Println(x, y)
}
```

### 2.5.3. Số Phức (Complex Numbers)

Số phức bao gồm một phần thực và một phần ảo. Trong Go, số phức được chia thành hai loại: complex64 và complex128.

```go
package main

import "fmt"

func main() {
    var z1 complex64 = 1 + 2i
    var z2 complex128 = 3 + 4i

    fmt.Println(z1, z2)
}
```

### 2.5.4. Chuỗi (Strings)

Chuỗi là một dãy các ký tự. Trong Go, chuỗi là bất biến, nghĩa là một khi chuỗi được tạo ra, giá trị của nó không thể thay đổi.

```go
package main

import "fmt"

func main() {
    var s1 string = "Hello, World!"
    s2 := "Go is awesome!"

    fmt.Println(s1)
    fmt.Println(s2)
}
```

### 2.5.5. Boolean

Kiểu dữ liệu boolean chỉ có hai giá trị: true và false. Kiểu dữ liệu này thường được sử dụng trong các biểu thức điều kiện.

```go
package main

import "fmt"

func main() {
    var b1 bool = true
    var b2 bool = false

    fmt.Println(b1)
    fmt.Println(b2)
}
```

### 2.5.6 Mảng (Arrays)

Mảng là một tập hợp các phần tử có cùng kiểu dữ liệu và có kích thước cố định. Các phần tử trong mảng được truy cập thông qua chỉ số.

```go
package main

import "fmt"

func main() {
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr)

    // Truy cập phần tử trong mảng
    fmt.Println(arr[0])
    fmt.Println(arr[4])
}
```

### 2.5.7 Cấu Trúc (Structs)

Cấu trúc là một tập hợp các trường (fields) có thể có các kiểu dữ liệu khác nhau. Cấu trúc cho phép chúng ta nhóm các dữ liệu liên quan lại với nhau.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var p Person
    p.Name = "John"
    p.Age = 30

    fmt.Println(p)
}
```

### 2.5.8 Con Trỏ (Pointers)

Con trỏ là một biến lưu trữ địa chỉ của một biến khác. Con trỏ cho phép chúng ta làm việc trực tiếp với địa chỉ bộ nhớ.

```go
package main

import "fmt"

func main() {
    var x int = 10
    var p *int = &x

    fmt.Println("Giá trị của x:", x)
    fmt.Println("Địa chỉ của x:", p)
    fmt.Println("Giá trị tại địa chỉ của p:", *p)
}
```

### 2.5.9 Slice

Slice là một kiểu dữ liệu tham chiếu đến một phần của mảng. Slice linh hoạt hơn mảng vì có thể thay đổi kích thước.

```go
package main

import "fmt"

func main() {
    var s []int = []int{1, 2, 3, 4, 5}
    fmt.Println(s)

    // Thêm phần tử vào slice
    s = append(s, 6)
    fmt.Println(s)
}
```

### 2.5.10 Bản Đồ (Maps)

Bản đồ là một tập hợp các cặp khóa-giá trị. Khóa và giá trị có thể có các kiểu dữ liệu khác nhau.

```go
package main

import "fmt"

func main() {
    var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
    fmt.Println(m)

    // Truy cập giá trị trong bản đồ
    fmt.Println(m["two"])

    // Thêm một cặp khóa-giá trị mới
    m["four"] = 4
    fmt.Println(m)
}
```

### 2.5.11 Các Kiểu Dữ Liệu Giao Diện

Giao diện là một tập hợp các phương thức mà một kiểu dữ liệu phải triển khai. Giao diện cho phép chúng ta định nghĩa các hành vi mà các kiểu dữ liệu khác nhau có thể thực hiện.

```go
package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func main() {
    var a Animal

    a = Dog{}
    fmt.Println(a.Speak())

    a = Cat{}
    fmt.Println(a.Speak())
}
```

## 2.6 Toán Tử trong Go Lang

Toán tử là các ký hiệu đặc biệt được sử dụng để thực hiện các phép toán trên các biến và giá trị. Trong Go, toán tử được chia thành nhiều loại khác nhau, bao gồm toán tử số học, toán tử so sánh, toán tử logic, toán tử gán, toán tử bitwise, và một số toán tử khác. Bài viết này sẽ cung cấp một cái nhìn tổng quan chi tiết về các loại toán tử trong Go, cùng với các ví dụ minh họa cụ thể.

### 2.6.1 Toán Tử Số Học (Arithmetic Operators)

Toán tử số học được sử dụng để thực hiện các phép toán cơ bản như cộng, trừ, nhân, chia và lấy phần dư.

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| + | Cộng hai toán hạng | A + B |
| - | Trừ toán hạng thứ hai khỏi toán hạng thứ nhất | A - B |
| * | Nhân hai toán hạng | A * B |
| / | Chia toán hạng thứ nhất cho toán hạng thứ hai | A / B |
| % | Lấy phần dư của phép chia | A % B |
| ++ | Tăng giá trị của biến lên 1 | A++ |
| -- | Giảm giá trị của biến đi 1 | A-- |

```go
package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10
    var c int

    c = a + b
    fmt.Printf("Line 1 - Value of c is %d\n", c)

    c = a - b
    fmt.Printf("Line 2 - Value of c is %d\n", c)

    c = a * b
    fmt.Printf("Line 3 - Value of c is %d\n", c)

    c = a / b
    fmt.Printf("Line 4 - Value of c is %d\n", c)

    c = a % b
    fmt.Printf("Line 5 - Value of c is %d\n", c)

    a++
    fmt.Printf("Line 6 - Value of a is %d\n", a)

    a--
    fmt.Printf("Line 7 - Value of a is %d\n", a)
}
```
Kết quả:
```plaintext
Line 1 - Value of c is 31
Line 2 - Value of c is 11
Line 3 - Value of c is 210
Line 4 - Value of c is 2
Line 5 - Value of c is 1
Line 6 - Value of a is 22
Line 7 - Value of a is 21
```

### 2.6.2 Toán Tử So Sánh (Comparison Operators)
Toán tử so sánh được sử dụng để so sánh hai giá trị. Kết quả của phép so sánh là một giá trị boolean (true hoặc false).

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| == | Bằng nhau | A == B |
| != | Không bằng nhau | A != B |
| > | Lớn hơn | A > B |
| < | Nhỏ hơn | A < B |
| >= | Lớn hơn hoặc bằng | A >= B |
| <= | Nhỏ hơn hoặc bằng | A <= B |

```go
package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10

    if a == b {
        fmt.Printf("Line 1 - a is equal to b\n")
    } else {
        fmt.Printf("Line 1 - a is not equal to b\n")
    }

    if a < b {
        fmt.Printf("Line 2 - a is less than b\n")
    } else {
        fmt.Printf("Line 2 - a is not less than b\n")
    }

    if a > b {
        fmt.Printf("Line 3 - a is greater than b\n")
    } else {
        fmt.Printf("Line 3 - a is not greater than b\n")
    }

    a = 5
    b = 20

    if a <= b {
        fmt.Printf("Line 4 - a is either less than or equal to b\n")
    }

    if b >= a {
        fmt.Printf("Line 5 - b is either greater than or equal to a\n")
    }
}
```

Kết quả:

```plaintext
Line 1 - a is not equal to b
Line 2 - a is not less than b
Line 3 - a is greater than b
Line 4 - a is either less than or equal to b
Line 5 - b is either greater than or equal to a
```

### 2.6.3 Toán Tử Logic (Logical Operators)
Toán tử logic được sử dụng để kết hợp các biểu thức điều kiện. Kết quả của phép toán logic là một giá trị boolean (true hoặc false).

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| && | Phép AND logic | A && B |
| \|\| | Phép OR logic | A \|\| B |
| ! | Phép NOT logic | !A |

```go
package main

import "fmt"

func main() {
    var a bool = true
    var b bool = false

    if a && b {
        fmt.Printf("Line 1 - Condition is true\n")
    } else {
        fmt.Printf("Line 1 - Condition is not true\n")
    }

    if a || b {
        fmt.Printf("Line 2 - Condition is true\n")
    }

    a = false
    b = true

    if a && b {
        fmt.Printf("Line 3 - Condition is true\n")
    } else {
        fmt.Printf("Line 3 - Condition is not true\n")
    }

    if !(a && b) {
        fmt.Printf("Line 4 - Condition is true\n")
    }
}
```

Kết quả:

```plaintext
Line 1 - Condition is not true
Line 2 - Condition is true
Line 3 - Condition is not true
Line 4 - Condition is true
```

### 2.6.4 Toán Tử Gán (Assignment Operators)
Toán tử gán được sử dụng để gán giá trị cho biến. Go hỗ trợ nhiều toán tử gán kết hợp với các toán tử số học.

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| = | Gán giá trị | A = B |
| += | Cộng và gán | A += B (tương đương với A = A + B) |
| -= | Trừ và gán | A -= B (tương đương với A = A - B) |
| *= | Nhân và gán | A *= B (tương đương với A = A * B) |
| /= | Chia và gán | A /= B (tương đương với A = A / B) |
| %= | Lấy phần dư và gán | A %= B (tương đương với A = A % B) |

```go
package main

import "fmt"

func main() {
    var a int = 21
    var b int = 10
    var c int

    c = a + b
    fmt.Printf("Line 1 - Value of c is %d\n", c)

    c += a
    fmt.Printf("Line 2 - Value of c is %d\n", c)

    c -= a
    fmt.Printf("Line 3 - Value of c is %d\n", c)

    c *= a
    fmt.Printf("Line 4 - Value of c is %d\n", c)

    c /= a
    fmt.Printf("Line 5 - Value of c is %d\n", c)

    c = 200
    c %= a
    fmt.Printf("Line 6 - Value of c is %d\n", c)
}
```
Kết quả:

```plaintext
Line 1 - Value of c is 31
Line 2 - Value of c is 52
Line 3 - Value of c is 31
Line 4 - Value of c is 651
Line 5 - Value of c is 31
Line 6 - Value of c is 11
```
### 2.6.5 Toán Tử Bitwise (Bitwise Operators)
Toán tử bitwise được sử dụng để thực hiện các phép toán trên các bit của số nguyên.

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| & | AND bitwise | A & B |
| \| | OR bitwise | A \| B |
| ^ | XOR bitwise | A ^ B |
| << | Dịch trái | A << 2 |
| >> | Dịch phải | A >> 2 |
| &^ | AND NOT bitwise | A &^ B |

```go
package main

import "fmt"

func main() {
    var a uint = 60      // 60 = 0011 1100
    var b uint = 13      // 13 = 0000 1101
    var c uint = 0

    c = a & b
    fmt.Printf("Line 1 - Value of c is %d\n", c)

    c = a | b
    fmt.Printf("Line 2 - Value of c is %d\n", c)

    c = a ^ b
    fmt.Printf("Line 3 - Value of c is %d\n", c)

    c = a << 2
    fmt.Printf("Line 4 - Value of c is %d\n", c)

    c = a >> 2
    fmt.Printf("Line 5 - Value of c is %d\n", c)
}
```
Kết quả:

```plaintext
Line 1 - Value of c is 12
Line 2 - Value of c is 61
Line 3 - Value of c is 49
Line 4 - Value of c is 240
Line 5 - Value of c is 15
```

### 2.6.6 Toán Tử Khác (Miscellaneous Operators)
Go cũng hỗ trợ một số toán tử khác như toán tử địa chỉ và toán tử con trỏ.

| Toán Tử | Mô Tả | Ví Dụ |
|---------|-------|-------|
| & | Trả về địa chỉ của biến | &a |
| * | Trỏ đến giá trị của biến | *a |

Ví dụ:

```go
package main

import "fmt"

func main() {
    var a int = 4
    var b int
    var ptr *int

    ptr = &a
    fmt.Printf("Address of a is %x\n", &a)
    fmt.Printf("Address stored in ptr is %x\n", ptr)
    fmt.Printf("Value of *ptr is %d\n", *ptr)

    b = *ptr
    fmt.Printf("Value of b is %d\n", b)
}
```
Kết quả:

```plaintext
Address of a is 0xc0000180a0
Address stored in ptr is 0xc0000180a0
Value of *ptr is 4
Value of b is 4
```