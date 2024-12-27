# A. Beautiful Matrix

**Time limit per test:** 2 seconds  
**Memory limit per test:** 256 megabytes  

You are given a 5×5 matrix consisting of 24 zeroes and a single number one (`1`). Rows are indexed from 1 to 5 (top to bottom) and columns are indexed from 1 to 5 (left to right).  

You are allowed to perform the following operations on the matrix:  

1. **Swap two neighboring rows** (rows with indexes `i` and `i+1` where `1 ≤ i < 5`).  
2. **Swap two neighboring columns** (columns with indexes `j` and `j+1` where `1 ≤ j < 5`).  

A matrix is considered **beautiful** if the single `1` is positioned in the center of the matrix (at the intersection of the third row and the third column).  

Your task is to calculate the **minimum number of moves** required to make the matrix beautiful.

## Input
The input consists of five lines, each containing five integers:  
The `j-th` integer in the `i-th` line represents the element of the matrix at the intersection of row `i` and column `j`.  

It is guaranteed that the matrix contains 24 zeroes and a single number `1`.

## Output
Print a single integer — the minimum number of moves needed to make the matrix beautiful.

## Example

### Input

0 0 0 0 0

0 0 0 0 1

0 0 0 0 0

0 0 0 0 0

0 0 0 0 0

### Output

3
