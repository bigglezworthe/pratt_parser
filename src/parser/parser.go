package parser

import (
	"github.com/bigglezworthe/pratt_parser/src/lexer"
	"github.com/bizzlezworthe/pratt_parser/src/ast"
)

type parser struct {
    //error []error <-- DA FUTURE 
    tokens []lexer.Token
    pos int
}

func Parse(tokens []lexer.Token) ast.BlockStmt
