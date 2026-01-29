# 💡 Solutions: Arrays

> [!NOTE]
> Implement your core logic in the code blocks below to master fixed-size data sequences.

---

### 1. Declare an array of 5 integers and initialize it with values.

```go
package main

import "fmt"

func main() {
    var intArray [5]int = [5]int{1,2,3,4,5}
    fmt.Println(intArray)
}
```

---

### 2. How do you find the length of an array in Go?

```go
package main

import "fmt"

func main() {
    var intArray [5]int = [5]int{1,2,3,4,5}
    fmt.Println(len(intArray))
}
```

---

### 3. What happens if you try to assign a value to an index that is out of bounds in an array?

```go
It will give array index out of bounds error
```

---

### 4. Write a program that finds the maximum value in an array of 10 integers.

```go
package main

import "fmt"

func main() {
    var intArray [5]int = [5]int{1,2,3,4,5,6,7,8,9,10}
    maximum := intArray[0]
    for i := 1; i < len(intArray); i++ {
        maximum = max(maximum, intArray[i])
    }
    fmt.Println(maximum)
}
```

---

### 5. Are arrays in Go passed by value or by reference? Explain the implication.

```go
In Go, arrays are passed by value, not by reference.
When you pass an array to a function, a copy of the entire array is made. The function receives this copy, so changes inside the function do not affect the original array.
```

---
