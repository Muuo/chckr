package main

import (
    "os"
    "log"
    "fmt"
    "hash"
    "strings"
    "io"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
)

func main() {
    var hs string
    if (len(os.Args) < 3) {
        log.Fatal("Please specify the desired hash algorithm and the target file")
    } else if (len(os.Args) == 4) {
	hs = strings.TrimSpace(os.Args[3])
    }else if (len(os.Args) > 4) {
        log.Fatal("Provide a maximum of 3 arguments")
    }

    method := strings.ToLower(strings.TrimSpace(os.Args[1]))
    fname := strings.TrimSpace(os.Args[2])

    f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    var h hash.Hash
    if (method == "sha512") {
        h = sha512.New()
    } else if (method == "sha256") {
        h = sha256.New()
    } else if (method == "sha1") {
        h = sha1.New()
    } else if (method == "md5") {
        h = md5.New()
    } else {
        f.Close()
        log.Fatal("Hash not implemented!")
    }

    if _, err := io.Copy(h, f); err != nil {
        log.Fatal(err)
    }
    hsc := fmt.Sprintf("%x", h.Sum(nil))
    fmt.Println(hsc)
    if (len(os.Args) == 4) {
	    if (hs == hsc) {
		    fmt.Println("Hash matches")
	    } else {
		    fmt.Println("Hash does not match!")
	    }
    }
}
