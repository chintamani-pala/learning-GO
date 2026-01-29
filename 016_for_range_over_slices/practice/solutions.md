# 💡 Solutions: For Range Over Slices

> [!NOTE]
> Implement your core logic in the code blocks below to master iterating with `range`.

---

### 1. Write a program that uses `range` to iterate over a slice of integers and print each index and value.

```go
package main

func main() {
    var intSlice []int = []int{1,2,3,4,5}
    for index, value := range intSlice {
        fmt.Println(index, value)
    }
}
```

---

### 2. How can you iterate over a slice using `range` if you only need the values and not the indices?

```go
package main

func main() {
    var intSlice []int = []int{1,2,3,4,5}
    for _, value := range intSlice {
        fmt.Println(value)
    }
}
```

---

### 3. Write a program that finds the sum of all elements in a float slice using `range`.

```go
package main

func main() {
    var floatSlice []float64 = []float64{1.1,2.2,3.3,4.4,5.5}
    var sum float64 = 0
    for _, value := range floatSlice {
        sum += value
    }
    fmt.Println(sum)
}
```

---

### 4. Can you modify the original slice elements directly inside a `range` loop? Give an example or explain why/why not.

```go
package main

func main() {
    var intSlice []int = []int{1,2,3,4,5}
    for index, value := range intSlice {
        intSlice[index] = value * 2
    }
    fmt.Println(intSlice)
}
```

---

### 5. Use `range` to iterate over a string and print each character and its index.

```go
package main

func main() {
    var str string = "Hello"
    for index, value := range str {
        fmt.Println(index, string(value))
    }
}
```

---
