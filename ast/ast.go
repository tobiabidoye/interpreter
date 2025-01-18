package ast
import ("interpreter/token"
)

type Node interface{
    TokenLiteral() string
}

type Statement interface{
    //statement node 
    Node
    statementNode()
}

type Expression interface{
   //expression node 
    Node
    expressionNode()
}

type Program struct{
    Statements []Statement //array or slice of Statements
}

func (p * Program) TokenLiteral() string{
    
    if len(p.Statements) > 0{
        return p.Statements[0].TokenLiteral()
    }else{
        return ""
    }

}

type LetStatement struct{
    Token token.Token
    Name * Identifier
    Value Expression



}

func (ls * LetStatement) statementNode(){}

func(ls * LetStatement) TokenLiteral(){} 

type Identifier struct{
    Token token.Token
    Value string
}





