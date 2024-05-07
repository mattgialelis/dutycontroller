package pd

import (
	"context"
	"fmt"
	"log"

	"github.com/PagerDuty/go-pagerduty"
)

type Service struct {
	Name               string
	Description        string
	EscalationPolicyID string
	AutoResolveTimeout int
	AcknowedgeTimeout  int
}

func (s *Service) ToPagerDutyService() pagerduty.Service {
	return pagerduty.Service{
		Name:        s.Name,
		Description: s.Description,
		EscalationPolicy: pagerduty.EscalationPolicy{
			APIObject: pagerduty.APIObject{
				ID:   s.EscalationPolicyID,
				Type: "escalation_policy_reference",
			},
		},
	}
}

func (p *Pagerduty) CreatePagerDutyService(service Service) (string, error) {

	serviceInput := service.ToPagerDutyService()

	newService, err := p.client.CreateServiceWithContext(context.TODO(), serviceInput)

	if err != nil {
		return "", fmt.Errorf("failed to create service: %w", err)
	}

	return newService.ID, nil
}

func (p *Pagerduty) UpdatePagerDutyService(service Service) error {

	serviceInput := service.ToPagerDutyService()

	existingService, _, err := p.GetPagerDutyServiceByNameDirect(service.Name)
	if err != nil {
		return err
	}

	serviceInput.ID = existingService.ID

	_, err = p.client.UpdateServiceWithContext(context.TODO(), serviceInput)

	if err != nil {
		return fmt.Errorf("failed to update service: %w", err)
	}

	return nil
}

func (p *Pagerduty) DeletePagerDutyService(name string) error {

	service, _, err := p.GetPagerDutyServiceByNameDirect(name)
	if err != nil {
		return fmt.Errorf("failed to get service by name: %w", err)
	}

	err = p.client.DeleteServiceWithContext(context.Background(), service.ID)
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}

	return nil
}

func (p *Pagerduty) GetPagerDutyServiceByNameDirect(name string) (pagerduty.Service, bool, error) {
	// Go directly to PagerDuty to get the service
	var allServices []pagerduty.Service
	var offset uint = 0
	for {
		services, err := p.client.ListServicesPaginated(context.Background(), pagerduty.ListServiceOptions{Limit: 100, Offset: offset})
		if err != nil {
			log.Println("Failed to refresh PagerDuty service cache:", err)
			return pagerduty.Service{}, false, err
		}
		allServices = append(allServices, services...)
		if len(services) < 100 {
			break
		}
		offset += 100
	}

	for _, svc := range allServices {
		if svc.Name == name {
			return svc, true, nil
		}
	}

	return pagerduty.Service{}, false, nil
}
