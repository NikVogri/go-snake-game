# Play Snake Game inside your terminal

A simple implementation of the famous Snake game for your terminal.

Your job is to eat as much fruit (marked as F) as you can without leaving the border or eating your own body!

You can specify the board size by changing the `BOARD_SIZE` variable (located inside the `main.go` file).

## Rules

After every 5 fruit you eat, the snake will start to move 5% faster.

Leaving the border or eating your own body will end the game.

## Prerequisites

-   GoLang installed on your machine

## Starting the game

1. Build the game:

```
$ go build .
```

2. Start game:

```
$ ./go-snake-game
```

## Controls

-   use **arrow keys** (UP, DOWN, LEFT, RIGHT) to move the snake
-   **esc key** to end the game
