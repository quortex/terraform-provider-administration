package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetPlan - Returns a specifc plan.
func (c *Client) GetPlan(planID string) (*Plan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/manage/billing/plans/%s", c.HostURL, planID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	plan := Plan{}
	err = json.Unmarshal(body, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

// CreatePlan - Create new plan.
func (c *Client) CreatePlan(plan Plan) (*Plan, error) {
	rb, err := json.Marshal(plan)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/manage/billing/plans", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rplan := Plan{}
	err = json.Unmarshal(body, &rplan)
	if err != nil {
		return nil, err
	}

	return &rplan, nil
}

// UpdatePlan - Updates a plan.
func (c *Client) UpdatePlan(planID string, plan Plan) (*Plan, error) {
	rb, err := json.Marshal(plan)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/1.0/manage/billing/plans/%s", c.HostURL, planID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rplan := Plan{}
	err = json.Unmarshal(body, &rplan)
	if err != nil {
		return nil, err
	}

	return &rplan, nil
}

// DeletePlan - Deletes a plan.
func (c *Client) DeletePlan(planID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/manage/billing/plans/%s", c.HostURL, planID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
