package campaign

import (
	"emailcampaing/internal/contract"
	internallerrors "emailcampaing/internal/internall-errors"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaignDto) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internallerrors.ErrInternal
	}

	return campaign.ID, nil
}
