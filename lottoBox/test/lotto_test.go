package example
import (
    . "launchpad.net/gocheck"
    . "../lotto"
    "testing" 
    )

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {TestingT(t) }

type MySuite struct{}
var _ = Suite(&MySuite{})

func (s *MySuite) TestInit( c *C) {

    myList := ListBox { Count: 0, List : make([]byte, 102, 102)}
    c.Assert(myList.Count, Equals, byte(0))
    c.Assert(len(myList.List), Equals, 102)
    c.Assert(cap(myList.List), Equals, 102)
}

func (s *MySuite) TestFillBox( c *C) {
    box := new(LottoBox)
    box.Fillbox()
    c.Assert(box.Num[0] , Equals, byte(1))
    c.Assert(box.Num[10], Equals, byte(11))
    c.Assert(box.Num[48], Equals, byte(49))
}

func (s *MySuite) TestPushandPop( c *C) {
    box := new(LottoBox)
    // PUSH
    box.Push(10)
    c.Assert(box.Num[0], Equals, byte(10))
    box.Push(33)
    c.Assert(box.Num[1], Equals, byte(33))
    box.Push(50)
    c.Assert(box.Num[2], Equals, byte(50))
    // POP
    box.Pop(1)
    c.Assert(box.Num[0], Equals, byte(0))
    box.Pop(2)
    c.Assert(box.Num[1], Equals, byte(0))
    box.Pop(3)
    c.Assert(box.Num[2], Equals, byte(0))
}
