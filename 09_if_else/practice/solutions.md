# 💡 Solutions: If-Else

> [!NOTE]
> Implement your core logic in the code blocks below to master conditional branching.

---

### 1. Write an `if` statement that checks if a number is positive, negative, or zero.

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    if num > 0 {
        fmt.Println("Positive")
    } else if num < 0 {
        fmt.Println("Negative")
    } else {
        fmt.Println("Zero")
    }
}
```

---

### 2. Create a program that determines if a person is eligible to vote based on an `age` variable.

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Scan(&age)
    if age >= 18 {
        fmt.Println("Eligible to vote")
    } else {
        fmt.Println("Not eligible to vote")
    }
}
```

---

### 3. Can you use an `if` statement without an `else` block? Give an example.

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    if num > 0 {
        fmt.Println("Positive")
    }
}
```

---

### 4. Check if a string is empty using an `if` statement.

```go
package main

import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    if str == "" {
        fmt.Println("Empty string")
    } else {
        fmt.Println("Not empty string")
    }
}
```

---

### 5. Write a program that checks if a number is even or odd using `if-else`.

```go
package main

import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    if num%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

---
