package main

import (
    "github.com/en666ki/MJson/pkg/json_reader"
    "fmt"
)

func main() {
    var json_data, err = json_reader.Parse("/tmp/json.json")
    if (err != nil) {
        panic(err)
    }
    fmt.Println(json_data)
}
