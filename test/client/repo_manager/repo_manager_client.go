// Code generated by go-swagger; DO NOT EDIT.

package repo_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repo manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repo manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRepo create repo API
*/
func (a *Client) CreateRepo(params *CreateRepoParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRepo",
		Method:             "POST",
		PathPattern:        "/api/RepoManager.CreateRepo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRepoOK), nil

}

/*
DeleteRepos delete repos API
*/
func (a *Client) DeleteRepos(params *DeleteReposParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRepos",
		Method:             "POST",
		PathPattern:        "/api/RepoManager.DeleteRepos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteReposOK), nil

}

/*
DescribeRepos describe repos API
*/
func (a *Client) DescribeRepos(params *DescribeReposParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeRepos",
		Method:             "GET",
		PathPattern:        "/api/RepoManager.DescribeRepos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeReposOK), nil

}

/*
ModifyRepo modify repo API
*/
func (a *Client) ModifyRepo(params *ModifyRepoParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyRepo",
		Method:             "POST",
		PathPattern:        "/api/RepoManager.ModifyRepo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyRepoOK), nil

}

/*
ValidateRepo validate repo API
*/
func (a *Client) ValidateRepo(params *ValidateRepoParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ValidateRepo",
		Method:             "GET",
		PathPattern:        "/api/RepoManager.ValidateRepo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateRepoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateRepoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
