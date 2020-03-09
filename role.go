package cas_client

func (cas *Cas) AddRole(code string, name string, description string) int {
	var response = struct {
		Response
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}{}
	cas.HttpPost(URL_ROLE_ADD, nil, RequestBody{
		"code":        code,
		"name":        name,
		"description": description,
	}, &response)
	return response.Data.ID
}

func (cas *Cas) UpdateRole(id int, role Role) bool {
	var response Response
	cas.HttpPost(URL_ROLE_UPDATE, nil, RequestBody{
		"id":          id,
		"code":        role.Code,
		"name":        role.Name,
		"description": role.Description,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) UpdateRoleByCode(code string, role Role) bool {
	var response Response
	cas.HttpPost(URL_ROLE_UPDATE_BY_CODE, nil, RequestBody{
		"code": code,
		"name": role.Name,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteRole(id int) bool {
	var response Response
	cas.HttpPost(URL_ROLE_DELETE, nil, RequestBody{
		"id": id,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetRole(id int) *Role {
	var response = struct {
		Response
		Role *Role `json:"data"`
	}{}
	cas.HttpGet(URL_ROLE_GET, Query{
		"id": id,
	}, &response)

	if response.Role == nil {
		return nil
	} else {
		return response.Role
	}
}

func (cas *Cas) GetRoleList(name string, page int, pageSize int) (int, []Role) {
	var response = struct {
		Response
		Data struct {
			Total int    `json:"total"`
			List  []Role `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ROLE_LIST, Query{
		"name":      name,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) AddRoleResource(roleIDList []int, resourceIDList []int) bool {
	var response Response
	cas.HttpPost(URL_ROLE_RESOURCE_ADD, nil, RequestBody{
		"role_id_list":     roleIDList,
		"resource_id_list": resourceIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteRoleResource(roleIDList []int, resourceIDList []int) bool {
	var response Response
	cas.HttpPost(URL_ROLE_RESOURCE_DELETE, nil, RequestBody{
		"role_id_list":     roleIDList,
		"resource_id_list": resourceIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetRoleResource(roleID int, page int, pageSize int) (int, []Resource) {
	var response = struct {
		Response
		Data struct {
			Total int        `json:"total"`
			List  []Resource `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ROLE_RESOURCE_LIST, Query{
		"role_id":   roleID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetResourceRole(resourceID int, page int, pageSize int) (int, []Role) {
	var response = struct {
		Response
		Data struct {
			Total int    `json:"total"`
			List  []Role `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_ROLE_LIST, Query{
		"resource_id": resourceID,
		"page":        page,
		"page_size":   pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) AddRoleAction(roleIDList []int, actionIDList []int) bool {
	var response Response
	cas.HttpPost(URL_ROLE_ACTION_ADD, nil, RequestBody{
		"role_id_list":   roleIDList,
		"action_id_list": actionIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteRoleAction(roleIDList []int, actionIDList []int) bool {
	var response Response
	cas.HttpPost(URL_ROLE_ACTION_DELETE, nil, RequestBody{
		"role_id_list":   roleIDList,
		"action_id_list": actionIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetRoleAction(roleID int, page int, pageSize int) (int, []Action) {
	var response = struct {
		Response
		Data struct {
			Total int      `json:"total"`
			List  []Action `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ROLE_ACTION_LIST, Query{
		"role_id":   roleID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetActionRole(actionID int, page int, pageSize int) (int, []Role) {
	var response = struct {
		Response
		Data struct {
			Total int    `json:"total"`
			List  []Role `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_ROLE_LIST, Query{
		"action_id": actionID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}
