package pd

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/log"

	"strings"

	pagerdutyv1beta1 "github.com/mattgialelis/dutycontroller/api/v1beta1"

	"github.com/PagerDuty/go-pagerduty"
)

type Service struct {
	Name               string
	Description        string
	Status             string
	EscalationPolicyID string
	AutoResolveTimeout int
	AcknowedgeTimeout  int
}

func ServicesSpectoService(bs pagerdutyv1beta1.Services, EscalationPolicyID string) Service {
	return Service{
		Name:               bs.Name,
		Description:        bs.Spec.Description,
		Status:             bs.Spec.Status,
		EscalationPolicyID: EscalationPolicyID,
		AutoResolveTimeout: bs.Spec.AutoResolveTimeout,
		AcknowedgeTimeout:  bs.Spec.AcknowedgeTimeout,
	}
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

	existingService, _, err := p.GetPagerDutyServiceByNameDirect(context.TODO(), service.Name)
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

func (p *Pagerduty) DeletePagerDutyService(ctx context.Context, id string) error {
	log := log.FromContext(ctx)

	err := p.client.DeleteServiceWithContext(context.Background(), id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Info("Service not found in PagerDdy, skipping deletion, no need to delete")
			return nil
		} else {
			return fmt.Errorf("failed to delete service: %w", err)
		}
	}

	return nil
}

func (p *Pagerduty) GetPagerDutyServiceByNameDirect(ctx context.Context, name string) (pagerduty.Service, bool, error) {
	// Go directly to PagerDuty to get the service
	var allServices []pagerduty.Service
	var offset uint = 0

	log := log.FromContext(ctx)

	for {
		services, err := p.client.ListServicesPaginated(
			context.Background(),
			pagerduty.ListServiceOptions{Limit: 100, Offset: offset},
		)
		if err != nil {
			log.Info("Failed to refresh PagerDuty service cache:", err)
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
