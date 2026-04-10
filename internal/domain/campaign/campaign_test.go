package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contact := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contact)

	if campaign.ID != "1" {
		t.Error("Expected 1 ", campaign.ID)
	} else if campaign.Name != name {
		t.Error("Expected correct name")
	} else if campaign.Content != content {
		t.Error("Expected correct content")
	} else if len(campaign.Contact) != len(contact) {
		t.Error("Expected correct contact length")
	}
}
