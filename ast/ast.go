package ast

type Node interface{
    TokenLiteral() string
}

type statement interface{
    //statement node 
    Node
    statementNode()
}

type Expression interface{
   //expression node 
    Node
    expressionNode()
}





