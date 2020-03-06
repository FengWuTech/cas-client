package cas_client

import "cas-client/models"

func (cas *Cas) AddResource(typ int, code string, name string, data string) int {
	var response = struct {
		models.Response
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}{}
	cas.HttpPost(URL_RESOURCE_ADD, nil, models.RequestBody{
		"type": typ,
		"code": code,
		"name": name,
		"data": data,
	}, &response)
	return response.Data.ID
}

func (cas *Cas) UpdateResource(id int, resource models.Resource) bool {
	var response models.Response
	cas.HttpPost(URL_RESOURCE_UPDATE, nil, models.RequestBody{
		"id":   id,
		"code": resource.Code,
		"name": resource.Name,
		"data": resource.Data,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteResource(id int) bool {
	var response models.Response
	cas.HttpPost(URL_RESOURCE_DELETE, nil, models.RequestBody{
		"id": id,
	}, &response)
	return response.Code == 200
}

func (cas *Cas) GetResource(id int) *models.Resource {
	var response = struct {
		models.Response
		Resource *models.Resource `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_GET, models.Query{"id": id}, &response)
	return response.Resource
}

// typ 资源类型，<=100的资源类型值含义由系统定义，>100的资源类型可以由业务方自己定义
// pid 父级资源id
func (cas *Cas) GetResourceList(typ int, pid int, name string, page int, pageSize int) (int, []models.Resource) {
	var response = struct {
		models.Response
		Data struct {
			Total int               `json:"total"`
			List  []models.Resource `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_RESOURCE_LIST, models.Query{
		"type":      typ,
		"pid":       pid,
		"name":      name,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}
