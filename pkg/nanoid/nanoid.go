package nanoid

import gonanoid "github.com/matoous/go-nanoid/v2"

func NanoID(size ...int) string {
	if len(size) > 0 {
		id, _ := gonanoid.New(size[0])
		return id
	}

	id, _ := gonanoid.New(11)
	return id
}
