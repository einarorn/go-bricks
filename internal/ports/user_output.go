package ports

import "go-bricks/internal/models"

type UserOutput interface {
	Draw(status models.GameStatus)
}
