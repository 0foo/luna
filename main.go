
package main

import (
    "fmt"
    "os"

    "seedutil"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go "faker_function|param1=value1|param2=value2"")
        return
    }

    input := os.Args[1]
    result, err := seedutil.CallFakerDynamic(input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(result)
}
