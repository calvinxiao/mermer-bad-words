package main

import (
	"fmt"
	"strings"

	"github.com/AfterShip/gopkg/uuid"
	"github.com/martinlindhe/base36"
	"github.com/spaolacci/murmur3"
)

func main() {
	showLen()
	showLen()
	showLen()
	as := "https://www.aftership.com"
	asId := GenerateShortID(as)
	fmt.Printf("Hash(\"%s\") = %s, len is %d\n", as, asId, len(asId))

	// calculate new uuid 1M times to see how many rude words
	tries := 100000
	rudeCount := 0
	for i := 0; i < tries; i++ {
		input := uuid.GenerateUUIDV4()
		id := GenerateShortID(input)
		for _, rudeWord := range rudeWords {
			// if rudeWord != "fuck" {
			// 	continue
			// }
			rudeWord = strings.ToUpper(rudeWord)
			if strings.Contains(id, rudeWord) {
				rudeCount++
				// fmt.Printf("uuid: %s, hash: %s, contains rude word: %s\n", input, id, rudeWord)
			}
		}
	}
	fmt.Printf("Found rude words: %d, total: %d\n", rudeCount, tries)
}

func showLen() {
	s := uuid.GenerateUUIDV4()
	id := GenerateShortID(s)
	fmt.Printf("Hash(\"%s\") = %s, len is %d\n", s, id, len(id))
}

func GenerateShortID(raw string) string {
	return base36.Encode(murmur3.Sum64([]byte(raw)))
}
