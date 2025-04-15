package logger

import (
	"fmt"
	_"log"
)

func Error(op string, event string) error{
	return fmt.Errorf("err: %s, happen: %s", op, event)
}