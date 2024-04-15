package language

import "fmt"

type TokenType string

const (
    //Single character tokens
    LEFT_PAREN TokenType = "LEFT_PAREN"
    RIGHT_PAREN TokenType = "RIGHT_PAREN"
    LEFT_BRACE TokenType = "LEFT_BRACE"
    RIGHT_BRACE TokenType = "RIGHT_BRACE"

    COMMA TokenType = "COMMA"
    DOT TokenType = "DOT"
    MINUS TokenType = "MINUS"
    PLUS TokenType = "PLUS"
    SEMICOLON TokenType = "SEMICOLON"
    SLASH TokenType = "SLASH"
    STAR TokenType = "STAR"

    //One or two character tokens
    BANG TokenType = "BANG"
    BANG_EQUAL TokenType = "BANG_EQUAL"
    EQUAL TokenType = "EQUAL"
    EQUAL_EQUAL TokenType = "EQUAL_EQUAL"
    GREATER TokenType = "GREATER"
    GREATER_EQUAL TokenType = "GREATER_EQUAL"
    LESS TokenType = "LESS"
    LESS_EQUAL TokenType = "LESS_EQUAL"

    //Literals
    IDENTIFIER TokenType = "IDENTIFIER"
    STRING TokenType = "STRING"
    NUMBER TokenType = "NUMBER"

    //Keywords
    AND TokenType = "AND"
    CLASS TokenType = "CLASS"
    ELSE TokenType = "ELSE"
    FALSE TokenType = "FALSE"
    FUN TokenType = "FUN"
    FOR TokenType = "FOR"
    IF TokenType = "IF"
    NIL TokenType = "NIL"
    OR TokenType = "OR"
    PRINT TokenType = "PRINT"
    RETURN TokenType = "RETURN"
    SUPER TokenType = "SUPER"
    THIS TokenType = "THIS"
    TRUE TokenType = "TRUE"
    VAR TokenType = "VAR"
    WHILE TokenType = "WHILE"
    BREAK TokenType = "BREAK"

    EOF TokenType = "EOF"

)

var Keywords = map[string]TokenType{
    "and": "AND",
    "class": "CLASS",
    "else": "ELSE",
    "false": "FALSE",
    "for": "FOR",
    "fun": "FUN",
    "if": "IF",
    "nil": "NIL",
    "or": "OR",
    "print": "PRINT",
    "return": "RETURN",
    "super": "SUPER",
    "this": "THIS",
    "true": "TRUE",
    "var": "VAR",
    "while": "WHILE",
    "break": "BREAK",
}

type Token struct{
    TokenType TokenType;
    Lexeme string;
    Literal interface{};
    Line int;
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) *Token {
    token := new(Token)
    token.TokenType = tokenType
    token.Lexeme = lexeme
    token.Literal = literal
    token.Line = line
    return token
}

func (token *Token) String() string {

    return fmt.Sprintf("%s %s %s",token.TokenType,token.Lexeme, token.Literal)
}
