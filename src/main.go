package main

import (
	"fmt"
	"os"

	"github.com/bigglezworthe/pratt_parser/src/lexer"
	"github.com/bigglezworthe/pratt_parser/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
    sourceBytes, _ := os.ReadFile("./examples/02.lang")

    tokens, err := lexer.Tokenize(string(sourceBytes))
    if err != nil {
        fmt.Println(err)
        return
    }

    //for _, token := range tokens {
    //    token.Debug()
    //}

    ast := parser.Parse(tokens)
    litter.Dump(ast)
}

