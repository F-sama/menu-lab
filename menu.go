package main  

import "fmt" 

func main(){  
    var cmd string
    for {
        fmt.Println("please input command")
        fmt.Scanln(&cmd)
        if cmd == "hello" {
            fmt.Println("hello commander.")
        } else if cmd == "quit" {
            fmt.Println("Bye")
            break
        } else {
            fmt.Println("please input right command!")
        }
    }
}
