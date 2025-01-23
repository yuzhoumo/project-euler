package main

import (
    "fmt"
    "project-euler/golang/solutions"
)

type sol struct {
    id int
    f func() string
}

func main() {
    allSolutions := []sol{
        sol{1, solutions.E1},
        sol{2, solutions.E2},
        sol{3, solutions.E3},
        sol{4, solutions.E4},
        sol{5, solutions.E5},
        sol{6, solutions.E6},
    }

    for _, s := range allSolutions {
        fmt.Printf("%d: %s\n", s.id, s.f())
    }
}
