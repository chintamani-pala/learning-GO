# 💡 Solutions: Basic Types - Int and Float

> [!NOTE]
> Implement your core logic in the code blocks below to master numeric operations and types.

---

### 1. Declare an `int` variable and a `float64` variable. Try to add them together. What happens and how do you fix it?

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b float64 = 20.5
    fmt.Println(a + b) //invalid operation: a + b (mismatched types int and float64)
    fmt.Println(float64(a) + b) //30.5
}
```

---

### 2. Write a program to calculate the area of a circle given a radius (use `math.Pi`).

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    radius := 5.0 //it should be a float64 because math.Pi is a float64
    fmt.Println(math.Pi * radius * radius) //78.5
}
```

---

### 3. What is the difference between `int8`, `int32`, and `int64`?

```go
package main

import "fmt"

func main() {
    fmt.Println(int8(127)) //int8 means it can store values from -128 to 127 ( -(2^7)+1 to 2^7 - 1)
    fmt.Println(int32(2147483647)) //int32 means it can store values from -2147483648 to 2147483647 ( -(2^31)+1 to 2^31 - 1)
    fmt.Println(int64(9223372036854775807)) //int64 means it can store values from -9223372036854775808 to 9223372036854775807 ( -(2^63)+1 to 2^63 - 1)

    //if you take a larger than the boundary it will consider as rotational
    var x int
    fmt.Scan(&x) //128

    var y int
    fmt.Scan(&y) //1

    fmt.Println(int8(x+y))//-127
}
```

---

### 4. Write a program that converts a `float64` value to an `int`. Does it round or truncate?

```go
package main

import "fmt"

func main() {
    var a float64 = 10.5
    fmt.Println(int(a)) //10 it will truncate the decimal
}
```

---

### 5. Find the remainder of 17 divided by 5 using the modulus operator.

```go
package main

import "fmt"

func main() {
    fmt.Println(17 % 5) //2
}
```

---
