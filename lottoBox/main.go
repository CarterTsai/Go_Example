package main

import (
    . "./lotto"
    "fmt"
    "math/rand"
    "time"
)

func msg(msg string) {
    fmt.Println("[LottoBox] : " + msg)
}

func main() {
    myList := ListBox { Count: 0, List : make([]byte, 102, 102)}

    box := new(LottoBox)
    rand.Seed(time.Now().Unix())

    msg("Lotto Go!\n")

    box.Drawnbox(49, &myList)

    box.Pop(box.Last_num)
    box.Drawnbox(48, &myList)

    box.Pop(box.Last_num)
    box.Drawnbox(5, &myList)

    print("00: ")
    for i:=1; i<=len(myList.List); i++ {
        if( i%6 != 0) {
            fmt.Printf("%0.2d ,", int(myList.List[i-1]))
        } else {
            fmt.Printf("%0.2d \n",myList.List[i-1])
            fmt.Printf("%0.2d: ", i/6)
        }
    }
}
