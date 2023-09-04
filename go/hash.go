package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := hash("password")
	fmt.Printf(hash)
}

func hash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
