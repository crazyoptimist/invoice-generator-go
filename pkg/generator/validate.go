package generator

import (
	"github.com/go-playground/validator/v10"
)

// Validate document fields
func (d *Document) Validate() error {
	validate := validator.New()
	if err := validate.Struct(d); err != nil {
		return err
	}

	// Prepare items
	for _, item := range d.Items {
		if err := item.Prepare(); err != nil {
			return err
		}
	}

	return nil
}
