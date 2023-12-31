// Code generated by Kitex v0.8.0. DO NOT EDIT.

package agentservice

import (
	"context"
	agent "github.com/Skyenought/tet/kitex_gen/agent"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddRepository(ctx context.Context, req *agent.AddRepositoryReq, callOptions ...callopt.Option) (r *agent.AddRepositoryRes, err error)
	DeleteRepositories(ctx context.Context, req *agent.DeleteRepositoriesReq, callOptions ...callopt.Option) (r *agent.DeleteRepositoriesRes, err error)
	UpdateRepository(ctx context.Context, req *agent.UpdateRepositoryReq, callOptions ...callopt.Option) (r *agent.UpdateRepositoryRes, err error)
	GetRepositories(ctx context.Context, req *agent.GetRepositoriesReq, callOptions ...callopt.Option) (r *agent.GetRepositoriesRes, err error)
	AddIDL(ctx context.Context, req *agent.AddIDLReq, callOptions ...callopt.Option) (r *agent.AddIDLRes, err error)
	DeleteIDL(ctx context.Context, req *agent.DeleteIDLsReq, callOptions ...callopt.Option) (r *agent.DeleteIDLsRes, err error)
	UpdateIDL(ctx context.Context, req *agent.UpdateIDLReq, callOptions ...callopt.Option) (r *agent.UpdateIDLRes, err error)
	GetIDLs(ctx context.Context, req *agent.GetIDLsReq, callOptions ...callopt.Option) (r *agent.GetIDLsRes, err error)
	SyncIDLsById(ctx context.Context, req *agent.SyncIDLsByIdReq, callOptions ...callopt.Option) (r *agent.SyncIDLsByIdRes, err error)
	AddTemplate(ctx context.Context, req *agent.AddTemplateReq, callOptions ...callopt.Option) (r *agent.AddTemplateRes, err error)
	DeleteTemplate(ctx context.Context, req *agent.DeleteTemplateReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateRes, err error)
	UpdateTemplate(ctx context.Context, req *agent.UpdateTemplateReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateRes, err error)
	GetTemplates(ctx context.Context, req *agent.GetTemplatesReq, callOptions ...callopt.Option) (r *agent.GetTemplatesRes, err error)
	AddTemplateItem(ctx context.Context, req *agent.AddTemplateItemReq, callOptions ...callopt.Option) (r *agent.AddTemplateItemRes, err error)
	DeleteTemplateItem(ctx context.Context, req *agent.DeleteTemplateItemReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateItemRes, err error)
	UpdateTemplateItem(ctx context.Context, req *agent.UpdateTemplateItemReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateItemRes, err error)
	GetTemplateItems(ctx context.Context, req *agent.GetTemplateItemsReq, callOptions ...callopt.Option) (r *agent.GetTemplateItemsRes, err error)
	UpdateTasks(ctx context.Context, req *agent.UpdateTasksReq, callOptions ...callopt.Option) (r *agent.UpdateTasksRes, err error)
	AddToken(ctx context.Context, req *agent.AddTokenReq, callOptions ...callopt.Option) (r *agent.AddTokenRes, err error)
	DeleteToken(ctx context.Context, req *agent.DeleteTokenReq, callOptions ...callopt.Option) (r *agent.DeleteTokenRes, err error)
	GetToken(ctx context.Context, req *agent.GetTokenReq, callOptions ...callopt.Option) (r *agent.GetTokenRes, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAgentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAgentServiceClient struct {
	*kClient
}

func (p *kAgentServiceClient) AddRepository(ctx context.Context, req *agent.AddRepositoryReq, callOptions ...callopt.Option) (r *agent.AddRepositoryRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddRepository(ctx, req)
}

func (p *kAgentServiceClient) DeleteRepositories(ctx context.Context, req *agent.DeleteRepositoriesReq, callOptions ...callopt.Option) (r *agent.DeleteRepositoriesRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRepositories(ctx, req)
}

func (p *kAgentServiceClient) UpdateRepository(ctx context.Context, req *agent.UpdateRepositoryReq, callOptions ...callopt.Option) (r *agent.UpdateRepositoryRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRepository(ctx, req)
}

func (p *kAgentServiceClient) GetRepositories(ctx context.Context, req *agent.GetRepositoriesReq, callOptions ...callopt.Option) (r *agent.GetRepositoriesRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRepositories(ctx, req)
}

func (p *kAgentServiceClient) AddIDL(ctx context.Context, req *agent.AddIDLReq, callOptions ...callopt.Option) (r *agent.AddIDLRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddIDL(ctx, req)
}

func (p *kAgentServiceClient) DeleteIDL(ctx context.Context, req *agent.DeleteIDLsReq, callOptions ...callopt.Option) (r *agent.DeleteIDLsRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteIDL(ctx, req)
}

func (p *kAgentServiceClient) UpdateIDL(ctx context.Context, req *agent.UpdateIDLReq, callOptions ...callopt.Option) (r *agent.UpdateIDLRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateIDL(ctx, req)
}

func (p *kAgentServiceClient) GetIDLs(ctx context.Context, req *agent.GetIDLsReq, callOptions ...callopt.Option) (r *agent.GetIDLsRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetIDLs(ctx, req)
}

func (p *kAgentServiceClient) SyncIDLsById(ctx context.Context, req *agent.SyncIDLsByIdReq, callOptions ...callopt.Option) (r *agent.SyncIDLsByIdRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SyncIDLsById(ctx, req)
}

func (p *kAgentServiceClient) AddTemplate(ctx context.Context, req *agent.AddTemplateReq, callOptions ...callopt.Option) (r *agent.AddTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTemplate(ctx, req)
}

func (p *kAgentServiceClient) DeleteTemplate(ctx context.Context, req *agent.DeleteTemplateReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTemplate(ctx, req)
}

func (p *kAgentServiceClient) UpdateTemplate(ctx context.Context, req *agent.UpdateTemplateReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTemplate(ctx, req)
}

func (p *kAgentServiceClient) GetTemplates(ctx context.Context, req *agent.GetTemplatesReq, callOptions ...callopt.Option) (r *agent.GetTemplatesRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTemplates(ctx, req)
}

func (p *kAgentServiceClient) AddTemplateItem(ctx context.Context, req *agent.AddTemplateItemReq, callOptions ...callopt.Option) (r *agent.AddTemplateItemRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTemplateItem(ctx, req)
}

func (p *kAgentServiceClient) DeleteTemplateItem(ctx context.Context, req *agent.DeleteTemplateItemReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateItemRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTemplateItem(ctx, req)
}

func (p *kAgentServiceClient) UpdateTemplateItem(ctx context.Context, req *agent.UpdateTemplateItemReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateItemRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTemplateItem(ctx, req)
}

func (p *kAgentServiceClient) GetTemplateItems(ctx context.Context, req *agent.GetTemplateItemsReq, callOptions ...callopt.Option) (r *agent.GetTemplateItemsRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTemplateItems(ctx, req)
}

func (p *kAgentServiceClient) UpdateTasks(ctx context.Context, req *agent.UpdateTasksReq, callOptions ...callopt.Option) (r *agent.UpdateTasksRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTasks(ctx, req)
}

func (p *kAgentServiceClient) AddToken(ctx context.Context, req *agent.AddTokenReq, callOptions ...callopt.Option) (r *agent.AddTokenRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddToken(ctx, req)
}

func (p *kAgentServiceClient) DeleteToken(ctx context.Context, req *agent.DeleteTokenReq, callOptions ...callopt.Option) (r *agent.DeleteTokenRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteToken(ctx, req)
}

func (p *kAgentServiceClient) GetToken(ctx context.Context, req *agent.GetTokenReq, callOptions ...callopt.Option) (r *agent.GetTokenRes, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetToken(ctx, req)
}
