package worktime

import (
	"fmt"
	"time"
)

const (
	layout1 = "03:04:05PM"
	layout2 = "15:04:05"
)

func WorkTime(value string) (string, string, error) {
	t, err := time.Parse(layout1, value)
	if err != nil {
		return "", "", err

	}
	fmt.Println(t.Format(layout1))
	fmt.Println(t.Format(layout2))
	return t.Format(layout1), t.Format(layout2), nil
}
