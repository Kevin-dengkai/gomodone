package gomodone

import (
	"errors"
	"fmt"
)

func SayHi(name, flag string) (string, error) {
	switch flag {
	case "php":
		return fmt.Sprintf("php Hi,%s", name), nil
	case "py":
		return fmt.Sprintf("php Hi,%s", name), nil
	default:
		return "", errors.New("unKnow lg")
	}
}
