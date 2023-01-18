package image

import "time"

type Core struct {
	ID         uint
	Image_url  string `validate:"required"`
	Post_ID    uint
	Created_at time.Time
}

type ImageService interface {
	CreatedImage(input Core) error
	Getall() ([]Core, error)
}

type ImageData interface {
	CreatedImage(input Core) error
	Getall() ([]Core, error)
}
