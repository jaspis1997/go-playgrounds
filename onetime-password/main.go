package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	const (
		secret = "test"
	)
	key, err := totp.GenerateCode(secret, time.Now())
	fmt.Println(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(os.Stdin)
	passcode, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	// valid := totp.Validate(key, secret)
	valid := totp.Validate(string(passcode), secret)
	if valid {
		fmt.Println("OK")
	}
}
