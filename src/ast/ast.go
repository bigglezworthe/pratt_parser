// AST = Abstract Syntax Tree
package ast

type Stmt interface {
    stmt()
}

type Expr interface {
    expr()
}
