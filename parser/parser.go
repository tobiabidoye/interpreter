package parser

import ("interpreter/ast"
        "interpreter/token"
        "interpreter/lexer"
)

type Parser struct{
    l * lexer.Lexer
    curToken token.Token
    peekToken token.Token
}

func New(l * lexer.Lexer) *Parser{
    p:= &Parser{l: l}
    p.nextToken()
    p.nextToken()

    return p
}


func (p * Parser) nextToken(){
   p.curToken = p.peekToken
   p.peekToken = p.l.NextToken()

}


func (p * Parser) parseProgram() *ast.Program{
      
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Type != token.EOF{
        stmt := p.parseStatement()
        if stmt != nil{
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program 
}

func (p * Parser) parseStatement() ast.Statement{
    switch p.curToken.Type{
    case token.LET:
        return p.parseLetStatement()
    default: 
        return nil
    
    }

}



