# GOLOX

Golox is a general purpose programming language. It is a Go implementation of Lox (by Robert Nystrom) with some differences

## Getting started

### Running the REPL
To run the REPL, simply type
`go run .`

It can interpret equality expressions
`> print 1==1;
true`
`> print 'test' == 'test';
true`

And comparison
`> print 5 > 10;
false`


## Debug Mode
Add the flag -d to enable debug mode
`go run . -d`


## Interpret a file
It looks for the file, passed with the -f flag, in project root
`go run . -f test.glx`
