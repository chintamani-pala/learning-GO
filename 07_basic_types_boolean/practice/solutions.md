# 💡 Solutions: Basic Types - Boolean

> [!NOTE]
> Implement your core logic in the code blocks below to master boolean logic and comparisons.

---

### 1. Declare a boolean variable `isGoFun` and assign it `true`. Print its value.

```go
package main

import "fmt"

func main() {
    var isGoFun bool = true
    fmt.Println(isGoFun)
}
```

---

### 2. Write an expression that checks if 10 is greater than 5 and store the result in a boolean.

```go
package main

import "fmt"

func main() {
    var isGreater bool = 10 > 5
    fmt.Println(isGreater)  
}
```

---

### 3. Use the logical AND (`&&`) operator to check if two conditions are both true.

```go
package main

import "fmt"

func main() {
    var isGreater bool = 10 > 5
    var isLess bool = 10 < 5
    fmt.Println(isGreater && isLess)
}
```

---

### 4. Use the logical NOT (`!`) operator to invert a boolean value.

```go
package main

import "fmt"

func main() {
    var isGreater bool = 10 > 5
    fmt.Println(!isGreater)
}
```

---

### 5. Create a program that prints the result of `(5 > 3) || (10 < 2)`. What is the expected output?

```go
package main

import "fmt"

func main() {
    fmt.Println((5 > 3) || (10 < 2)) // true
}
```

---
