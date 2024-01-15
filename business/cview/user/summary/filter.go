package summary

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/theo-bot/service4.1-video/business/sys/validate"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	UserID   *uuid.UUID `validate:"omitempty,uuid4"`
	UserName *string    `validate:"omitempty,min=3"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithUserID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithUserID(userID uuid.UUID) {
	qf.UserID = &userID
}

// WithUserName sets the UserName field of the QueryFilter value.
func (qf *QueryFilter) WithUserName(userName string) {
	qf.UserName = &userName
}
