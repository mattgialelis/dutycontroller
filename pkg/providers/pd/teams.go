package pd

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
)

func (pd *Pagerduty) GetTeambyName(name string) (string, bool, error) {

	teamResp, err := pd.client.ListTeamsWithContext(context.Background(), pagerduty.ListTeamOptions{})
	if err != nil {
		return "", false, err
	}

	for _, t := range teamResp.Teams {
		if t.Name == name {
			return t.ID, true, nil
		}
	}

	return "", false, nil
}
