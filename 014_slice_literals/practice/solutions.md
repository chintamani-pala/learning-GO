# 💡 Solutions: Slice Literals

> [!NOTE]
> Implement your core logic in the code blocks below to master dynamic sequences.

---

### 1. Create a slice of strings containing your favorite colors.

```go
package main

import "fmt"

func main() {
    var colors []string = []string{"red", "green", "blue"}
    fmt.Println(colors)
}
```

---

### 2. What is the main difference between an array and a slice in Go?

```go
In Go, arrays are fixed-size sequences, while slices are dynamic sequences that can grow or shrink.
```

---

### 3. Initialize a slice of integers using the `[]int{...}` syntax.

```go
package main

import "fmt"

func main() {
    var intSlice []int = []int{1,2,3,4,5}
    fmt.Println(intSlice)
}
```

---

### 4. How do you add a new element to an existing slice?

```go
package main

import "fmt"

func main() {
    var intSlice []int = []int{1,2,3,4,5}
    intSlice = append(intSlice, 6)
    fmt.Println(intSlice)
}
```

---

### 5. Create a slice of structs (e.g., a `Person` struct with name and age) and initialize it with two items.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func main() {
    var people []Person = []Person{
        Person{Name: "John", Age: 30},
        Person{Name: "Jane", Age: 25},
    }
    fmt.Println(people)
}
```

---
