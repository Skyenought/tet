// Code generated by thriftgo (0.3.4). DO NOT EDIT.

package agent

import (
	"fmt"
	"github.com/Skyenought/tet/kitex_gen/model"
)

type AddRepositoryReq struct {
	RepositoryType   int32  `thrift:"repository_type,1" frugal:"1,default,i32" json:"repository_type"`
	RepositoryDomain string `thrift:"repository_domain,2" frugal:"2,default,string" json:"repository_domain"`
	RepositoryOwner  string `thrift:"repository_owner,3" frugal:"3,default,string" json:"repository_owner"`
	RepositoryName   string `thrift:"repository_name,4" frugal:"4,default,string" json:"repository_name"`
	Branch           string `thrift:"branch,5" frugal:"5,default,string" json:"branch"`
	StoreType        int32  `thrift:"store_type,6" frugal:"6,default,i32" json:"store_type"`
}

func NewAddRepositoryReq() *AddRepositoryReq {
	return &AddRepositoryReq{}
}

func (p *AddRepositoryReq) InitDefault() {
	*p = AddRepositoryReq{}
}

func (p *AddRepositoryReq) GetRepositoryType() (v int32) {
	return p.RepositoryType
}

func (p *AddRepositoryReq) GetRepositoryDomain() (v string) {
	return p.RepositoryDomain
}

func (p *AddRepositoryReq) GetRepositoryOwner() (v string) {
	return p.RepositoryOwner
}

func (p *AddRepositoryReq) GetRepositoryName() (v string) {
	return p.RepositoryName
}

func (p *AddRepositoryReq) GetBranch() (v string) {
	return p.Branch
}

func (p *AddRepositoryReq) GetStoreType() (v int32) {
	return p.StoreType
}
func (p *AddRepositoryReq) SetRepositoryType(val int32) {
	p.RepositoryType = val
}
func (p *AddRepositoryReq) SetRepositoryDomain(val string) {
	p.RepositoryDomain = val
}
func (p *AddRepositoryReq) SetRepositoryOwner(val string) {
	p.RepositoryOwner = val
}
func (p *AddRepositoryReq) SetRepositoryName(val string) {
	p.RepositoryName = val
}
func (p *AddRepositoryReq) SetBranch(val string) {
	p.Branch = val
}
func (p *AddRepositoryReq) SetStoreType(val int32) {
	p.StoreType = val
}

func (p *AddRepositoryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryReq(%+v)", *p)
}

type AddRepositoryRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewAddRepositoryRes() *AddRepositoryRes {
	return &AddRepositoryRes{}
}

func (p *AddRepositoryRes) InitDefault() {
	*p = AddRepositoryRes{}
}

func (p *AddRepositoryRes) GetCode() (v int32) {
	return p.Code
}

func (p *AddRepositoryRes) GetMsg() (v string) {
	return p.Msg
}
func (p *AddRepositoryRes) SetCode(val int32) {
	p.Code = val
}
func (p *AddRepositoryRes) SetMsg(val string) {
	p.Msg = val
}

func (p *AddRepositoryRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryRes(%+v)", *p)
}

type DeleteRepositoriesReq struct {
	Ids []int64 `thrift:"ids,1" frugal:"1,default,list<i64>" json:"ids"`
}

func NewDeleteRepositoriesReq() *DeleteRepositoriesReq {
	return &DeleteRepositoriesReq{}
}

func (p *DeleteRepositoriesReq) InitDefault() {
	*p = DeleteRepositoriesReq{}
}

func (p *DeleteRepositoriesReq) GetIds() (v []int64) {
	return p.Ids
}
func (p *DeleteRepositoriesReq) SetIds(val []int64) {
	p.Ids = val
}

func (p *DeleteRepositoriesReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteRepositoriesReq(%+v)", *p)
}

type DeleteRepositoriesRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewDeleteRepositoriesRes() *DeleteRepositoriesRes {
	return &DeleteRepositoriesRes{}
}

func (p *DeleteRepositoriesRes) InitDefault() {
	*p = DeleteRepositoriesRes{}
}

func (p *DeleteRepositoriesRes) GetCode() (v int32) {
	return p.Code
}

func (p *DeleteRepositoriesRes) GetMsg() (v string) {
	return p.Msg
}
func (p *DeleteRepositoriesRes) SetCode(val int32) {
	p.Code = val
}
func (p *DeleteRepositoriesRes) SetMsg(val string) {
	p.Msg = val
}

func (p *DeleteRepositoriesRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteRepositoriesRes(%+v)", *p)
}

type UpdateRepositoryReq struct {
	Id     int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Branch string `thrift:"branch,2" frugal:"2,default,string" json:"branch"`
	Status int32  `thrift:"status,3" frugal:"3,default,i32" json:"status"`
}

func NewUpdateRepositoryReq() *UpdateRepositoryReq {
	return &UpdateRepositoryReq{}
}

func (p *UpdateRepositoryReq) InitDefault() {
	*p = UpdateRepositoryReq{}
}

func (p *UpdateRepositoryReq) GetId() (v int64) {
	return p.Id
}

func (p *UpdateRepositoryReq) GetBranch() (v string) {
	return p.Branch
}

func (p *UpdateRepositoryReq) GetStatus() (v int32) {
	return p.Status
}
func (p *UpdateRepositoryReq) SetId(val int64) {
	p.Id = val
}
func (p *UpdateRepositoryReq) SetBranch(val string) {
	p.Branch = val
}
func (p *UpdateRepositoryReq) SetStatus(val int32) {
	p.Status = val
}

func (p *UpdateRepositoryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryReq(%+v)", *p)
}

type UpdateRepositoryRes struct {
	Code int32  `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
}

func NewUpdateRepositoryRes() *UpdateRepositoryRes {
	return &UpdateRepositoryRes{}
}

func (p *UpdateRepositoryRes) InitDefault() {
	*p = UpdateRepositoryRes{}
}

func (p *UpdateRepositoryRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateRepositoryRes) GetMsg() (v string) {
	return p.Msg
}
func (p *UpdateRepositoryRes) SetCode(val int32) {
	p.Code = val
}
func (p *UpdateRepositoryRes) SetMsg(val string) {
	p.Msg = val
}

func (p *UpdateRepositoryRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryRes(%+v)", *p)
}

type GetRepositoriesReq struct {
	Page             int32  `thrift:"page,1" frugal:"1,default,i32" json:"page"`
	Limit            int32  `thrift:"limit,2" frugal:"2,default,i32" json:"limit"`
	Order            int32  `thrift:"order,3" frugal:"3,default,i32" json:"order"`
	OrderBy          string `thrift:"order_by,4" frugal:"4,default,string" json:"order_by"`
	RepositoryType   int32  `thrift:"repository_type,5" frugal:"5,default,i32" json:"repository_type"`
	StoreType        int32  `thrift:"store_type,6" frugal:"6,default,i32" json:"store_type"`
	RepositoryDomain string `thrift:"repository_domain,7" frugal:"7,default,string" json:"repository_domain"`
	RepositoryOwner  string `thrift:"repository_owner,8" frugal:"8,default,string" json:"repository_owner"`
	RepositoryName   string `thrift:"repository_name,9" frugal:"9,default,string" json:"repository_name"`
}

func NewGetRepositoriesReq() *GetRepositoriesReq {
	return &GetRepositoriesReq{}
}

func (p *GetRepositoriesReq) InitDefault() {
	*p = GetRepositoriesReq{}
}

func (p *GetRepositoriesReq) GetPage() (v int32) {
	return p.Page
}

func (p *GetRepositoriesReq) GetLimit() (v int32) {
	return p.Limit
}

func (p *GetRepositoriesReq) GetOrder() (v int32) {
	return p.Order
}

func (p *GetRepositoriesReq) GetOrderBy() (v string) {
	return p.OrderBy
}

func (p *GetRepositoriesReq) GetRepositoryType() (v int32) {
	return p.RepositoryType
}

func (p *GetRepositoriesReq) GetStoreType() (v int32) {
	return p.StoreType
}

func (p *GetRepositoriesReq) GetRepositoryDomain() (v string) {
	return p.RepositoryDomain
}

func (p *GetRepositoriesReq) GetRepositoryOwner() (v string) {
	return p.RepositoryOwner
}

func (p *GetRepositoriesReq) GetRepositoryName() (v string) {
	return p.RepositoryName
}
func (p *GetRepositoriesReq) SetPage(val int32) {
	p.Page = val
}
func (p *GetRepositoriesReq) SetLimit(val int32) {
	p.Limit = val
}
func (p *GetRepositoriesReq) SetOrder(val int32) {
	p.Order = val
}
func (p *GetRepositoriesReq) SetOrderBy(val string) {
	p.OrderBy = val
}
func (p *GetRepositoriesReq) SetRepositoryType(val int32) {
	p.RepositoryType = val
}
func (p *GetRepositoriesReq) SetStoreType(val int32) {
	p.StoreType = val
}
func (p *GetRepositoriesReq) SetRepositoryDomain(val string) {
	p.RepositoryDomain = val
}
func (p *GetRepositoriesReq) SetRepositoryOwner(val string) {
	p.RepositoryOwner = val
}
func (p *GetRepositoriesReq) SetRepositoryName(val string) {
	p.RepositoryName = val
}

func (p *GetRepositoriesReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesReq(%+v)", *p)
}

type GetRepositoriesRes struct {
	Code int32                   `thrift:"code,1" frugal:"1,default,i32" json:"code"`
	Msg  string                  `thrift:"msg,2" frugal:"2,default,string" json:"msg"`
	Data *GetRepositoriesResData `thrift:"data,3" frugal:"3,default,GetRepositoriesResData" json:"data"`
}

func NewGetRepositoriesRes() *GetRepositoriesRes {
	return &GetRepositoriesRes{}
}

func (p *GetRepositoriesRes) InitDefault() {
	*p = GetRepositoriesRes{}
}

func (p *GetRepositoriesRes) GetCode() (v int32) {
	return p.Code
}

func (p *GetRepositoriesRes) GetMsg() (v string) {
	return p.Msg
}

var GetRepositoriesRes_Data_DEFAULT *GetRepositoriesResData

func (p *GetRepositoriesRes) GetData() (v *GetRepositoriesResData) {
	if !p.IsSetData() {
		return GetRepositoriesRes_Data_DEFAULT
	}
	return p.Data
}
func (p *GetRepositoriesRes) SetCode(val int32) {
	p.Code = val
}
func (p *GetRepositoriesRes) SetMsg(val string) {
	p.Msg = val
}
func (p *GetRepositoriesRes) SetData(val *GetRepositoriesResData) {
	p.Data = val
}

func (p *GetRepositoriesRes) IsSetData() bool {
	return p.Data != nil
}

func (p *GetRepositoriesRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesRes(%+v)", *p)
}

type GetRepositoriesResData struct {
	Repositories []*model.Repository `thrift:"repositories,1" frugal:"1,default,list<model.Repository>" json:"repositories"`
	Total        int32               `thrift:"total,2" frugal:"2,default,i32" json:"total"`
}

func NewGetRepositoriesResData() *GetRepositoriesResData {
	return &GetRepositoriesResData{}
}

func (p *GetRepositoriesResData) InitDefault() {
	*p = GetRepositoriesResData{}
}

func (p *GetRepositoriesResData) GetRepositories() (v []*model.Repository) {
	return p.Repositories
}

func (p *GetRepositoriesResData) GetTotal() (v int32) {
	return p.Total
}
func (p *GetRepositoriesResData) SetRepositories(val []*model.Repository) {
	p.Repositories = val
}
func (p *GetRepositoriesResData) SetTotal(val int32) {
	p.Total = val
}

func (p *GetRepositoriesResData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesResData(%+v)", *p)
}
