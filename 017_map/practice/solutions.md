# 💡 Solutions: Maps

> [!NOTE]
> Implement your core logic in the code blocks below to master key-value pairs in Go.

---

### 1. Declare and initialize a map where keys are strings (country names) and values are strings (capitals).

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    fmt.Println(countryName)
}
```

---

### 2. How do you add a new key-value pair to an existing map?

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    countryName["Japan"] = "Tokyo"
    fmt.Println(countryName)
}
```

---

### 3. Write a program that checks if a specific key exists in a map and prints its value if it does.

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    if value, ok:=countryName["India"]; ok {
        fmt.Println(value)
    }
}
```

---

### 4. How do you delete a key from a map?

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    delete(countryName, "India")
    fmt.Println(countryName)
}
```

---

### 5. Write a program that uses `range` to iterate over a map and print all keys and values.

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    for key, value := range countryName {
        fmt.Println(key, value)
    }
}
```

---
