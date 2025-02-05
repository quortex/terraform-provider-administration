package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func addParamToUrl(url string, param string, value string) string {
	if value != "" {
		if strings.Contains(url, "?") {
			return fmt.Sprintf("%s&%s=%s", url, param, value)
		}
		return fmt.Sprintf("%s?%s=%s", url, param, value)
	}
	return url
}

// GetProcessingPresets - Returns a specifc preset.
func (c *Client) GetManageProcessingPresets(presetUUID string, org string, ptype string) (*ProcessingPresets, error) {
	url := fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetUUID)

	url = addParamToUrl(url, "org", org)
	url = addParamToUrl(url, "type", ptype)

	req, err := http.NewRequest("GET", url, nil)
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
func (c *Client) CreateManageProcessingPresets(preset ProcessingPresets, org string, ptype string) (*ProcessingPresets, error) {
	rb, err := json.Marshal(preset)
	if err != nil {
		return nil, err
	}

	// tflog.Info(tflog.Debug, "Processing Preset: %s", string(rb))

	url := fmt.Sprintf("%s/1.0/manage/ott/presets/processings/", c.HostURL)
	url = addParamToUrl(url, "org", org)
	url = addParamToUrl(url, "type", ptype)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
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
func (c *Client) UpdateManageProcessingPresets(presetUUID string, preset ProcessingPresets, org string, ptype string) (*ProcessingPresets, error) {
	rb, err := json.Marshal(preset)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetUUID)
	url = addParamToUrl(url, "org", org)
	url = addParamToUrl(url, "type", ptype)

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(rb)))
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
func (c *Client) DeleteManageProcessingPresets(presetUUID string, org string, ptype string) error {
	url := fmt.Sprintf("%s/1.0/manage/ott/presets/processings/%s", c.HostURL, presetUUID)

	url = addParamToUrl(url, "org", org)
	url = addParamToUrl(url, "type", ptype)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
