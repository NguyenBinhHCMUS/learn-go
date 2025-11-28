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