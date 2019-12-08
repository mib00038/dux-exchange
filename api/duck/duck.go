package duck

import "math/rand"

type color string

const (
	Yellow color = "yellow"
	Red    color = "red"
	Blue   color = "blue"
	Green  color = "green"
)

type Duck struct {
	Color color `json:"color"`
}

func Rand() Duck {
	switch rand.Intn(3) + 1 {
	case 4:
		return Duck{Red}
	case 3:
		return Duck{Red}
	case 2:
		return Duck{Red}
	case 1:
		fallthrough
	default:
		return Duck{Yellow}
	}
}
