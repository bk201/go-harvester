package goharv

import (
	"encoding/json"
	"net/http"

	"github.com/rancher/apiserver/pkg/types"
	corev1 "k8s.io/api/core/v1"
)

type Node corev1.Node

type NodeList struct {
	types.Collection
	Data []*Node `json:"data"`
}

type NodesClient struct {
	*apiClient
}

func newNodesClient(c *Client) *NodesClient {
	return &NodesClient{
		apiClient: newAPIClient(c, "nodes"),
	}
}

func (s *NodesClient) List() (*NodeList, error) {
	var collection NodeList
	respCode, respBody, err := s.apiClient.List()
	if err != nil {
		return nil, err
	}
	if respCode != http.StatusOK {
		return nil, NewResponseError(respCode, respBody)
	}
	err = json.Unmarshal(respBody, &collection)
	return &collection, err
}

func (s *NodesClient) Get(name string) (*Node, error) {
	var obj *Node
	respCode, respBody, err := s.apiClient.Get(name)
	if err != nil {
		return nil, err
	}
	if respCode != http.StatusOK {
		return nil, NewResponseError(respCode, respBody)
	}
	if err = json.Unmarshal(respBody, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}