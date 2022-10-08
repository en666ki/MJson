package json_reader

import (
    "encoding/json"
    "os"
    "errors"
)

func Parse(input string) (string, error) {
    var data, err = os.ReadFile(input)
    if err != nil {
        return "", err
    }
    if !json.Valid(data) {
        return "", errors.New("Invalid json")
    }
    return string(data[:]), nil
}


