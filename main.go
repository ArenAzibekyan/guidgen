package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type generateCommand struct {
	quantity  int
	hex       bool
	uppercase bool
}

func (cmd generateCommand) Validate() error {
	if cmd.quantity >= 1 {
		return nil
	}
	return fmt.Errorf("invalid quantity %d", cmd.quantity)
}

func (cmd generateCommand) uuidString() string {
	var s string
	if id := uuid.New(); cmd.hex {
		s = hex.EncodeToString(id[:])
	} else {
		s = id.String()
	}
	if !cmd.uppercase {
		return s
	}
	return strings.ToUpper(s)
}

func (cmd generateCommand) Print() {
	for i := 0; i < cmd.quantity; i++ {
		fmt.Println(cmd.uuidString())
	}
}

func main() {
	var cmd generateCommand

	flag.IntVar(&cmd.quantity, "q", 1, "quantity")
	flag.BoolVar(&cmd.hex, "h", false, "hex")
	flag.BoolVar(&cmd.uppercase, "u", false, "uppercase")
	flag.Parse()

	if err := cmd.Validate(); err != nil {
		log.Fatal(err)
	}
	cmd.Print()
}
