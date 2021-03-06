package addontester

import (
	"fmt"

	"github.com/bitrise-io/bitrise-addon-test/addonprovisioner"
	"github.com/bitrise-io/bitrise-addon-test/utils"
)

// ChangePlanTesterParams ...
type ChangePlanTesterParams struct {
	AppSlug   string
	Plan      string
	WithRetry bool
}

// ChangePlan ...
func (t *Tester) ChangePlan(params ChangePlanTesterParams, remainingRetries int) error {
	if len(params.AppSlug) == 0 {
		var err error
		params.AppSlug, err = utils.RandomHex(8)
		if err != nil {
			return fmt.Errorf("Failed to generate app slug: %s", err)
		}
	}

	t.logger.Printf("\nPlan changing details:")
	t.logger.Printf("App slug: %s", params.AppSlug)
	t.logger.Printf("Plan: %s", params.Plan)
	t.logger.Printf("Should retry: %v", params.WithRetry)
	if params.WithRetry {
		t.logger.Printf("No. of test: %d.", numberOfTestsWithRetry-remainingRetries)
	}

	status, body, err := t.addonClient.ChangePlan(addonprovisioner.ChangePlanRequestParams{
		Plan: params.Plan,
	}, params.AppSlug)

	if err != nil {
		return fmt.Errorf("Plan changing failed: %s", err)
	}

	t.logger.Printf("\nResponse status: %d", status)
	t.logger.Printf("Response body: %v\n", body)

	if status < 200 || status > 299 {
		return fmt.Errorf("Plan changing request resulted in a non-2xx response")
	}

	t.logger.Println("\nPlan changing success.")

	if params.WithRetry && remainingRetries > 0 {
		return t.ChangePlan(params, remainingRetries-1)
	}

	return nil
}
