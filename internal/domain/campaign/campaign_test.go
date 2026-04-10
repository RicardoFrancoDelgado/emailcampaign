package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contact := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contact)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}
