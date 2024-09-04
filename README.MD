# Code Challenge wc Tool

This is my implementation of the [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc/) challenge. It is written in GO and can be built using `make`.

The implementation supports command line file name inputs and also standard input. Available command line flags are:

- `-h` outputs text explaining how to use the tool
- `-c` outputs number of bytes if chosen
- `-l` outputs number of lines if chosen
- `-m` outputs number of characters if chosen
- `-w` outputs number of white space separated words if chosen

If no flags are chosen, `l`, `w` and `c` are enabled by default.

## Example usage

`ccwc test.txt test2.txt`

`cat test.txt | ccwc -l`
