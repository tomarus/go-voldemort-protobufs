go-voldemort-protobufs
======================

Protocol buffer bindings for current Voldemort version 1.16

## Details

These bindings were generated with the Protocol buffers https://code.google.com/p/protobuf/
There's a good guide about them here - https://developers.google.com/protocol-buffers/docs/overview

To create these for yourself

1) Install protocol buffers - see link below
2) Download current Voldemort code  - see link below
3) Using the .proto files under src/ generate the language specific bindings using 

> protoc --go_out=. *.proto
  
4) Include these files in your Go program

> go get github.com/matzhouse/go-voldemort-protobufs

```go
import github.com/matzhouse/go-voldemort-protobufs
```

I'll link to some example code at some point in the future when I have it sorted.

## Questions?

Find me on twitter, and ask me a question :) 

**@mattalkingshit**


## Links

Project Voldemort 
http://www.project-voldemort.com/

Protocol Buffers
https://code.google.com/p/protobuf/

Go Protocol Buffers
https://code.google.com/p/goprotobuf/
