# 🐒 Monkey Programming Language

A complete implementation of the Monkey programming language from the book "Writing An Interpreter In Go" by Thorsten Ball, with additional features and improvements.

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/monkey-lang)](https://goreportcard.com/report/github.com/yourusername/monkey-lang)
[![GoDoc](https://godoc.org/github.com/yourusername/monkey-lang?status.svg)](https://godoc.org/github.com/yourusername/monkey-lang)

## 📖 Overview

Monkey is a programming language designed to learn about interpreter and compiler implementation. It features:

- C-like syntax
- Variable bindings
- Integer and boolean data types
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

## 🚀 Quick Start

### Prerequisites
- Go 1.16 or higher

### Installation
```bash
# Clone the repository
git clone https://github.com/yourusername/monkey-lang.git

# Navigate to the project directory
cd monkey-lang

# Build the project
go build -o monkey cmd/monkey/main.go

# Run the REPL
./monkey
```

## 💻 Usage

### REPL (Interactive Mode)
```bash
$ ./monkey
This is the Monkey programming language!
Feel free to type in commands
>> let x = 5;
>> x + 10;
15
>> 
```

### Running Monkey Files
```bash
$ ./monkey run examples/fibonacci.monkey
```

## 🔍 Language Features

### Variables
```monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

### Functions
```monkey
// Function definition
let add = fn(a, b) { return a + b; };

// Higher-order functions
let twice = fn(f, x) {
    return f(f(x));
};

// Closures
let newAdder = fn(x) {
    return fn(y) { return x + y };
};
```

### Arrays
```monkey
let myArray = [1, 2, 3, 4, 5];
let first = myArray[0];
```

### Hash Maps
```monkey
let myHash = {"name": "Monkey", "age": 1};
let name = myHash["name"];
```

## 🏗️ Project Structure

```
monkey-lang/
├── ast/            # Abstract Syntax Tree definitions
├── evaluator/      # Expression evaluation logic
├── lexer/          # Lexical analysis
├── object/         # Object system
├── parser/         # Parsing logic
├── repl/           # Read-Eval-Print Loop
├── token/          # Token definitions
├── cmd/
│   └── monkey/     # Main executable
├── examples/       # Example Monkey programs
└── tests/          # Test suite
```

## 🧪 Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./lexer
go test ./parser
go test ./evaluator
```

## 🔧 Extended Features

Beyond the book's implementation, this version includes:

1. **Standard Library**
   - Array methods (map, reduce, filter)
   - String manipulation functions
   - Math operations

2. **Additional Data Types**
   - Float numbers
   - Null type
   - Regular expressions

3. **Error Handling**
   - Try-catch mechanism
   - Stack traces
   - Better error messages

## 📝 Examples

### Fibonacci Sequence
```monkey
let fibonacci = fn(x) {
    if (x < 2) { return x; }
    return fibonacci(x - 1) + fibonacci(x - 2);
};

fibonacci(10);
```

### Map Function
```monkey
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        }
        return iter(rest(arr), push(accumulated, f(first(arr))));
    };
    return iter(arr, []);
};

let double = fn(x) { return x * 2 };
map([1, 2, 3, 4], double);
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📚 Resources

- [Writing An Interpreter In Go](https://interpreterbook.com/) - The original book
- [Writing A Compiler In Go](https://compilerbook.com/) - The sequel
- [Monkey Language Specification](./docs/SPEC.md)
- [Contributing Guidelines](./CONTRIBUTING.md)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thorsten Ball for writing the excellent books
- The Go team for creating a great language
- All contributors to this project

---
Made with ❤️ by Akshay
