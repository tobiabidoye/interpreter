package lexer

type Lexer struct{
    input string
    position int //represents current position in input
    readPosition int // current reading position in input
    ch byte // current byte that is being looked at
}

func New (input string) *Lexer{
    l := &Lexer{input: input}
    return l

}



