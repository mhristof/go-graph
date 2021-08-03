# go-graph

[![Go Report Card](https://goreportcard.com/badge/github.com/mhristof/go-graph)](https://goreportcard.com/report/github.com/mhristof/go-graph)

Graph implementation in golang

## DependencyMap

A dependency map as described in [this](https://www.youtube.com/watch?v=ddTC4Zovtbc)

Given this dep graph,

```
      A            B
       -\         /|
         -\     /- |
           -C --   |
            /      |
          /-       |
        /-         |
      E-           D
     --\         /-|
   -/   -\   /---  |
 -/       ---      |
 H        F        G
```

the `Sort("h")` would return `h, g, f, e, c, a`
and the `SortAll("h")` would return `h, g, f, e, c, a, d, b`
