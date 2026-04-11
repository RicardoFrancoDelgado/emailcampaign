package campaign

import "emailcampaing/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) create(newCampaign contract.NewCampaignDto) error {
	return nil
}
