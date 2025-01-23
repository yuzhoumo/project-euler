package main

import (
    "fmt"
    "github.com/yuzhoumo/project-euler/golang/solutions"
)

func main() {
    sols := map[string]func()string{
        "1": solutions.E1,
    }

    for k, v := range sols {
        fmt.Printf("%s: %s\n", k, v())
    }
}
