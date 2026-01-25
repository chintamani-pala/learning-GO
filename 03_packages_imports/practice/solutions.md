# 💡 Solutions: Packages and Imports

> [!NOTE]
> Implement your core logic in the code blocks below to master standard library usage.

---

### 1. Import the `fmt` and `math` packages and print the value of `math.Pi`.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}
```

---

### 2. Write a program that uses `fmt.Printf` to format a string with a placeholder.

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello, %s!", "Go Developer")
}
```

---

### 3. How do you import multiple packages in a single `import` block?

```go
import (
    "fmt"
    "math"
)
```

---

### 4. Use the `rand` package (from `math/rand`) to generate a random integer between 1 and 100.

```go
package main

import (
    "time"
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    //⚠️ Important runtime detail (very common Go gotcha)

    //Without seeding the random number generator, Go produces the same “random” sequence every time you run the program.
    fmt.Println(rand.Intn(100))
}
```

---

### 5. Create a program that uses the `time` package to print the current local time.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println(time.Now().Local())
}
```

---
