package utils

import (
	"math/rand"
	"time"
)

func GenerateAuthCode()(int, error){
	min := 100000
	max := 999999

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(max-min+1) + min

	return num, nil
}