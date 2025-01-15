package token

type TokenType string

type Token struct{
    Type TokenType
    Literal string
}

const(
    ILLEGAL = "ILLEGAL"
    EOF = "EOF" 


    //identifiers and literals
    IDENT = "IDENT"
    INT = "INT"
    
    //Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERIK = "*" 
    SLASH = "/"

    LT = "<" 
    GT = ">"

    //Delimiters
    COMMA = ","
    SEMICOLONS = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    //Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE" 
    FALSE = "FALSE" 
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)

var keywords = map[string]TokenType{
    "fn":FUNCTION,
    "let":LET,
    "true": TRUE, 
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN, 
}

func LookUpIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{ //if identifier exists in the map execute the if statement
       return tok; 
    } 


   return IDENT 

}
