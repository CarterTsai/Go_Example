Password
--------

Generate L1, L2 and L3 password

### L1, L2, L3 Password

1. L1 : sha1(md5(pw+salty)))
2. L2 : sha512(sha256(pw+salty)))
3. L3 : sha512(sha256(pw+salty)))

### Testing

1. Get gocheck  

```
$> go get launchpad.net/gocheck
```

2. Update gocheck

```
$> go get -u launchpad.net/gocheck
```

3. Testing  

```
$>  make 
```

4. Benchmarks

```
$> go test -gocheck.b
```
