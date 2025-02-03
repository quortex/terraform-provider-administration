package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetProcessingPresets - Returns a specifc preset.
func (c *Client) GetManageProcessingPresets(presetID string) (*ProcessingPresets, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	preset := ProcessingPresets{}
	err = json.Unmarshal(body, &preset)
	if err != nil {
		return nil, err
	}

	return &preset, nil
}

// CreateProcessingPresets - Create new preset.
func (c *Client) CreateManageProcessingPresets(preset ProcessingPresets) (*ProcessingPresets, error) {
	rb, err := json.Marshal(preset)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/manage/ott/presets/processings/", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rpreset := ProcessingPresets{}
	err = json.Unmarshal(body, &rpreset)
	if err != nil {
		return nil, err
	}

	return &rpreset, nil
}

// UpdateProcessingPresets - Updates a preset.
func (c *Client) UpdateManageProcessingPresets(presetID string, preset ProcessingPresets) (*ProcessingPresets, error) {
	rb, err := json.Marshal(preset)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	rpreset := ProcessingPresets{}
	err = json.Unmarshal(body, &rpreset)
	if err != nil {
		return nil, err
	}

	return &rpreset, nil
}

// DeleteProcessingPresets - Deletes a preset.
func (c *Client) DeleteManageProcessingPresets(presetID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
