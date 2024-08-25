# N-Queens Problem

- backtracking with the N-Queens problem.

## Problem:

Place N queens on an N x N chessboard so that no two queens can attack each other.


## Backtracking Approach:

   1. Start with an empty board.

   2. Place a queen in the first row, and then move to the next row.

   3. In the next row, place a queen in a column where it’s not attacked by any previously placed queens.

 4. if no column is safe, backtrack to the previous row and move the queen to the next column.

5. Continue this process until either all queens are placed (solution found) or all columns in the first row are tried (no solution).

    ## Algorithm:
   ### 1.  Board Setup:

      +  We initialize an n x n chessboard with . representing empty spaces.

    ### 2. Backtracking Function:
     +   The backtrack function tries to place a queen in each column of the current row. If placing a queen is safe (checked by isSafe), it moves to the next row (backtrack(row + 1)).

    +    If a solution is found (all rows filled with queens), it records the solution. If no solution is possible, it backtracks by removing the last placed queen.

    ### 3. Safety Check (isSafe):
     +   Ensures no queens are in the same column or diagonal before placing a new queen.

   ### 4. Output:
        The solveNQueens function returns all possible solutions. Each solution is a different configuration of the board where all queens are safely placed.

## Benefits of Backtracking:

-  Optimal for Complex Problems: Effective for problems with large solution spaces where brute-force would be inefficient.

 -    Generality: Can be applied to a wide variety of problems, including puzzles, games, and optimization problems.

## Drawbacks:

- Efficiency: Backtracking can be slow for very large or complex problems because it explores many possibilities. Pruning and heuristics can help, but it’s not always the fastest approach.

-  Complexity: Implementing a backtracking solution can be complex, especially for problems with intricate constraints.

