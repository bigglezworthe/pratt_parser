package lexer

import "fmt"

type Token struct {
    Kind TokenKind
    Value string
}

func (t Token) isOfKind(expectedTokens ...TokenKind) bool {
    for _, expected := range expectedTokens {
        if expected == t.Kind{
            return true
        }
    }

    return false
}

func (t Token) Debug() {
    if t.isOfKind(NUMBER, STRING, IDENTIFIER) {
        fmt.Printf("%s (%s)\n", TokenKindString(t.Kind), t.Value)
    } else { 
        fmt.Printf("%s ()\n", TokenKindString(t.Kind))
    }
}

func NewToken(kind TokenKind, value string) Token {
    return Token{
        Kind: kind, 
        Value: value,
    }
}
