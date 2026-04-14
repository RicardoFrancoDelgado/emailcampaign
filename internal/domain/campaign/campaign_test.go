package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body"
	contact = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contact)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contact), len(contact))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contact)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contact)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_NameMustBeValidate(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contact)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_ContentMustBeRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contact)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_ContactMustBeRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contact is required", err.Error())
}
