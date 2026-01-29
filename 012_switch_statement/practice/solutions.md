# 💡 Solutions: Switch Statement

> [!NOTE]
> Implement your core logic in the code blocks below to master multi-way branching.

---

### 1. Write a switch statement that prints the name of the day based on an integer (1=Monday, 2=Tuesday, etc.).

```go
package main

import "fmt"

func main() {
	var day int

	fmt.Print("Enter day number: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day number")
	}
}

```

---

### 2. How is a Go `switch` different from `switch` in languages like C++ or Java regarding `break` statements?

```go
break is automatic

Each case stops after executing

No accidental fall-through

```

---

### 3. Create a switch statement with no expression (switch true) that checks multiple boolean conditions.

```go
package main

import "fmt"

func main() {
	var day int = 1
	switch {
        case day < 1:
            fmt.Println("Invalid")
        case day <= 5:
            fmt.Println("Weekday")
        default:
            fmt.Println("Weekend")
	}
}
```

---

### 4. Use the `fallthrough` keyword in a switch statement and explain what it does.

```go
package main

import "fmt"

func main() {
	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}
}

```
 - Forces execution to continue to the next case

 - Ignores the next case’s condition

 - Executes exactly one additional case

 - Must be the last statement inside a case

---

### 5. Write a switch statement that categorizes a grade (A, B, C, D, F) based on a score range.

```go
package main

import "fmt"

func main() {
	score := 85

	switch {
	case score >= 90 && score <= 100:
		fmt.Println("Grade A")
	case score >= 80 && score < 90:
		fmt.Println("Grade B")
	case score >= 70 && score < 80:
		fmt.Println("Grade C")
	case score >= 60 && score < 70:
		fmt.Println("Grade D")
	case score >= 0 && score < 60:
		fmt.Println("Grade F")
	default:
		fmt.Println("Invalid score")
	}
}

```

---
