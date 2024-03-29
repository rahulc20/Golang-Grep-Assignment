package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: grep <pattern> [file]")
        return
    }

    pattern := os.Args[1]
    var input *os.File

    if len(os.Args) > 2 {
        file := os.Args[2]
        f, err := os.Open(file)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
            return
        }
        defer f.Close()
        input = f
    } else {
        input = os.Stdin
    }

    scanner := bufio.NewScanner(input)
    regex := regexp.MustCompile(pattern)

    for scanner.Scan() {
        line := scanner.Text()
        if regex.MatchString(line) {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
    }
}
