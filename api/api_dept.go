package api

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"net/http"
	"net/url"
	"strconv"
)

// CreateDept:创建部门
// name:部门名称
// parentId:父部门id
func (ding *DingTalk) CreateDept(name string, parentId int) (rsp model.DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["name"] = name
	form["parentid"] = parentId

	err = ding.Request(http.MethodPost, constant.CreateDeptKey, nil, form, &rsp)

	return rsp, err
}

//CreateDetailDept:创建详细的部门
func (ding *DingTalk) CreateDetailDept(res model.CreateDetailDeptRequest) (req model.DeptCreateResponse, err error) {

	//组装参数
	res.JoinOuterPermitUsers()
	res.JoinOuterPermitDepts()
	res.JoinDeptPermits()
	res.JoinUserPermits()

	err = ding.Request(http.MethodPost, constant.CreateDeptKey, nil, res, &req)

	return req, err
}

//GetDeptDetail:获取部门详情
func (ding *DingTalk) GetDeptDetail(deptId int, lang string) (rsp model.DeptDetail, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))
	params.Set("lang", lang)

	err = ding.Request(http.MethodGet, constant.GetDeptDetailKey, params, nil, &rsp)

	return rsp, err
}

//DeleteDept:删除部门
func (ding *DingTalk) DeleteDept(deptId int) (rsp model.Response, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.DeleteDeptKey, params, nil, &rsp)

	return rsp, err
}

//UpdateDept:更新部门
func (ding *DingTalk) UpdateDept(res model.CreateDetailDeptRequest) (rsp model.DeptCreateResponse, err error) {

	//组装参数
	res.JoinOuterPermitUsers()
	res.JoinOuterPermitDepts()
	res.JoinDeptPermits()
	res.JoinUserPermits()
	res.JoinDeptManagerUserIds()

	err = ding.Request(http.MethodPost, constant.UpdateDeptKey, nil, res, &rsp)

	return rsp, err
}

//GetSubDeptList:获取子部门列表
func (ding *DingTalk) GetSubDeptList(deptId int, lang string, fetch bool) (rsp model.GetSubDeptResponse, err error) {

	if lang != "en_US" {
		lang = "zh_CN"
	}

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))
	params.Set("lang", lang)
	params.Set("fetch_child", strconv.FormatBool(fetch))

	err = ding.Request(http.MethodGet, constant.GetSubDeptListKey, params, nil, &rsp)

	return rsp, err
}

//GetDeptUserIds:获取部门用户userid列表
//deptId:部门id
func (ding *DingTalk) GetDeptUserIds(deptId int) (req model.DeptUserIdsResponse, err error) {

	params := url.Values{}
	params.Set("deptId", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetDeptUserIdKey, params, nil, &req)
	return req, err
}

//GetDeptUserDetail:获取部门用户详情
func (ding *DingTalk) GetDeptUserDetail(deptId, offset, size int, lang string) (req model.GetDeptUserDetailResponse, err error) {
	if lang != "en_US" {
		lang = "zh_CN"
	}

	if size < 0 || size > 100 {
		size = 100
	}
	params := url.Values{}
	params.Set("lang", lang)
	params.Set("department_id", strconv.Itoa(deptId))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("size", strconv.Itoa(size))
	params.Set("order", "entry_desc")

	err = ding.Request(http.MethodGet, constant.GetDeptUserDetailKey, params, nil, &req)
	return req, err
}

//GetDeptUserDetail:获取子部门ID列表
func (ding *DingTalk) GetSubDeptIds(deptId int) (req model.GetSubDeptIdsResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetSubDeptIdsKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询指定用户的所有上级父部门路径
func (ding *DingTalk) GetParentIdsByUserId(userId string) (req model.GetParentIdsByUserIdResponse, err error) {

	params := url.Values{}
	params.Set("userId", userId)

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByUserKey, params, nil, &req)
	return req, err
}

//GetParentIdsByUserId:查询部门的所有上级父部门路径
func (ding *DingTalk) GetParentIdsByDeptId(deptId int) (req model.GetParentIdsByDeptIdResponse, err error) {

	params := url.Values{}
	params.Set("id", strconv.Itoa(deptId))

	err = ding.Request(http.MethodGet, constant.GetParentDeptsByDeptKey, params, nil, &req)
	return req, err
}