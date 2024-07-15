package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func loginServer(userInput string)error {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        return (err);
    }
    defer conn.Close();

    _, err = conn.Write([]byte(userInput))
    if err != err {
        return err;
    }

    return nil;
}

func getUserInput() string{
    scanner := bufio.NewScanner(os.Stdin);
    scanner.Scan();
    return scanner.Text();
}

func main() {
    fmt.Println("type 'Quit!' if you want to leave");
    closeserver := "Quit!";
    for {
        fmt.Println("Input: ");
        userInput := getUserInput();

        if strings.TrimSpace(userInput) == closeserver {
            break;
        }

        err := loginServer(userInput);
        if err != nil{
            fmt.Println("Error,", err);
        }
    }
}
