package cas_client

func (cas *Cas) AddUserRole(userIDList []int, roleIDList []int) bool {
	var response Response
	cas.HttpPost(URL_USER_ROLE_ADD, nil, RequestBody{
		"user_id_list": userIDList,
		"role_id_list": roleIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteUserRole(userIDList []int, roleIDList []int) bool {
	var response Response
	cas.HttpPost(URL_USER_ROLE_DELETE, nil, RequestBody{
		"user_id_list": userIDList,
		"role_id_list": roleIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) ClearUserRole(userID int) bool {
	var response Response
	cas.HttpPost(URL_USER_ROLE_CLEAR, nil, RequestBody{"user_id": userID}, &response)
	return response.Code == 200
}

func (cas *Cas) GetUserRole(userID int, page int, pageSize int) (int, []Role) {
	var response = struct {
		Response
		Data struct {
			Total int    `json:"total"`
			List  []Role `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_USER_ROLE_LIST, Query{
		"user_id":   userID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetRoleUser(roleID int, page int, pageSize int) (int, []int) {
	var response = struct {
		Response
		Data struct {
			Total int   `json:"total"`
			List  []int `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ROLE_USER_LIST, Query{
		"role_id":   roleID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetUserResource(userID int, resourceType int, page int, pageSize int) (int, []Resource) {
	var response = struct {
		Response
		Data struct {
			Total int        `json:"total"`
			List  []Resource `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_USER_RESOURCE_LIST, Query{
		"user_id":       userID,
		"resource_type": resourceType,
		"page":          page,
		"page_size":     pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetUserAction(userID int, page int, pageSize int) (int, []Action) {
	var response = struct {
		Response
		Data struct {
			Total int      `json:"total"`
			List  []Action `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_USER_ACTION_LIST, Query{
		"user_id":   userID,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}
