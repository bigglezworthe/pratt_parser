package lexer 

//Go doesn't have enumerations, so we have to do this
type TokenKind int
const (
    EOF TokenKind = iota  //iota starts autonumbering 

    NULL
    TRUE
    FALSE
    STRING 
    NUMBER  
    IDENTIFIER

    //groupings 
    OPEN_BRACKET
    CLOSE_BRACKET
    OPEN_BRACE
    CLOSE_BRACE
    OPEN_PAREN
    CLOSE_PAREN

    //compare 
    EQUAL  
    ASSIGN
    NOT_EQUAL  
    GREATER 
    LESSER  
    GREAT_EQUAL 
    LESS_EQUAL 

    //logical
    NOT
    AND
    OR

    //math
    PLUS  
    MINUS  
    STAR 
    SLASH 
    CARROT
    PERCENT

    //shorthand 
    PLUS_PLUS
    MINUS_MINUS
    PLUS_EQUALS
    MINUS_EQUALS
    STAR_EQUALS
    SLASH_EQUALS

    //syntax
    DOT
    DOT_DOT     //like 1..10
    SEMICOLON
    COLON
    QUESTION
    COMMA

    //keywords
    LET
    CONST 
    CLASS 
    NEW 
    IMPORT 
    EXPORT 
    FROM 
    FN 
    IF
    ELSE 
    FOR
    WHILE
    FOR_EACH
    TYPEOF
    IN 

    //Misc
    NUM_TOKENS
)

func TokenKindString(t TokenKind) string {
    //Gives the name of the tokenKind (poor man's enum)
    switch t {
    case EOF:
        return "eof" 
    case STRING:
        return "string"
    case NUMBER:
        return "number"
    case IDENTIFIER:
        return "identifier"
    case OPEN_BRACKET:
        return "open_bracket"
    case CLOSE_BRACKET:
        return "close_bracket"
    case OPEN_BRACE:
        return "open_brace"
    case CLOSE_BRACE:
        return "close_brace"
    case OPEN_PAREN:
        return "open_paren"
    case CLOSE_PAREN:
        return "close_paren"
    case EQUAL:
        return "equal"
    case ASSIGN:
        return "assign"
    case NOT_EQUAL:
        return "not_equal"
    case GREATER:
        return "greater"
    case LESSER:
        return "lesser"
    case GREAT_EQUAL:
        return "great_equal"
    case LESS_EQUAL:
        return "less_equal"
    case NOT:
        return "not"
    case AND:
        return "and"
    case OR:
        return "or"
    case PLUS:
        return "plus"
    case MINUS:
        return "minus"
    case STAR:
        return "star"
    case SLASH:
        return "slash"
    case CARROT:
        return "carrot"
    case PERCENT:
        return "percent"
    case PLUS_PLUS:
        return "plus_plus"
    case MINUS_MINUS:
        return "minus_minus"
    case PLUS_EQUALS:
        return "plus_equals"
    case MINUS_EQUALS:
        return "minus_equals"
    case STAR_EQUALS:
        return "star_equals"
    case SLASH_EQUALS:
        return "slash_equals"
    case DOT:
        return "dot"
    case DOT_DOT:
        return "dot_dot"
    case SEMICOLON:
        return "semicolon"
    case COLON:
        return "colon"
    case QUESTION:
        return "question"
    case COMMA:
        return "comma"
    case LET:
        return "let"
    case CONST:
        return "const"
    case CLASS:
        return "class"
    case NEW:
        return "new"
    case IMPORT:
        return "import"
    case EXPORT:
        return "export"
    case FROM:
        return "from"
    case FN:
        return "fn"
    case IF:
        return "if"
    case ELSE:
        return "else"
    case FOR:
        return "for"
    case WHILE:
        return "while"
    case FOR_EACH:
        return "for_each"
    case TYPEOF:
        return "typeof"
    case IN:
        return "in"
    default:
        return "unknown"
    }
}

var reserved_lu map[string]TokenKind = map[string]TokenKind{
    "true": TRUE,
    "false": FALSE,
    "null": NULL,
    "let": LET,
    "const": CONST,
    "class": CLASS,
    "new": NEW,
    "import": IMPORT,
    "from": FROM,
    "fn": FN,
    "if": IF,
    "else": ELSE,
    "foreach": FOR_EACH,
    "while": WHILE,
    "for": FOR, 
    "export": EXPORT,
    "typeof": TYPEOF,
    "in": IN,
}
