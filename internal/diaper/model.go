package diaper

import (
	"errors"
	"time"
)

type Type string

const (
	TypeWet   Type = "wet"
	TypeDry   Type = "dry"
	TypeMixed Type = "mixed"
)

type Diaper struct {
	BabyID     string
	OccurredAt time.Time
	Type       Type
	Notes      string
}

func (d Diaper) Validate() error {
	errs := make([]error, 0)

	if d.BabyID == "" {
		errs = append(errs, errors.New("diaper requires a valid BabyID"))
	}
	if len(d.Notes) > 500 {
		errs = append(errs, errors.New("diaper notes cannot exceed 500 characters."))
	}
	if d.Type != TypeDry && d.Type != TypeMixed && d.Type != TypeWet {
		errs = append(errs, errors.New("diaper type must be wet, dry or mixed."))
	}
	if d.OccurredAt.IsZero() {
		errs = append(errs, errors.New("diaper must occur at a valid time."))
	}
	if d.OccurredAt.After(time.Now()) {
		errs = append(errs, errors.New("diaper cannot occur in the future."))
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}
