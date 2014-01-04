package password

import (
    "crypto/sha512"
    "crypto/sha256"
    "crypto/sha1"
    "crypto/md5"
    "crypto/hmac"
    "io"
)


var salty string
var key []byte

func init() {
    salty = "what is your salte"
    key = []byte("146c07dahsdj2j3k12j312k11kkkkjfl;lb;hshdgdd3d")
}

/**
 *  @rerutn sha1(md5(pw+salty)))
 */

func L1PwGenerator(pw string) []byte {
    _hash1 := md5.New()
    _hash2 := sha1.New()

    io.WriteString(_hash1, pw+salty)
    io.WriteString(_hash2, string(_hash1.Sum(nil)))
    return _hash2.Sum(nil)
}

/**
 *  @rerutn sha512(sha256(pw+salty)))
 */

func L2PwGenerator(pw string) []byte {
    _hash1 := sha256.New()
    _hash2 := sha512.New()

    io.WriteString(_hash1, pw+salty)
    io.WriteString(_hash2, string(_hash1.Sum(nil)))
    return _hash2.Sum(nil)
}

/**
 *  @rerutn sha512(sha256(pw+salty)))
 */

func L3PwGenerator(pw string) []byte {
    _hmac := hmac.New(sha512.New, key)
    _hmac.Write([]byte(pw+salty))
    return _hmac.Sum(nil)
}

func SetSalty(text string) {
    salty = text
}

func GetSalty() string {
    return salty
}

func SetKey(text string) {
    key = []byte(text)
}

func GetKey() []byte {
    return key
}
