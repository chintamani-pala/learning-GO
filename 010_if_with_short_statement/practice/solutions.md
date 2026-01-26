# 💡 Solutions: If with Short Statement

> [!NOTE]
> Implement your core logic in the code blocks below to master compact conditions.

---

### 1. Write an `if` statement that initializes a variable `num := 10` and then checks if it's greater than 5.

```go
package main

import "fmt"

func main() {
    if num := 10; num > 5 {
        fmt.Println("Greater than 5")
    }
}
```

---

### 2. Is the variable initialized in the `if` short statement accessible outside the `if` block?

```go
package main

import "fmt"

func main() {
    if num := 10; num > 5 {
        fmt.Println("Greater than 5")
    }
    fmt.Println(num)//Error: undefined: num
}
```

---

### 3. Use a short statement in an `if` to check the length of a string.

```go
package main

import "fmt"

func main() {
    if str := "hello"; len(str) > 5 {
        fmt.Println("Greater than 5")
    }
}
```

---

### 4. Rewrite a standard `if` block (with a variable declared before) into an `if` with a short statement.

```go
package main

import "fmt"

func main() {
    
    
    if num := 10; num > 5 {
        fmt.Println("Greater than 5")
    }
}
```

---

### 5. Explain a scenario where using an `if` with a short statement is cleaner than traditional declaration.

```go
Using an if with a short statement is cleaner when a variable is only required for the condition, because it reduces scope, improves readability, and prevents unnecessary variable leakage.
```

---
