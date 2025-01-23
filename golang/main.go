package main

import (
    "fmt"
    "project-euler/golang/solutions"
)

func main() {
    sols := map[string]func()string{
        "1": solutions.E1,
        "2": solutions.E2,
    }

    for k, v := range sols {
        fmt.Printf("%s: %s\n", k, v())
    }
}
