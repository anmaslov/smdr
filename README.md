# smdr
Golang package for extracting and parse SMDR data from NEC SV8100 PBX systems.

## Usage

### Install package
`go get -v github.com/anmaslov/smdr/`

or use as go modules:
```go
import (
    "github.com/anmaslov/smdr"
)
```
and run `go mod init` then run `go build`

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
r1 := smdr.SetRequest(smdr.DataRequest())
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
res := smdr.CDR{}
err = res.Parser(buff[:rd])
if err != nil {
    log.Fatal(err)
}
//store to file or db, example to MongoDB

//all ok, sending client response sequence
r4 := smdr.SetRequest(smdr.ClientResponse(res.sequence))
if wr, err := conn.Write([]byte(r4)); //request #4
    wr == 0 || err != nil {
    log.Fatal(err)
}
``` 

## Tests
Only tested on the NEC SV8300.

## License
**SMDR package** is released under the [MIT License](http://opensource.org/licenses/mit-license.php).
