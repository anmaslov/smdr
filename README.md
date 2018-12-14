# smdr
Golang package for extracting and parse SMDR data from NEC SV8100 PBX systems.

## Usage

### Install package
`go get -v github.com/anmaslov/smdr/`

### Creating a client
```go
conn, err := net.Dial("tcp", addr)
if err != nil {
    log.Fatal("dial error:", err)
    return
}
defer conn.Close()
```

### Send request to PBX

```go
r1 := setRequest(dataRequest())
if wr, err := conn.Write([]byte(r1)); //Request #1
    wr == 0 || err != nil {
        log.Fatal(err)
}

conn.SetReadDeadline(time.Now().Add(5 * time.Second))
//trying to get response from PBX
buff := make([]byte, 1024)
rd, err := conn.Read(buff)
if err != nil{
	log.Fatal(err)
}
//response received

//trying parse data
res := CDR{}
err = res.parser(buff[:rd])
if err != nil {
    log.Fatal(err)
}
//store to file or db, example to MongoDB

//all ok, sending client response sequence
r4 := setRequest(clientResponse(res.sequence))
if wr, err := conn.Write([]byte(r4)); //request #4
    wr == 0 || err != nil {
    log.Fatal(err)
}
``` 
More examples is coming soon.

## Tests
Only tested on the NEC SV8300.

## License
**SMDR package** is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
