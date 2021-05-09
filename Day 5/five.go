package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("File reading error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result := scanner.Text()
        fmt.Println(result)
    }
}
