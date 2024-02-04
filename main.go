package main

import (
	"fmt"
	"github.com/CoboGlobal/cobo-go-api/cobo_custody"
)

func main() {
	apiSecret, apiKey := cobo_custody.GenerateKeyPair()
	fmt.Println("API_SECRET:", apiSecret)
	fmt.Println("API_KEY:", apiKey)
}
