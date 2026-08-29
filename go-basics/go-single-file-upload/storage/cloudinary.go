package storage

import (
	"context"
	"go-single-file-upload/config"

	"github.com/cloudinary/cloudinary-go/v2"
)

type Cloudinary struct {
	cld *cloudinary.Cloudinary
}

func InitCloudinary(ctx context.Context, conf *config.Cloudinary) (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(conf.Name, conf.Key, conf.Secret)
	if err != nil {
		return &Cloudinary{}, err
	}

	return &Cloudinary{
		cld: cld,
	}, nil
}

func (c *Cloudinary) GetCld() *cloudinary.Cloudinary {
	return c.cld
}
