package generator

import (
	encHex "encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func Generate(quantity uint, hex, uppercase bool) []string {
	res := make([]string, 0, quantity)
	for i := uint(0); i < quantity; i++ {
		res = append(res, generate(hex, uppercase))
	}
	return res
}

func generate(hex, uppercase bool) string {
	var s string
	if id := uuid.New(); hex {
		s = encHex.EncodeToString(id[:])
	} else {
		s = id.String()
	}
	if uppercase {
		s = strings.ToUpper(s)
	}
	return s
}
