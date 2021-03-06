package cfclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type StacksResponse struct {
	Count     int              `json:"total_results"`
	Pages     int              `json:"total_pages"`
	NextUrl   string           `json:"next_url"`
	Resources []StacksResource `json:"resources"`
}

type StacksResource struct {
	Meta   Meta  `json:"metadata"`
	Entity Stack `json:"entity"`
}

type Stack struct {
	Guid        string `json:"guid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	c           *Client
}

func (c *Client) ListStacksByQuery(query url.Values) ([]Stack, error) {
	var stacks []Stack
	requestUrl := "/v2/stacks?" + query.Encode()
	for {
		stacksResp, err := c.getStacksResponse(requestUrl)
		if err != nil {
			return []Stack{}, err
		}
		for _, stack := range stacksResp.Resources {
			stack.Entity.Guid = stack.Meta.Guid
			stack.Entity.c = c
			stacks = append(stacks, stack.Entity)
		}
		requestUrl = stacksResp.NextUrl
		if requestUrl == "" {
			break
		}
	}
	return stacks, nil
}

func (c *Client) ListStacks() ([]Stack, error) {
	return c.ListStacksByQuery(nil)
}

func (c *Client) getStacksResponse(requestUrl string) (StacksResponse, error) {
	var stacksResp StacksResponse
	r := c.NewRequest("GET", requestUrl)
	resp, err := c.DoRequest(r)
	if err != nil {
		return StacksResponse{}, fmt.Errorf("Error requesting stacks %v", err)
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return StacksResponse{}, fmt.Errorf("Error reading stacks body %v", err)
	}
	err = json.Unmarshal(resBody, &stacksResp)
	if err != nil {
		return StacksResponse{}, fmt.Errorf("Error unmarshalling stacks %v", err)
	}
	return stacksResp, nil
}
