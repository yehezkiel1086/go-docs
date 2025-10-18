package cloudinary

import (
	"go-gin-nextjs-file-cloudinary/config"

	"github.com/cloudinary/cloudinary-go/v2"
)

type Cloudinary struct {
	*cloudinary.Cloudinary
}

func InitCloudinary(conf *config.Cloudinary) (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(conf.Name, conf.APIKey, conf.APISecret)
	if err != nil {
		return nil, err
	}

	return &Cloudinary{
		cld,
	}, nil
}
