package generator

import (
	encHex "encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func Generate(quantity uint, hex, uppercase bool) ([]string, error) {
	res := make([]string, quantity)
	for i := uint(0); i < quantity; i++ {
		s, err := generate(hex, uppercase)
		if err != nil {
			return nil, err
		}
		res[i] = s
	}
	return res, nil
}

func generate(hex, uppercase bool) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	var s string
	if hex {
		s = encHex.EncodeToString(id[:])
	} else {
		s = id.String()
	}

	if uppercase {
		s = strings.ToUpper(s)
	}

	return s, nil
}
