package main

import (
    "fmt"
    "math/rand"
    "time"
)

type ListBox struct{
    list []byte
    count byte
}

type LottoBox struct{
    num [49]byte
    count byte
    last_num byte
}

func (box *LottoBox) push(ball_num byte) bool {
    box.num[box.count] = ball_num
    box.count++
    return true;
}

func (box *LottoBox) pop(ball_num byte) byte {
    var tmp byte
    tmp, box.num[ball_num -1 ] = box.num[ball_num - 1], 0
    box.count--
    box.last_num = tmp
    return tmp;
}

func (box *LottoBox) fillbox() bool {
    // put ball in to LottoBox
    box.count = 0
    for i:=1 ; i<=49; i++ {
        box.push(byte(i))
    }
    return true;
}

func (box *LottoBox) drawnbox(num int, myList *ListBox) {
    box.fillbox()
    for i:=0; i<num; i++ {
        tmp := box.pop(byte(rand.Intn(49)+1))
        for tmp == 0 {
            tmp = box.pop(byte(rand.Intn(49)+1))
        }
        myList.list[myList.count] = tmp
        myList.count++
    }
}

func msg(msg string) {
    fmt.Println("[LottoBox] : " + msg)
}

func main() {
    myList := ListBox { count: 0, list : make([]byte, 102, 102)}

    box := new(LottoBox)
    rand.Seed(time.Now().Unix())

    msg("Lotto Go!\n")

    box.drawnbox(49, &myList)

    box.pop(box.last_num)
    box.drawnbox(48, &myList)

    box.pop(box.last_num)
    box.drawnbox(5, &myList)

    print("00: ")
    for i:=1; i<=len(myList.list); i++ {
        if( i%6 != 0) {
            fmt.Printf("%0.2d ,", int(myList.list[i-1]))
        } else {
            fmt.Printf("%0.2d \n",myList.list[i-1])
            fmt.Printf("%0.2d: ", i/6)
        }
    }
}
