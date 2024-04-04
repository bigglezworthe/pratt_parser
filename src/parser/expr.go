package parser

import (
	"fmt"
	"strconv"

	"github.com/bigglezworthe/pratt_parser/src/ast"
	"github.com/bigglezworthe/pratt_parser/src/lexer"
)

func parse_expr(p *parser, bp bindingPower) ast.Expr {
    //First parse the NUD 
    tokenKind := p.currentTokenKind()
    nud_fn, exists := nud_lu[tokenKind]
    if !exists {
        panic(fmt.Sprintf("NUD Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
    }

    //While LED and current BP < current token BP, parse LHS (unrolls recursion)
    left := nud_fn(p)
    for bp_lu[p.currentTokenKind()] > bp {
        tokenKind = p.currentTokenKind()
        led_fn, exists := led_lu[tokenKind]
        if ; !exists {
            panic(fmt.Sprintf("LED Handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
        }

        left = led_fn(p, left, bp)
    }

    return left
}


func parse_primary_expr(p *parser) ast.Expr {
    switch p.currentTokenKind() {
    case lexer.NUMBER:
        number, _ := strconv.ParseFloat(p.advance().Value, 64)
        return &ast.NumberExpr{
            Value: number,
        }
    case lexer.STRING:
        return &ast.StringExpr{
            Value: p.advance().Value,
        }
    case lexer.IDENTIFIER:
        return &ast.SymbolExpr{
            Value: p.advance().Value,
        }
    default:
        panic(fmt.Sprintf("Cannot create primary_expression from %s\n", lexer.TokenKindString(p.currentTokenKind())))
    }
}

func parse_binary_expr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
    operatorToken := p.advance()
    right := parse_expr(p, bp)

    return ast.BinaryExpr{
        Left: left, 
        Operator: operatorToken,
        Right: right,
    }
    
}
