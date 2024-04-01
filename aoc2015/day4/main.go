package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	str1 := "yzbqklnj"
	for i := 0; i < 10000000; i++ {
		concat := fmt.Sprintf("%s%d", str1, i)
		if GetMD5Hash(concat)[:6] == "000000" {
			fmt.Println(i)
			break
		}
	}
}
