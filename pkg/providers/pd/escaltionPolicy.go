package pd

import (
	"context"
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
)

func (p *Pagerduty) GetEscalationPolicyByName(name string) (string, bool, error) {
	policies, err := p.client.ListEscalationPoliciesWithContext(context.TODO(), pagerduty.ListEscalationPoliciesOptions{})
	if err != nil {
		return "", false, fmt.Errorf("failed to fetch escalation policies: %w", err)
	}

	if len(policies.EscalationPolicies) == 0 {
		return "", false, fmt.Errorf("no escalation policy returned by listEscalationPolicies")
	}

	for _, policy := range policies.EscalationPolicies {
		if policy.Name == name {
			return policy.ID, true, nil
		}
	}

	return "", false, fmt.Errorf("escalation policy with name %s not found", name)
}
