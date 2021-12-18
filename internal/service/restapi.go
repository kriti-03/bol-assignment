package service

import (
	"github.com/pablocrivella/mancala/internal/utility"
)

type (
	RequestBody interface {
		Validate() utility.ValidationErrors
	}

	Resources struct {
		Games GamesResource
	}
)
