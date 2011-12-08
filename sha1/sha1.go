package main

import "crypto/sha1"
import "flag"
import "fmt"

func main() {
	flag.Parse()
	sha1 := sha1.New()
	sha1.Write([]byte(flag.Arg(0)))
	fmt.Printf("%x\n", sha1.Sum(nil))
}
