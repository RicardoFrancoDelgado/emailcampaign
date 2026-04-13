package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contact   []Contact
}


// método que inicializa uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	// validando os campos se estão todos preenchidos
	if name == "" {
		return nil, errors.New("Name must be required")
	} else if content == "" {
		return nil, errors.New("Content must be required")
	} else if len(emails) == 0 {
		return nil, errors.New("Contact must be required")
	}

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	// retornando a nova campanha
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contact:   contacts,
	}, nil
}
