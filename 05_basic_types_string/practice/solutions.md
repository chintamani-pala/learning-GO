# 💡 Solutions: Basic Types - String

> [!NOTE]
> Implement your core logic in the code blocks below to master string manipulation.

---

### 1. Create a string literal using double quotes and another using backticks. What is the difference?

| Double quotes `" "`                          | Backticks `` ` ` ``           |
| -------------------------------------------- | ----------------------------- |
| Supports escape sequences (`\n`, `\t`, `\"`) | No escape sequences           |
| Newlines need `\n`                           | Newlines are written directly |
| Interpreted string                           | Raw string                    |

**In short:**

* `" "` → processed by Go
* `` ` ` `` → taken exactly as written


---

### 2. How do you find the length of a string in Go?

```go
package main

import "fmt"

func main() {
    str := "Hello"
    fmt.Println(len(str))
}
```

---

### 3. Concatenate two strings "Golang" and "is Fun" and print the result.

```go
package main

import "fmt"

func main() {
    str1 := "Golang"
    str2 := "is Fun"
    fmt.Println(str1 + " " + str2)
}
```

---

### 4. Access the first character of the string "Hello" and print its numeric value (rune/byte).

```go
package main

import "fmt"

func main() {
    str := "Hello"
    fmt.Println(str[0])  // it will print the numeric value of the first character(H -> 72)
}
```

---

### 5. Use a loop to print each character of a string on a new line.

```go
package main

import "fmt"

func main() {
    str := "Hello"
    for _,i := range str {
        fmt.Println(string(i))
    }
}
```

---
