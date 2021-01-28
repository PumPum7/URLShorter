package api

import (
	"math/rand"
	"urlShortener/database"
	"urlShortener/utils"
)

func GenName(length int, tries int) string {
	if tries > 6 {
		utils.Log.Error("Failed to generate name")
		return ""
	}

	for i := 0; i < 10; i++ {
		sequence := RandSeq(length)
		if !database.CheckKey(sequence) {
			return sequence
		}
	}
	return GenName(length+1, tries+1)

}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
