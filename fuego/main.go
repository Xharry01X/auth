package main

import (
	"github.com/go-fuego/fuego"
)


func main(){
s :=fuego.NewServer()  //s is a server 

fuego.Get(s,"/",helloworld) //hellowrold is a controller
s.Run()
}

func helloworld(c fuego.ContextNoBody)(string, error){
	return "Hello world", nil
}