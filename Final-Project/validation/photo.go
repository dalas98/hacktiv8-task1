package validation

import (
	"dalas98/golang-lesson/Final-Project/model/modelphoto"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidatePhotoCreate(data modelphoto.Request) error {
	return validation.Errors{
		"title":     validation.Validate(data.Title, validation.Required),
		"caption":   validation.Validate(data.Caption),
		"photo_url": validation.Validate(data.PhotoURL, validation.Required),
		"userId":    validation.Validate(data.UserID),
	}.Filter()
}
