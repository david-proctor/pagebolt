package main

import (
    "fmt"
    "os"
    "github.com/pagebolt/templates"
)

func main() {
    PrintLogo()

    args := os.Args[1:]
    if len(args) != 1 {
        PrintHelp()
        return
    }

    scanner := templates.DirectoryScannerImpl{
        RootPath : args[0],
    }
    fmt.Println("Building templates from path", scanner.RootPath)
    templates := templates.AssembleTemplates(scanner)

    for _,template := range templates {
        fmt.Println("=============================")
        fmt.Println(template.String())
    }
}

func PrintLogo() {
    fmt.Println("___   ___  ____  _____   ___  __  __ ________")
    fmt.Println("|   \\/   |/ ___||   __| /   \\/  \\/ //__   __/")
    fmt.Println("|  _/ _  | \\/ \\ |  _|  /   </ / / /__/  /")
    fmt.Println("|_|/_/ |_|\\__/_||____|/____/\\__/____/__/")
}

func PrintHelp() {
    fmt.Println("PageBolt is a tool for building simple, fast-loading web sites.")
    fmt.Println()
    fmt.Println("Syntax: pagebolt path")
    fmt.Println("    path    the absolute path to the folder containing your pagebolt template files")
}