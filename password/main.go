package main

import (
    pw "./password"
    "fmt"
)


func main() {
    fmt.Printf("Level1 : %x\n", pw.L1PwGenerator("test1"))
    fmt.Printf("Level1 : %s\n", pw.L1PwGenerator("test1"))
    fmt.Printf("Level2 : %x\n", pw.L2PwGenerator("test1"))
    fmt.Printf("Level3 : %x\n", pw.L3PwGenerator("test1"))
}
