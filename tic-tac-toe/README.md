
# Tic-Tac-Toe

Tic-tac-toe is played on a 3 by 3 board.  One player is crosses (`x`) the other (`o`).  Players take it in turn, to place a `x` or `o` on the board.  If a player gets 3 in a row either diagonally, vertically or horizontally they win.  If neither player achieves this when the grid is filled up then it is a draw.

Write a function that takes the final board as an array and gives the result of who won.

For example the following board:

```
X O X
O X O
O X X

```

Would be given from top to bottom left to right in a single array as `"x", "o", "x", "o", "x", "o", "o", "x", "x"`, your function should return `"x"` as `x` won.

The function signature should be:

```
func CalculateWinner(b []rune) string
```

You should return:
    - `x` if the player with x won
    - `o` if the player with o won
    - `-` if it was a draw

The func is already defined, the challenge is to make all of tests pass with the fewest characters possible for the implmentation.  That is we will take the characters between the curly braces of the implementation.  We will use the https://charactercounttool.com/ to count the number of characters excluding spaces in your answer file.