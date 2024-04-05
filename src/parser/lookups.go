package parser

import (
	"github.com/bigglezworthe/pratt_parser/src/ast"
	"github.com/bigglezworthe/pratt_parser/src/lexer"
)

type bindingPower int 

//higher power = tighter binding (inverse precedence) 
const (
    default_bp bindingPower = iota
    comma
    assignment
    logical
    relational
    additive
    multiplicative
    unary
    call
    member
    primary
)

//NUD = Null Denoted (like !, no LHS expected)
//LED = Left Denoted (like +, LHS expected)
type stmtHandler func(p *parser) ast.Stmt
type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp bindingPower) ast.Expr

type stmtLookup map[lexer.TokenKind]stmtHandler
type nudLookup map[lexer.TokenKind]nudHandler
type ledLookup map[lexer.TokenKind]ledHandler
type bpLookup map[lexer.TokenKind]bindingPower

//global lookup tables
var bp_lu = bpLookup{}
var nud_lu = nudLookup{}
var led_lu = ledLookup{}
var stmt_lu = stmtLookup{}

// Helper Methods
func led(kind lexer.TokenKind, bp bindingPower, ledFn ledHandler) {
    bp_lu[kind] = bp
    led_lu[kind] = ledFn
}

func nud(kind lexer.TokenKind, bp bindingPower, nudFn nudHandler) {
    bp_lu[kind] = bp
    nud_lu[kind] = nudFn
}

func stmt(kind lexer.TokenKind, stmtFn stmtHandler) {
    bp_lu[kind] = default_bp
    stmt_lu[kind] = stmtFn
}

func createTokenLookups() {

    //Logical 
    led(lexer.AND, logical, parse_binary_expr)
    led(lexer.OR, logical, parse_binary_expr)
    led(lexer.DOT_DOT, logical, parse_binary_expr)

    //Relational
    led(lexer.LESSER, relational, parse_binary_expr)
    led(lexer.GREATER, relational, parse_binary_expr)
    led(lexer.LESS_EQUAL, relational, parse_binary_expr)
    led(lexer.GREAT_EQUAL, relational, parse_binary_expr)
    led(lexer.EQUAL, relational, parse_binary_expr)
    led(lexer.NOT_EQUAL, relational, parse_binary_expr)

    //Additive & Multiplicative 
    led(lexer.PLUS, additive, parse_binary_expr)
    led(lexer.MINUS, additive, parse_binary_expr)
    led(lexer.STAR, additive, parse_binary_expr)
    led(lexer.SLASH, additive, parse_binary_expr)
    led(lexer.PERCENT, additive, parse_binary_expr)
    led(lexer.CARROT, additive, parse_binary_expr)

    //Literals & Symbols
    nud(lexer.NUMBER, primary, parse_primary_expr)
    nud(lexer.STRING, primary, parse_primary_expr)
    nud(lexer.IDENTIFIER, primary, parse_primary_expr)
    
    //Statements
    stmt(lexer.CONST, parse_var_decl_stmt)
    stmt(lexer.LET, parse_var_decl_stmt)

}
