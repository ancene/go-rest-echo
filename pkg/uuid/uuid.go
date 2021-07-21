package uuid

import guuid "github.com/google/uuid"

func UUID() string {
	return guuid.New().String()
}
