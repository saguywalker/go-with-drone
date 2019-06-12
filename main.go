package main

import (
    "fmt"
    "time"
    "flag"
)

func main(){
    var commit = flag.String("commit", "", "Commit SHA1")
    fmt.Println("GO with CI")
    flag.Parse()
    fmt.Println("This file was from commit:",*commit)
    fmt.Printf("Result of %d**%d is %d \n", 2, 15, Pow(2,15))
    fmt.Printf("Result of %d**%d is %d \n", 3,6, Pow(3,7))
    fmt.Printf("Result of %d**%d is %d \n", 5,5, Pow(5,5))
    fmt.Printf("Result of %d**%d is %d \n", 7,7, Pow(7,7))
    fmt.Printf("Result of %d**%d is %d \n", 32, 4, Pow(32,4))
    fmt.Println(time.Now().Format("20060102150405"))
}

func Pow(x, y int) int{
    total := 1
    i := 0
    for i < y{
        total *= x
        i += 1
    }
    return total
}

