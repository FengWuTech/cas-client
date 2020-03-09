package cas_client

func (cas *Cas) AddResource(typ int, code string, name string, data string) int {
	var response = struct {
		Response
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}{}
	cas.HttpPost(URL_RESOURCE_ADD, nil, RequestBody{
		"type": typ,
		"code": code,
		"name": name,
		"data": data,
	}, &response)
	return response.Data.ID
}

func (cas *Cas) UpdateResource(id int, resource Resource) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_UPDATE, nil, RequestBody{
		"id":   id,
		"code": resource.Code,
		"name": resource.Name,
		"data": resource.Data,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) UpdateResourceByCode(code string, resource Resource) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_UPDATE_BY_CODE, nil, RequestBody{
		"code": code,
		"name": resource.Name,
		"data": resource.Data,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteResource(id int) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_DELETE, nil, RequestBody{
		"id": id,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteResourceByCode(code string) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_DELETE_BY_CODE, nil, RequestBody{
		"code": code,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetResource(id int) *Resource {
	var response = struct {
		Response
		Resource *Resource `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_GET, Query{"id": id}, &response)
	return response.Resource
}

func (cas *Cas) GetResourceByCode(code string) *Resource {
	var response = struct {
		Response
		Resource *Resource `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_GET_BY_CODE, Query{"code": code}, &response)
	return response.Resource
}

// typ 资源类型，<=100的资源类型值含义由系统定义，>100的资源类型可以由业务方自己定义
// pid 父级资源id
func (cas *Cas) GetResourceList(typ int, pid int, name string, page int, pageSize int) (int, []Resource) {
	var response = struct {
		Response
		Data struct {
			Total int        `json:"total"`
			List  []Resource `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_LIST, Query{
		"type":      typ,
		"pid":       pid,
		"name":      name,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}

func (cas *Cas) AddResourceAction(resourceIDList []int, actionIDList []int) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_ACTION_ADD, nil, RequestBody{
		"resource_id_list": resourceIDList,
		"action_id_list":   actionIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteResourceAction(resourceIDList []int, actionIDList []int) bool {
	var response Response
	cas.HttpPost(URL_RESOURCE_ACTION_DELETE, nil, RequestBody{
		"resource_id_list": resourceIDList,
		"action_id_list":   actionIDList,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetResourceAction(resourceID int, page int, pageSize int) (int, []Action) {
	var response = struct {
		Response
		Data struct {
			Total int      `json:"total"`
			List  []Action `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_ACTION_LIST, Query{
		"resource_id": resourceID,
		"page":        page,
		"page_size":   pageSize,
	}, &response)

	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetActionResource(actionID int, page int, pageSize int) (int, []Resource) {
	var response = struct {
		Response
		Data struct {
			Total int        `json:"total"`
			List  []Resource `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_RESOURCE_LIST, Query{
		"action_id": actionID,
		"page":      page,
		"page_size": pageSize,
	}, &response)

	return response.Data.Total, response.Data.List
}

func (cas *Cas) GetResourceTree(typ int, pid int) []ResourceTree {
	var response = struct {
		Response
		Data []ResourceTree `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_TREE, Query{
		"pid":  pid,
		"type": typ,
	}, &response)
	return response.Data
}
