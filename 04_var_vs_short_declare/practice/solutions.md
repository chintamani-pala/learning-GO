# 💡 Solutions: Var vs Short Declare

> [!NOTE]
> Implement your core logic in the code blocks below to master Go declaration syntax.

---

### 1. When is it mandatory to use `var` instead of short declaration `:=`?

| Scenario                      | `var`      | `:=`          |
| ----------------------------- | ---------- | ------------- |
| Outside function              | ✅ required | ❌ not allowed |
| No initial value              | ✅          | ❌             |
| Explicit type needed          | ✅          | ❌             |
| Inside function, new variable | ✅          | ✅             |
| Reassignment only             | ✅          | ❌             |


---

### 2. Rewrite `var score int = 100` using the short declaration operator.

```go
// var score int = 100
score := 100
```

---

### 3. Can you use `:=` outside of a function body? Explain why or why not.


**No, you cannot use `:=` outside a function body.**

**Why?**

* `:=` works **only inside functions**
* Outside functions (package level), Go allows **only `var`, `const`, `type`, `func`**


---

### 4. Declare two variables `a` and `b` using a single `:=` statement.

```go
a, b := 10, 20
```

---

### 5. What happens if you try to use `:=` to reassign a value to an already declared variable in the same scope?

```go
Error: no new variables on left side of :=
```

---
