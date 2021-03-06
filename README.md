# gophercises

My solution to Jon Calhoun's [gophercises](https://github.com/gophercises/)

## ex1: Quiz Game

A timed quiz game that asks you basic questions and gives you the score after you finish or the quiz times out (whichever comes first).
Build the program using `go build quiz.go` <br>
Run using :
`./quiz -file <filename> -time <timelimit> [-shuf]`

The filename argument contains the questions and answers in a headless CSV format. <br>
The time limit argument is the timer length in seconds. <br>
The shuf argument asks the user whether the questions are to be shuffled on each run of the quiz <br>

## ex2: URL Redirector

A redirector that uses the `http` package to implement a simple URL Redirector. The program reads input from a JSON or YAML file, and 
starts a server that redirects when the path is entered in the browser. 
Build the program using `cd ex2/main; go build main.go` <br>
Run using :
`./main -f <filename>`

## ex3: Choose your own adventure

Implement a simple "Choose your own adventure" game on the browser using HTTP templates. The program reads input from a JSON or YAML file whose
format is described in the help section of the program. 
Build the program using `cd ex3/; go build` <br>
Run using :
`./cyoa -f <filename>`