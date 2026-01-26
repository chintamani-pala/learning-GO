# 💡 Solutions: Constants

> [!NOTE]
> Implement your core logic in the code blocks below to master typed and untyped constants.

---

### 1. Declare a constant named `SecondsInMinute` and set it to 60.

```go
package main

import "fmt"

func main() {
    const SecondsInMinute int = 60
    fmt.Println(SecondsInMinute)
}
```

---

### 2. Can you change the value of a constant after it has been declared? Try it and see the compiler error.

```go
package main

import "fmt"

func main() {
    const SecondsInMinute int = 60
    fmt.Println(SecondsInMinute)
    SecondsInMinute =  100 // cannot assign to SecondsInMinute (declared const)
    fmt.Println(SecondsInMinute)

}
```

---

### 3. What is an "untyped constant" in Go?

```go
package main

import "fmt"

func main() {
    const SecondsInMinute = 60 //notype is given explicitly
    fmt.Println(SecondsInMinute)
}
```

---

### 4. Declare a group of constants using a `const` block.

```go
package main

import "fmt"
func main() {
    const (
        SecondsInMinute = 60
        MinutesInHour = 60
        HoursInDay = 24
    )

    fmt.Println(SecondsInMinute)
    fmt.Println(MinutesInHour)
    fmt.Println(HoursInDay)
}
```

---

### 5. Use `iota` to create an enumeration for the days of the week (starting from Sunday = 0).

```go
package main

import "fmt"

func main() {
    const (
        Sunday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )

    fmt.Println(Sunday)
    fmt.Println(Monday)
    fmt.Println(Tuesday)
    fmt.Println(Wednesday)
    fmt.Println(Thursday)
    fmt.Println(Friday)
    fmt.Println(Saturday)
}
```
🔹 What is `iota` in Go?

iota is a built-in identifier used inside const blocks
It represents successive untyped integer constants, starting from 0.


---
