# 💡 Solutions: Setup First Program

> [!NOTE]
> Implement your core logic in the code blocks below to master basic Go program structure.

---

### 1. Write a program that prints "Hello, Go Developer!" to the console.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go Developer!")
}
```

---

### 2. Explain the purpose of `package main` in a Go file.

```go
It tells the Go compiler that this file belongs to the main package

The main package is special: it produces an executable program, not a library
```

---

### 3. How do you run a Go program using the command line?

```bash
go run filename.go
```

---

### 4. What is the role of the `func main()` in a Go application?

```go
It’s the first function that runs when your Go program starts

The Go runtime calls it automatically

When main() finishes, the program exits
```

---

### 5. Write a program that prints your name and your favorite programming language on two separate lines.

```go
package main

import "fmt"

func main() {
    fmt.Println("Chintamani")
    fmt.Println("Go")
}
```

---
