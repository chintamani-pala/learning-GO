# 💡 Solutions: Read Value OK

> [!NOTE]
> Implement your core logic in the code blocks below to master the map access pattern.

---

### 1. Explain what the "comma ok" idiom is when reading from a map.

```go
It checks is actually that value exist or not if not then also the value will be 0 but the "ok" will be false
```

---

### 2. Write a code snippet that attempts to read a value from a map and prints "Found" or "Not Found" based on the boolean result.

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
        fmt.Println("Found", value)
    } else {
        fmt.Println("Not Found")
    }
}
```

---

### 3. What is the value of the first variable in `val, ok := myMap["key"]` if the key does not exist?

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    if value, ok:=countryName["Japan"]; ok {
        fmt.Println("Found", value)
    } else {
        fmt.Println("Not Found")
    }
}
```

---

### 4. Why is it safer to use the `ok` check rather than just checking if the returned value is the zero value?

```go
Because if we use ok then we will know is that value is actually present or not.
```

---

### 5. Write a program that tries to access a key in a map of integers; if the key doesn't exist, it should initialize it with a default value.

```go
package main

import "fmt"

func main() {
    var countryName := map[string]string {
        "India":"New Delhi",
        "USA":"Washington D.C.",
        "China":"Beijing",
    }
    if value, ok:=countryName["Japan"]; ok {
        fmt.Println("Found", value)
    } else {
        fmt.Println("Not Found", "Default value", value)
    }
}
```

---
