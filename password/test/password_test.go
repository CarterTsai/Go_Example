package testing

import (
    pw "../password"
    . "launchpad.net/gocheck"
    "testing"    
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {TestingT(t) }

type MySuite struct{}
var _ = Suite(&MySuite{})

func (s *MySuite) TestInit( c *C) {

    c.Assert(pw.GetSalty(), Equals, "what is your salte")
    pw.SetSalty("test1")
    c.Assert(pw.GetSalty(), Equals, "test1")
}

func (s *MySuite) TestAllLevelPassword( c *C) {

    l1 := []byte{0xdd, 0xf8, 0x2b, 0xcc, 0x16, 0x80, 0x26, 0x62, 0x87, 0xde, 0x4f, 0xd7, 0xe0, 0xf3, 0x68, 0x17, 0xc9, 0xe1, 0xc, 0x92}

    l2 := []byte{0x35, 0xb5, 0x3e, 0xea, 0x31, 0xf3, 0x94, 0xc8, 0xf9, 0x1b, 0xb6, 0x3c, 0x9, 0xb, 0xcd, 0xd6, 0xa6, 0x98, 0xd0, 0x15, 0x90, 0x39, 0x1d, 0x21, 0x3d, 0xad, 0xc0, 0x46, 0x7a, 0xca, 0xf2, 0x58, 0x9d, 0x2, 0x94, 0x18, 0xb4, 0xc5, 0x30, 0xbe, 0x49, 0x99, 0x18, 0x26, 0x76, 0x50, 0x5, 0xf2, 0x6e, 0x2c, 0xd2, 0x4b, 0x87, 0xd3, 0x1c, 0x79, 0xfd, 0xb6, 0x86, 0x54, 0x41, 0xdd, 0xb5, 0x13}

    l3 := []byte{0x8b, 0x38, 0x45, 0xab, 0x67, 0x6b, 0xf1, 0x2, 0x99, 0xa3, 0x77, 0xe5, 0x6c, 0x2a, 0x38, 0x50, 0xa9, 0x57, 0xca, 0xca, 0x4f, 0x29, 0xdf, 0x5c, 0xd2, 0x68, 0xa, 0x19, 0xd6, 0x26, 0xa5, 0xd6, 0xcd, 0xd0, 0x18, 0x94, 0x71, 0x9d, 0x44, 0x8b, 0x66, 0x35, 0x65, 0xe8, 0x3e, 0x74, 0xe9, 0xef, 0xed, 0x7c, 0xa, 0xca, 0x46, 0x49, 0xd4, 0x88, 0x17, 0x73, 0xd7, 0x63, 0xa5, 0x36, 0x92, 0x84} 

    c.Assert(string(pw.L1PwGenerator("test1")), Equals, string(l1))
    c.Assert(string(pw.L2PwGenerator("test1")), Equals, string(l2))
    c.Assert(string(pw.L3PwGenerator("test1")), Equals, string(l3))
}


func (s *MySuite) BenchmarkL1(c *C) {
    pw.L1PwGenerator("test1");
}

func (s *MySuite) BenchmarkL2(c *C) {
    pw.L2PwGenerator("test1");
}

func (s *MySuite) BenchmarkL3(c *C) {
    pw.L3PwGenerator("test1");
}