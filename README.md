# Go Basics Learning

A collection of Go programs and notes covering fundamental concepts of the Go programming language.

## Project Structure

This repository is organized into numbered directories, each focusing on a specific topic:

- **01_setup_first_program**: Setting up the environment and writing the first "Hello World" program.
- **02_variables_and_types**: Introduction to variables and basic types.
- **03_packages_imports**: Understanding Go packages and imports.
- **04_var_vs_short_declare**: Difference between standard `var` declaration and short declaration `:=`.
- **05_basic_types_string**: Working with strings.
- **06_basic_types_int_float**: Integer and floating-point number operations.
- **07_basic_types_boolean**: Boolean values and logic.
- **08_constants**: Defining and using constants.
- **09_if_else**: Conditional logic with `if` and `else`.
- **010_if_with_short_statement**: Using short statements within `if` blocks.
- **011_for_classic_loop**: The classic `for` loop syntax.
- **012_switch_statement**: Switch case control flow.
- **013_arrays**: Working with fixed-size arrays.
- **014_slice_literals**: Introduction to slices and slice literals.
- **015_length_and_capacity_of_slice**: Understanding slice length and capacity.
- **016_for_range_over_slices**: Iterating over slices using the `range` keyword.
- **017_map**: Working with key-value pairs using maps.
- **018_read_value_ok**: Checking for existence of keys in maps using the "comma ok" idiom.
- **019_functions**: Defining and using functions.
- **020_function_named_return_values**: Using named return values in functions.
- **021_variadic_functions**: Functions that accept a variable number of arguments.

## Getting Started

To run any of the examples, navigate to the directory and run the `main.go` file:

```bash
cd 01_setup_first_program
go run main.go
```

## Requirements

- Go 1.25.5 or higher
