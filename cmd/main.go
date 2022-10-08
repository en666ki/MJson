package main

import (
    "github.com/en666ki/MJson/pkg/json_reader"
    "fmt"
    "os"
)

func main() {
    if (len(os.Args) != 2) {
        fmt.Println("usage:\n\tMJson JSON_FILE")
        return
    }
    var input_file = os.Args[1]
    var json_data, err = json_reader.Parse(input_file)
    if (err != nil) {
        panic(err)
    }
    fmt.Println(json_data)
}
