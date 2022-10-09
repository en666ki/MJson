package json_reader

import (
    "encoding/json"
    "os"
)

func Parse(input string) (bool, error) {
    var data, err = os.ReadFile(input)
    if err != nil {
        return false, err
    }
    if !json.Valid(data) {
        return false, nil
    }
    return true, nil
}


