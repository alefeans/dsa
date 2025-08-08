# Quiz Game

[Quiz Game](https://gophercises.com/exercises/quiz) implementation.

## Build

```
go build .
```

## Usage 

It's possible to use your own questions file (like [problems.csv](problems.csv)) and set a time limit to complete the Quiz:

```
Usage of ./quiz_game:
  -file string
        CSV file with the quiz questions (default "problems.csv")
  -time int
        Time limit to complete the Quiz (default 30)
```

## Running

```
./quiz_game
Problem: 5+5 = 10
Problem: 1+1 = 2
Problem: 8+3 = 11
Problem: 1+2 = 4
You scored 3 out of 12.
```