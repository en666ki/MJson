package json_reader

import (
    "encoding/json"
    "log"
    "fmt"
    "flag"
)

func main() {
    var input = flag.String("input", "", "path to input file")
    var output = flag.String("output", "", "path to output file")
    flag.Parse()
    fmt.Println(input, output)
}

