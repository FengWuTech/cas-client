package cas_client

func (cas *Cas) AddAction(code string, name string, description string) int {
	var response = struct {
		Response
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}{}
	var requestBody = RequestBody{
		"code":        code,
		"name":        name,
		"description": description,
	}
	cas.HttpPost(URL_ACTION_ADD, nil, requestBody, &response)
	return response.Data.ID
}

func (cas *Cas) UpdateAction(id int, action Action) bool {
	var response Response
	action.ID = &id
	cas.HttpPost(URL_ACTION_UPDATE, nil, action, &response)
	return response.Code == 200
}

func (cas *Cas) UpdateActionByCode(code string, action Action) bool {
	var response Response
	cas.HttpPost(URL_ACTION_UPDATE_BY_CODE, nil, RequestBody{
		"code": code,
		"name": action.Name,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteAction(id int) bool {
	var response Response
	cas.HttpPost(URL_ACTION_DELETE, nil, map[string]int{"id": id}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteActionByCode(code string) bool {
	var response Response
	cas.HttpPost(URL_ACTION_DELETE_BY_CODE, nil, RequestBody{
		"code": code,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetAction(id int) *Action {
	var response = struct {
		Response
		Action *Action `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_GET, Query{
		"id": id,
	}, &response)
	return response.Action
}

func (cas *Cas) GetActionByCode(code string) *Action {
	var response = struct {
		Response
		Action *Action `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_GET_BY_CODE, Query{
		"code": code,
	}, &response)
	return response.Action
}

func (cas *Cas) GetActionList(name string, page int, pageSize int) (int, []Action) {
	var response = struct {
		Response
		Data struct {
			Total int      `json:"total"`
			List  []Action `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_LIST, Query{
		"name":      name,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}
