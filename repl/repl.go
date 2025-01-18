package repl

import("bufio"
        "fmt"
        "io"
        "interpreter/lexer"
        "interpreter/token"

)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer){
    scanner := bufio.NewScanner(in)
    for{
        fmt.Fprintf(out, PROMPT)
        scanned := scanner.Scan()
        if !scanned{
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        //we already scanned first token so we tokenize the remaining
        //from next token till end of file
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken(){
            fmt.Fprintf(out, "%+v\n", tok)
        }


    }   


}
