package tet

import (
	"context"
	agent "github.com/Skyenought/tet/kitex_gen/agent"

	"github.com/Skyenought/tet/kitex_gen/agent/agentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() agentservice.Client
	Service() string
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

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := agentservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient agentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() agentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddRepository(ctx context.Context, req *agent.AddRepositoryReq, callOptions ...callopt.Option) (r *agent.AddRepositoryRes, err error) {
	return c.kitexClient.AddRepository(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteRepositories(ctx context.Context, req *agent.DeleteRepositoriesReq, callOptions ...callopt.Option) (r *agent.DeleteRepositoriesRes, err error) {
	return c.kitexClient.DeleteRepositories(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateRepository(ctx context.Context, req *agent.UpdateRepositoryReq, callOptions ...callopt.Option) (r *agent.UpdateRepositoryRes, err error) {
	return c.kitexClient.UpdateRepository(ctx, req, callOptions...)
}

func (c *clientImpl) GetRepositories(ctx context.Context, req *agent.GetRepositoriesReq, callOptions ...callopt.Option) (r *agent.GetRepositoriesRes, err error) {
	return c.kitexClient.GetRepositories(ctx, req, callOptions...)
}

func (c *clientImpl) AddIDL(ctx context.Context, req *agent.AddIDLReq, callOptions ...callopt.Option) (r *agent.AddIDLRes, err error) {
	return c.kitexClient.AddIDL(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteIDL(ctx context.Context, req *agent.DeleteIDLsReq, callOptions ...callopt.Option) (r *agent.DeleteIDLsRes, err error) {
	return c.kitexClient.DeleteIDL(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateIDL(ctx context.Context, req *agent.UpdateIDLReq, callOptions ...callopt.Option) (r *agent.UpdateIDLRes, err error) {
	return c.kitexClient.UpdateIDL(ctx, req, callOptions...)
}

func (c *clientImpl) GetIDLs(ctx context.Context, req *agent.GetIDLsReq, callOptions ...callopt.Option) (r *agent.GetIDLsRes, err error) {
	return c.kitexClient.GetIDLs(ctx, req, callOptions...)
}

func (c *clientImpl) SyncIDLsById(ctx context.Context, req *agent.SyncIDLsByIdReq, callOptions ...callopt.Option) (r *agent.SyncIDLsByIdRes, err error) {
	return c.kitexClient.SyncIDLsById(ctx, req, callOptions...)
}

func (c *clientImpl) AddTemplate(ctx context.Context, req *agent.AddTemplateReq, callOptions ...callopt.Option) (r *agent.AddTemplateRes, err error) {
	return c.kitexClient.AddTemplate(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteTemplate(ctx context.Context, req *agent.DeleteTemplateReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateRes, err error) {
	return c.kitexClient.DeleteTemplate(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateTemplate(ctx context.Context, req *agent.UpdateTemplateReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateRes, err error) {
	return c.kitexClient.UpdateTemplate(ctx, req, callOptions...)
}

func (c *clientImpl) GetTemplates(ctx context.Context, req *agent.GetTemplatesReq, callOptions ...callopt.Option) (r *agent.GetTemplatesRes, err error) {
	return c.kitexClient.GetTemplates(ctx, req, callOptions...)
}

func (c *clientImpl) AddTemplateItem(ctx context.Context, req *agent.AddTemplateItemReq, callOptions ...callopt.Option) (r *agent.AddTemplateItemRes, err error) {
	return c.kitexClient.AddTemplateItem(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteTemplateItem(ctx context.Context, req *agent.DeleteTemplateItemReq, callOptions ...callopt.Option) (r *agent.DeleteTemplateItemRes, err error) {
	return c.kitexClient.DeleteTemplateItem(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateTemplateItem(ctx context.Context, req *agent.UpdateTemplateItemReq, callOptions ...callopt.Option) (r *agent.UpdateTemplateItemRes, err error) {
	return c.kitexClient.UpdateTemplateItem(ctx, req, callOptions...)
}

func (c *clientImpl) GetTemplateItems(ctx context.Context, req *agent.GetTemplateItemsReq, callOptions ...callopt.Option) (r *agent.GetTemplateItemsRes, err error) {
	return c.kitexClient.GetTemplateItems(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateTasks(ctx context.Context, req *agent.UpdateTasksReq, callOptions ...callopt.Option) (r *agent.UpdateTasksRes, err error) {
	return c.kitexClient.UpdateTasks(ctx, req, callOptions...)
}

func (c *clientImpl) AddToken(ctx context.Context, req *agent.AddTokenReq, callOptions ...callopt.Option) (r *agent.AddTokenRes, err error) {
	return c.kitexClient.AddToken(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteToken(ctx context.Context, req *agent.DeleteTokenReq, callOptions ...callopt.Option) (r *agent.DeleteTokenRes, err error) {
	return c.kitexClient.DeleteToken(ctx, req, callOptions...)
}

func (c *clientImpl) GetToken(ctx context.Context, req *agent.GetTokenReq, callOptions ...callopt.Option) (r *agent.GetTokenRes, err error) {
	return c.kitexClient.GetToken(ctx, req, callOptions...)
}
