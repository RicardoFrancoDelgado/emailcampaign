package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contact := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contact)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contact := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contact)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contact := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contact)

	assert.Greater(campaign.CreatedOn, now)
}
