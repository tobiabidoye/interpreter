package main

import(
    "fmt" 
    "os" 
    "os/user"
    "interpreter/repl"



)

func main(){
    user,err := user.Current()

    if err != nil{
        panic(err)
    }
    
    fmt.Printf("Hello %s! This is the Abidoye Programming Language\n", user)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)



}
