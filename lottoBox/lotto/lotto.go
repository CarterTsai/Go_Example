package lotto

import (
    "math/rand"
)

type ListBox struct{
    List []byte
    Count byte
}

type LottoBox struct{
    Num [49]byte
    Count byte
    Last_num byte
}

func init() {
}

func (box *LottoBox) Push(ball_num byte) bool {
    box.Num[box.Count] = ball_num
    box.Count++
    return true;
}

func (box *LottoBox) Pop(ball_num byte) byte {
    var tmp byte
    tmp, box.Num[ball_num -1 ] = box.Num[ball_num - 1], 0
    box.Count--
    box.Last_num = tmp
    return tmp;
}

func (box *LottoBox) Fillbox() bool {
    // put ball in to LottoBox
    box.Count = 0
    for i:=1 ; i<=49; i++ {
        box.Push(byte(i))
    }
    return true;
}

func (box *LottoBox) Drawnbox(num int, myList *ListBox) {
    box.Fillbox()
    for i:=0; i<num; i++ {
        tmp := box.Pop(byte(rand.Intn(49)+1))
        for tmp == 0 {
            tmp = box.Pop(byte(rand.Intn(49)+1))
        }
        myList.List[myList.Count] = tmp
        myList.Count++
    }
}
