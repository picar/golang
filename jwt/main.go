package main

import (
	"fmt"
	jwtToken "jwt/jwtToken"
)

func main () {
	fmt.Println(jwtToken.CreateToken("pierre", "secret"))
}