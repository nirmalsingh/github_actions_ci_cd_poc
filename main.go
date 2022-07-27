package main

import "fmt"

func main() {
    msg := sayHello("Alice")
    fmt.Println(msg)
}

func sayHello(name string) string {
    return fmt.Sprintf("Hi %s", name)
// Change this to "Hello %s" instead of "Hi %s". 
    return fmt.Sprintf("Hello %s", name)
}
