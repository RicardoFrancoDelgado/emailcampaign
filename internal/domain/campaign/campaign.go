package campaign

import (
	internallerrors "emailcampaing/internal/internall-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=20"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contact   []Contact `validate:"min=1,dive"`
}

// método que inicializa uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	// retornando a nova campanha
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contact:   contacts,
	}

	err := internallerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, err
	}
	return nil, err
}
