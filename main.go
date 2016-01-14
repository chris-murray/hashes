// takes a string and prints the sha256 hash
package main

import "fmt"
import "flag"
import "crypto/sha256"
import "crypto/md5"
import "encoding/hex"

//var m = flag.Bool("n", false, "omit trailing newline")
var userString = flag.String("s", "p4ssw0rd1", "string to hash")
var mode = flag.String("m", "all", "which type of hash to generate, all, md5, sha256")

//GetMD5Hash pass a string, get an md5 hash back
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//GetSHA256Sum pass a string, get a sha256 hash back.
func GetSHA256Sum(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	flag.Parse()

	fmt.Printf("Your String: %s\n", *userString)

	if *mode == "all" {
		fmt.Printf("sha256 sum: %s\n", GetSHA256Sum(*userString))
		fmt.Printf("md5 Hash: %s", GetMD5Hash(*userString))
	} else if *mode == "md5" {
		fmt.Printf("md5 Hash: %s", GetMD5Hash(*userString))
	} else if *mode == "sha256" {
		fmt.Printf("sha256 sum: %s\n", GetSHA256Sum(*userString))
	} else {
		fmt.Printf("Sorry, I don't reconize that mode: %s\n", *mode)
	}
}
