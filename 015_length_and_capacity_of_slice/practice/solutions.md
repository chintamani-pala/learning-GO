# 💡 Solutions: Length and Capacity of Slice

> [!NOTE]
> Implement your core logic in the code blocks below to master slice internals.

---

### 1. Create a slice with a length of 5 and a capacity of 10 using `make`.

```go
package main

import "fmt"

func main() {
    var intSlice []int = make([]int, 5, 10)
    fmt.Println(intSlice)
}
```

---

### 2. What does `len(s)` and `cap(s)` represent for a slice `s`?

```go
len(s) represents the number of elements in the slice.
cap(s) represents the total number of elements that the underlying array can hold.
```

---

### 3. Create a slice, append elements until its capacity is exceeded, and print the capacity before and after.

```go
package main

import "fmt"

func main() {
    var intSlice []int = make([]int, 5, 10)
    fmt.Println(intSlice)
    fmt.Println(cap(intSlice))
    intSlice = append(intSlice, 1,2,3,4,5,6,7,8,9,0)
    fmt.Println(intSlice)
    fmt.Println(cap(intSlice))
}
```

---

### 4. If you create a sub-slice `s2 := s[1:3]`, how are the length and capacity of `s2` determined based on `s`?

```go
For s2 := s[low:high], length is high - low, and capacity is cap(s) - low, since the sub-slice shares the underlying array starting from index low.
```

---

### 5. Write a program that demonstrates a "nil slice" and its length and capacity.

```go
package main

import "fmt"

func main() {
    var intSlice []int
    fmt.Println(intSlice)
    fmt.Println(len(intSlice))
    fmt.Println(cap(intSlice))
}
```

---
