package pd

import (
	"context"
	"errors"

	"github.com/PagerDuty/go-pagerduty"
)

type OrchestrationRoute struct {
	EventOrchestrationName string
	Label                  string
	Expression             []string
	RouteTo                string
}

func (or OrchestrationRoute) ConditionsFromOrchestrationRoute() []*pagerduty.OrchestrationRouterRuleCondition {
	conditions := make([]*pagerduty.OrchestrationRouterRuleCondition, 0)
	for _, e := range or.Expression {
		conditions = append(conditions, &pagerduty.OrchestrationRouterRuleCondition{
			Expression: e,
		})
	}

	return conditions
}

func (or OrchestrationRoute) RuleFromOrchestrationRoute() *pagerduty.OrchestrationRouterRule {
	return &pagerduty.OrchestrationRouterRule{
		Label:      or.Label,
		Conditions: or.ConditionsFromOrchestrationRoute(),
		Actions: &pagerduty.OrchestrationRouterActions{
			RouteTo: or.RouteTo,
		},
	}
}

func (or *OrchestrationRoute) GenerateUpdateFromOrchestrationRouterRule(rule *pagerduty.OrchestrationRouterRule) {
	if or.Label == "" {
		or.Label = rule.Label
	}
	if or.RouteTo == "" {
		or.RouteTo = rule.Actions.RouteTo
	}
	if len(or.Expression) == 0 {
		for _, c := range rule.Conditions {
			or.Expression = append(or.Expression, c.Expression)
		}
	}
}

// Get Orchestration by Name
func (pd *Pagerduty) GetEventOrchestrationByName(name string) (string, bool, error) {
	orch, err := pd.client.ListOrchestrationsWithContext(context.Background(), pagerduty.ListOrchestrationsOptions{})
	if err != nil {
		return "", false, err
	}

	for _, o := range orch.Orchestrations {
		if o.Name == name {
			return o.ID, true, nil
		}
	}

	return "", false, nil
}

// Add new Rule to the existing Orchestrations
func (pd *Pagerduty) AddOrchestrationServiceRoute(OrchestrationRoute OrchestrationRoute) error {
	orchID, ok, err := pd.GetEventOrchestrationByName(OrchestrationRoute.EventOrchestrationName)
	if err != nil {
		return err
	} else {
		if !ok {
			return errors.New("orchestration not found")
		}
	}

	//Load existing Routes for the Orchestrations
	router, err := pd.client.GetOrchestrationRouterWithContext(context.Background(), orchID, &pagerduty.GetOrchestrationRouterOptions{})
	if err != nil {
		return err
	}

	rule := OrchestrationRoute.RuleFromOrchestrationRoute()

	//Add new Rule to the existing Orchestrations
	router.Sets[0].Rules = append(router.Sets[0].Rules, rule)

	//Apply the new Rule to the Orchestrations
	_, err = pd.client.UpdateOrchestrationRouterWithContext(context.Background(), orchID, *router)
	if err != nil {
		return err
	}

	return nil
}

// Func Update
func (pd *Pagerduty) UpdateOrchestrationServiceRoute(OrchestrationRoute OrchestrationRoute) error {
	orchID, ok, err := pd.GetEventOrchestrationByName(OrchestrationRoute.EventOrchestrationName)
	if err != nil {
		return err
	} else {
		if !ok {
			return errors.New("orchestration not found")
		}
	}

	//Load existing Routes for the Orchestrations
	router, err := pd.client.GetOrchestrationRouterWithContext(context.Background(), orchID, &pagerduty.GetOrchestrationRouterOptions{})
	if err != nil {
		return err
	}

	for i, rule := range router.Sets[0].Rules {
		if rule.Actions.RouteTo == OrchestrationRoute.RouteTo {
			OrchestrationRoute.GenerateUpdateFromOrchestrationRouterRule(router.Sets[0].Rules[i])
			router.Sets[0].Rules[i] = OrchestrationRoute.RuleFromOrchestrationRoute()
		}
	}

	_, err = pd.client.UpdateOrchestrationRouterWithContext(context.Background(), orchID, *router)
	if err != nil {
		return err
	}

	return nil
}

// Func Delete
func (pd *Pagerduty) DeleteOrchestrationServiceRoute(OrchestrationRoute OrchestrationRoute) error {
	orchID, ok, err := pd.GetEventOrchestrationByName(OrchestrationRoute.EventOrchestrationName)
	if err != nil {
		return err
	} else {
		if !ok {
			return errors.New("orchestration not found")
		}
	}

	//Load existing Routes for the Orchestrations
	router, err := pd.client.GetOrchestrationRouterWithContext(context.Background(), orchID, &pagerduty.GetOrchestrationRouterOptions{})
	if err != nil {
		return err
	}

	for i, rule := range router.Sets[0].Rules {
		if rule.Actions.RouteTo == OrchestrationRoute.RouteTo {
			// Remove the matching rule
			router.Sets[0].Rules = append(router.Sets[0].Rules[:i], router.Sets[0].Rules[i+1:]...)
			break
		}
	}

	_, err = pd.client.UpdateOrchestrationRouterWithContext(context.Background(), orchID, *router)
	if err != nil {
		return err
	}

	return nil
}
