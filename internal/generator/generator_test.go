package generator

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name      string
		quantity  uint
		hex       bool
		uppercase bool
		reg       string
	}{
		{
			name:     "000",
			quantity: 0,
		},
		{
			name:      "100",
			quantity:  1,
			hex:       false,
			uppercase: false,
			reg:       `\A[0-9a-f]{8}[-][0-9a-f]{4}[-][0-9a-f]{4}[-][0-9a-f]{4}[-][0-9a-f]{12}\z`,
		},
		{
			name:      "201",
			quantity:  2,
			hex:       false,
			uppercase: true,
			reg:       `\A[0-9A-F]{8}[-][0-9A-F]{4}[-][0-9A-F]{4}[-][0-9A-F]{4}[-][0-9A-F]{12}\z`,
		},
		{
			name:      "310",
			quantity:  3,
			hex:       true,
			uppercase: false,
			reg:       `\A[0-9a-f]{32}\z`,
		},
		{
			name:      "411",
			quantity:  4,
			hex:       true,
			uppercase: true,
			reg:       `\A[0-9A-F]{32}\z`,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			reg := regexp.MustCompile(c.reg)
			res, err := Generate(c.quantity, c.hex, c.uppercase)
			require.NoError(t, err)
			require.Len(t, res, int(c.quantity))
			for _, s := range res {
				require.True(t, reg.MatchString(s))
			}
		})
	}
}
