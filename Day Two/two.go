package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("File reading error:", err)
        return
    }
    defer file.Close()

    type password struct {
        Min, Max int
        Password string
        Policy string
    }

    var passwords = make([]password, 0)
    re := regexp.MustCompile(`(\d*)-(\d*)\s(\w):\s(\w*)`)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result := re.FindStringSubmatch(scanner.Text())
        min, err := strconv.Atoi(result[1])
        if err != nil {
            fmt.Println("Integer conversion error:", err)
        }
        max, err := strconv.Atoi(result[2])
        if err != nil {
            fmt.Println("Integer conversion error:", err)
        }
        passwords = append(passwords, password{
            Min: min,
            Max: max,
            Policy: result[3],
            Password: result[4],
        })

    }
    good := 0
    for _, a := range passwords {
        count := strings.Count(a.Password, a.Policy)
        if count >= a.Min && count <= a.Max {
            good++
        }
    }
    fmt.Println("Total good passwords:", good)
    good1 := 0
    for _, a := range passwords {
        password_good := false
        test1 := a.Password[a.Min-1:a.Min] == a.Policy
        test2 := a.Password[a.Max-1:a.Max] == a.Policy
        if test1 && !test2 {
            password_good = true
        } else if !test1 && test2 {
            password_good = true
        }
        if password_good {
            good1++
        }
    }
    fmt.Println("Total actually good1 passwords:", good1)
}
