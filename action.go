package cas_client

import "cas-client/models"

func (cas *Cas) AddAction(code string, name string) int {
	var response = struct {
		models.Response
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}{}
	var requestBody = models.RequestBody{
		"code": code,
		"name": name,
	}
	cas.HttpPost(URL_ACTION_ADD, nil, requestBody, &response)
	return response.Data.ID
}

func (cas *Cas) UpdateAction(id int, action models.Action) bool {
	var response models.Response
	action.ID = &id
	cas.HttpPost(URL_ACTION_UPDATE, nil, action, &response)
	return response.Code == 200
}

func (cas *Cas) DeleteAction(id int) bool {
	var response models.Response
	cas.HttpPost(URL_ACTION_DELETE, nil, map[string]int{"id": id}, &response)
	return response.Code == 200
}

func (cas *Cas) GetAction(id int) *models.Action {
	var response = struct {
		models.Response
		Action *models.Action `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_GET, models.Query{
		"id": id,
	}, &response)
	return response.Action
}

func (cas *Cas) GetActionList(name string, page int, pageSize int) (int, []models.Action) {
	var response = struct {
		models.Response
		Data struct {
			Total int             `json:"total"`
			List  []models.Action `json:"list"`
		} `json:"data"`
	}{}
	cas.HttpGet(URL_ACTION_LIST, models.Query{
		"name":      name,
		"page":      page,
		"page_size": pageSize,
	}, &response)
	return response.Data.Total, response.Data.List
}
