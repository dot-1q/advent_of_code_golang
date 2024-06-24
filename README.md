# Advent of Code in Go

![image](aoc_go.png "aoc")

The various Advent of Code challenges solved in golang.
This was done as a way to learn and get experienced with Go, so the first implementations might not be very idiomatic.

| Year | Done        |
|------|-------------|
| 2015 | Yes         |
| 2016 | Yes*        |
| 2017 | Yes         |
| 2018 | In Progress+|

<sub><sup>\* I did not do day11. Might revisit in the future</sup></sub>

<sub><sup>\+ I did not do day15. Its convoluted just for the sake of being convoluted. Not in the mood to be burned out</sup></sub>

<sub><sup>\I will revisit these days later, when I'm more open to them</sup></sub>

# Notable Challenges
[Day 22 of 2015](2015/day22/twenty-two.go) Implemented Dijkstra's algorithm with a Priority Queue with respect to each game state and its mana spent.

[Day 10 of 2016](2016/day10/ten.go) I'm always happy to make recursion work lol. Plus, browsed reddit and it also wasn't the default solution implemented by the other people.

[Day 11 of 2017](2017/day11/eleven.go) Made a very concise solution for a new grid system (Hexagonal). Made me learn about it. [Hex Grid Information](https://www.redblobgames.com/grids/hexagons)

[Day 21 of 2017](2017/day21/twenty-one.go) Performed various matrix transformations (rotate and flip). ChatGPT helped me write the Merging algorithm for the matrices.

[Day 24 of 2017](2017/day24/twenty-four.go) Depth First Search to create all the possible paths from one node to the next valid ones. Very concise solution.

[Day 16 of 2018](2018/day16/sixteen.go) First time creating an array of functions to callback later in the runtime of the program. Also fmt.Sscanf usage to revisit later.
