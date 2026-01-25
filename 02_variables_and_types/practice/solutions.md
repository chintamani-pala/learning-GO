# 💡 Solutions: Variables and Types

> [!NOTE]
> Implement your core logic in the code blocks below to master variable declarations and types.

---

### 1. Declare an integer variable `age` and assign it the value 25, then print it.

```go
package main

import "fmt"

func main() {
    //age := 25 //way 1
    // var age int = 25 //way 2

    var age int //way 3
    age = 25
    fmt.Println(age)
}
```

---

### 2. Create a string variable `city` and an integer variable `zipCode`. Assign values and print them together.

```go
package main 

import "fmt"

func main() {
    city := "New York"
    zipCode := 10001
    fmt.Println(city, zipCode)
}
```

---

### 3. What is the difference between `var x int` and `x := 10`?

```go
In var x int, x is declared but not initialized (zero value) 
In x := 10, x is declared and initialized (type inference)
```

---

### 4. Declare three variables of different types (bool, float64, string) in a single `var` block.

```go
package main

import "fmt"

func main() {
    var (
        isStudent  = true
        gpa  = 3.8
        name = "John Doe"
    )
    fmt.Println(isStudent, gpa, name)
}
```

---

### 5. Write a program that declares a variable without an initial value and prints its "zero value".

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Println(age) // Output: 0
}
```

---
