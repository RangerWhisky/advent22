# advent22
Advent of Code 22 solutions/attempts, written in Go for personal use/enjoyment

Intent/Focus
============

Intent is to learn Golang while practising good clean code development techniques.  TDD and legibility are key, with many commits to git as failing tests are written and the code written to match the behaviour.


About the code
==============

Main calls into each day, with separate methods for parts one and two\

Puzzle input is stored within each day, filepaths passed in from main.go for reading by the function.  This means that these functions can be tested using small/targeted test files before running against the main puzzle input

There's a bit of duplicated code between each of the days - the flow is emerging as "read a file, pass it to a function, sum the outputs" so there's scope for writing an intermediary that does this repeatedly with a function pointer to save rewrite
