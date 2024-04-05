package lexer

import (
	"fmt"
	"regexp"
)

type regexPattern struct {
    regex *regexp.Regexp
    handler regexHandler
}

type lexer struct {
    patterns []regexPattern
    Tokens []Token
    source string
    pos int
}

func (lex *lexer) advanceN(n int) {
    lex.pos += n
}

func (lex *lexer) push(t Token) {
    lex.Tokens = append(lex.Tokens, t)
}

func (lex *lexer) at() byte {
    return lex.source[lex.pos]   //treating string as array gives byte value 
}

func (lex *lexer) atEOF() bool {
    return lex.pos >= len(lex.source) 
}

func (lex *lexer) remainder() string {
    return lex.source[lex.pos:] 
}

func Tokenize(source string) ([]Token, error) {
    lex := createLexer(source)
    
    for !lex.atEOF() {
        matched := false 
        for _, pattern := range lex.patterns {
            loc := pattern.regex.FindStringIndex(lex.remainder())
            if loc != nil && loc[0] == 0 {
                pattern.handler(lex, pattern.regex)
                matched = true 
                break
            }
        }
        
        //should probably extend this
        if !matched {
            return nil, fmt.Errorf("Lexer::Error -> unrecognized token near %s", lex.remainder())
        }
    }

    lex.push(NewToken(EOF, "eof"))
    return lex.Tokens, nil
}

func createLexer(source string) *lexer {
    return &lexer{
        pos: 0,
        source: source,
        Tokens: make([]Token, 0),
        patterns: []regexPattern{
            //define matching patterns
            
            //ignore whitespace
            {regexp.MustCompile(`\s+`), skipHandler},

            //non-constants
            {regexp.MustCompile(`\/\/.*`), commentHandler},
            {regexp.MustCompile(`"[^"]*"`), stringHandler},
            {regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
            {regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},

            //default handlers (constants) 
            {regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
            {regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
            {regexp.MustCompile(`\{`), defaultHandler(OPEN_BRACE, "{")},
            {regexp.MustCompile(`\}`), defaultHandler(CLOSE_BRACE, "}")},
            {regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
            {regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
            {regexp.MustCompile(`==`), defaultHandler(EQUAL, "==")},
            {regexp.MustCompile(`!=`), defaultHandler(NOT_EQUAL, "!=")},
            {regexp.MustCompile(`<=`), defaultHandler(LESS_EQUAL, "<=")},
            {regexp.MustCompile(`>=`), defaultHandler(GREAT_EQUAL, ">=")},
            {regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
            {regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
            {regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
            {regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
            {regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
            {regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
            {regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
            {regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
            {regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
            {regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
            {regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
            {regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
            {regexp.MustCompile(`\*=`), defaultHandler(STAR_EQUALS, "*=")},
            {regexp.MustCompile(`/=`), defaultHandler(SLASH_EQUALS, "/=")},
            {regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
            {regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
            {regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
            {regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
            {regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
            {regexp.MustCompile(`\^`), defaultHandler(CARROT, "^")},
            
            //these have to come after all the compound matches
            {regexp.MustCompile(`=`), defaultHandler(ASSIGN, "=")},
            {regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
            {regexp.MustCompile(`<`), defaultHandler(LESSER, "<")},
            {regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
        },
    }
}

//handles a particular token 
type regexHandler func(lex *lexer, regex *regexp.Regexp)

func defaultHandler(kind TokenKind, value string) regexHandler { 
    return func(lex *lexer, regex *regexp.Regexp) {
        lex.advanceN(len(value)) 
        lex.push(NewToken(kind, value)) 
    } 
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
    match := regex.FindStringIndex(lex.remainder())
    lex.advanceN(match[1])
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
    match := regex.FindStringIndex(lex.remainder())
    stringLiteral := lex.remainder()[match[0] + 1 : match[1] - 1] //omit quotes 

    lex.push(NewToken(STRING, stringLiteral))
    lex.advanceN(len(stringLiteral) + 2)      //include omitted quotes 
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
    match := regex.FindString(lex.remainder())
    lex.push(NewToken(NUMBER, match))
    lex.advanceN(len(match))
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
    match := regex.FindString(lex.remainder())

    //define vars in conditional
    if kind, exists := reserved_lu[match]; exists {
        lex.push(NewToken(kind, match))
    } else {
        lex.push(NewToken(IDENTIFIER, match))
    }

    lex.advanceN(len(match))
}

func commentHandler(lex *lexer, regex *regexp.Regexp) {
    match := regex.FindStringIndex(lex.remainder())
    if match != nil {
        //ignore comment
        lex.advanceN(match[1])
    }
}


