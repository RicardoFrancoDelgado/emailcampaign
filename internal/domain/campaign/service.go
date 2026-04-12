package campaign

import (
	"emailcampaing/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDto) (string, error) {

	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	s.Repository.Save(campaign)

	return campaign.ID, nil
}
