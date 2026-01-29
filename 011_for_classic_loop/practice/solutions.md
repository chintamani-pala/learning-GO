# 💡 Solutions: For Classic Loop

> [!NOTE]
> Implement your core logic in the code blocks below to master iteration concepts.

---

### 1. Write a for loop that prints numbers from 1 to 10.

```go
package main

import "fmt"

func main() {
    for i:=1; i<=10; i++ {
        fmt.Println(i)
    }
}
```

---

### 2. Create a program that calculates the sum of the first 100 positive integers.

```go
package main
import "fmt"

func main() {
    sum:=0
    for i:=1; i<=100; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

---

### 3. Write a for loop that decrementally prints numbers from 10 down to 1.

```go
package main
import "fmt"

func main() {
    for i:=10; i<=1; i-- {
        fmt.Println(i)
    }
}
```

---

### 4. How do you implement an "infinite loop" in Go?

```go
package main
import "fmt"

func main() {
    for true {
        fmt.Println("infinite")
    }
}
```

---

### 5. Write a program that prints only the even numbers between 1 and 20 using a classic for loop.

```go
package main
import "fmt"

func main() {
    for i:=1;i<=20; i++ {
        if i%2==0{
            fmt.Println(i)
        }
    }
}
```

---
