package main

import (
	"fmt"
	"os"

	"github.com/bigglezworthe/pratt_parser/src/lexer"
)

func main() {
    sourceBytes, _ := os.ReadFile("test2.lang")

    tokens, err := lexer.Tokenize(string(sourceBytes))
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, token := range tokens {
        token.Debug()
    }
}

