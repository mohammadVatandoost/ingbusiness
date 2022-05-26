package iam

import (
	"context"
	iamV1 "github.com/mohammadVatandoost/ingbusiness/api/services/iam/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/ingpages"
)

func (s *Service) AddIngPage(ctx context.Context,
	ingPage *iamV1.IngPage, creatorID int32) (ingpages.IngPage, error) {
	return s.ingDirectory.AddIngPage(ctx, ingpages.AddIngPageParams{
		Name:           ingPage.Name,
		OrganizationID: ingPage.OrganizationId,
		Token:          ingPage.Token,
		CreatorID:      creatorID,
	})
}

func (s *Service) GetIngPageByOrganizationID(ctx context.Context,
	organizationID int32) ([]ingpages.IngPage, error) {
	return s.ingDirectory.GetIngPageByOrganizationID(ctx, organizationID)
}

func (s *Service) DeleteIngPage(ctx context.Context,
	id int32) (ingpages.IngPage, error) {
	return s.ingDirectory.DeleteIngPage(ctx, id)
}
