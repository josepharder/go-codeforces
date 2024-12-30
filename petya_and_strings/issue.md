# A. Petya and Strings

**Time limit per test:** 2 seconds  
**Memory limit per test:** 256 megabytes  

Little Petya loves presents. For his birthday, his mum bought him two strings of the same size, consisting of uppercase and lowercase Latin letters. Petya wants to compare the two strings **lexicographically**, but the comparison is case-insensitive (uppercase and lowercase letters are treated as equivalent). Help Petya perform the comparison.

## Input
- The first and second lines each contain one string.  
- The strings' lengths range from 1 to 100 inclusive.  
- The strings are guaranteed to be of the same length and consist only of uppercase and lowercase Latin letters.

## Output
Print:
- `-1` if the first string is lexicographically less than the second one.  
- `1` if the first string is lexicographically greater than the second one.  
- `0` if the strings are equal.

**Note:** The comparison is case-insensitive.

## Example

### Input

abcdefg
AbCdEfF

### Output

1

### Input

abs

Abz

### Output

-1
