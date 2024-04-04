# GOLOX

Golox is a general purpose programming language. It is a Go implementation of Lox (by Robert Nystrom) with some differences

## Getting started

### Running the REPL
To run the REPL, simply type

`go run .`

### Executing a file
To execute a file in the current directory use the -f flag
`go run. -f test.glx`

### Usage example: Print fibonacci numbers
```
fun fib(n) {
  if (n <= 1) return n;
  return fib(n - 2) + fib(n - 1);
}
for (var i = 0; i < 20; i = i + 1) {
  print fib(i);
}
```

Output
```
0
1
1
2
..
4181
```

## Debug Mode
Add the flag -d to enable debug mode
`go run . -d`


## Interpret a file
It looks for the file, passed with the -f flag, in project root
`go run . -f test.glx`
