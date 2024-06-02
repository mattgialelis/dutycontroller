package pd

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/PagerDuty/go-pagerduty"
)

type Pagerduty struct {
	client       *pagerduty.Client
	serivceCache []pagerduty.Service
	refresh      time.Duration
	mu           sync.Mutex
}

// NewPagerduty creates a new Pagerduty client
// Input:
//
//	authToken:  PagerDuty API token
//	refreshInterval:  Interval to refresh service the cache
//
// Returns:
//
//	Pagerduty:  Pagerduty client
//	error:  Error
func NewPagerduty(authToken string, refreshInterval int) (*Pagerduty, error) {
	if authToken == "" {
		return nil, fmt.Errorf("PAGERDUTY_TOKEN environment variable not set")
	}

	client := pagerduty.NewClient(authToken)

	pd := &Pagerduty{
		client:  client,
		refresh: time.Duration(refreshInterval) * time.Second,
	}

	return pd, nil
}

func (pd *Pagerduty) RefreshCache(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	ticker := time.NewTicker(pd.refresh)
	defer ticker.Stop()

	// Refresh the cache immediately when the function is started
	// Assume refresh() now returns an error
	err := pd.refreshCache()
	if err != nil {
		log.Println("Error refreshing PagerDuty service cache:", err)
		wg.Done()
		return
	}
	wg.Done()

	for {
		select {
		case <-ticker.C:
			err := pd.refreshCache()
			if err != nil {
				log.Println("Error refreshing PagerDuty service cache:", err)
				return
			}
		case <-stopCh:
			log.Println("Stopping refresh of PagerDuty service cache")
			return
		}
	}
}

func (pd *Pagerduty) refreshCache() error {
	var allServices []pagerduty.Service
	var offset uint = 0
	// Lock the mutex to ensure cache refresh is atomic
	pd.mu.Lock()
	defer pd.mu.Unlock()

	for {
		services, err := pd.client.ListServicesPaginated(
			context.Background(),
			pagerduty.ListServiceOptions{Limit: 100, Offset: offset},
		)

		if err != nil {
			log.Println("Failed to refresh PagerDuty service cache:", err)
			return err
		}
		allServices = append(allServices, services...)
		if len(services) < 100 {
			break
		}
		offset += 100
	}
	pd.serivceCache = allServices
	log.Printf("Synced %d services from PagerDuty\n", len(allServices))
	return nil
}
