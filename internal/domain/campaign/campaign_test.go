package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body teste"
	contact = []string{"email1@e.com", "email2@e.com"}

	fake = faker.New()
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

func Test_NewCampaign_NameMustBeValidateWithMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contact)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_NameMustBeValidateWithMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(35), content, contact)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_ContentMustBeRequiredWithMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contact)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_ContentMustBeRequiredWithMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1045), contact)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_ContactMustBeRequiredWithMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contact is required with min 1", err.Error())
}

func Test_NewCampaign_ContactMustBeRequired(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
