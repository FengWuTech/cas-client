package cas_client

const (
	RESOURCE_TYPE_MENU = 1
	RESOURCE_TYPE_DATA = 2

	URL_ACTION_ADD            = "/cas/action/add"
	URL_ACTION_UPDATE         = "/cas/action/update"
	URL_ACTION_UPDATE_BY_CODE = "/cas/action/update_by_code"
	URL_ACTION_DELETE         = "/cas/action/delete"
	URL_ACTION_DELETE_BY_CODE = "/cas/action/delete_by_code"
	URL_ACTION_GET            = "/cas/action/get"
	URL_ACTION_GET_BY_CODE    = "/cas/action/get_by_code"
	URL_ACTION_LIST           = "/cas/action/list"

	URL_RESOURCE_ADD            = "/cas/resource/add"
	URL_RESOURCE_UPDATE         = "/cas/resource/update"
	URL_RESOURCE_UPDATE_BY_CODE = "/cas/resource/update_by_code"
	URL_RESOURCE_DELETE         = "/cas/resource/delete"
	URL_RESOURCE_DELETE_BY_CODE = "/cas/resource/delete_by_code"
	URL_RESOURCE_GET            = "/cas/resource/get"
	URL_RESOURCE_GET_BY_CODE    = "/cas/resource/get_by_code"
	URL_RESOURCE_LIST           = "/cas/resource/list"
	URL_RESOURCE_TREE           = "/cas/resource/tree"

	URL_RESOURCE_ACTION_ADD    = "/cas/resource/action/add"
	URL_RESOURCE_ACTION_DELETE = "/cas/resource/action/delete"
	URL_RESOURCE_ACTION_LIST   = "/cas/resource/action/list"
	URL_ACTION_RESOURCE_LIST   = "/cas/action/resource/list"

	URL_ROLE_ADD            = "/cas/role/add"
	URL_ROLE_UPDATE         = "/cas/role/update"
	URL_ROLE_UPDATE_BY_CODE = "/cas/role/update_by_code"
	URL_ROLE_DELETE         = "/cas/role/delete"
	URL_ROLE_DELETE_BY_CODE = "/cas/role/delete_by_code"
	URL_ROLE_GET            = "/cas/role/get"
	URL_ROLE_GET_BY_CODE    = "/cas/role/get_by_code"
	URL_ROLE_LIST           = "/cas/role/list"

	URL_ROLE_RESOURCE_ADD    = "/cas/role/resource/add"
	URL_ROLE_RESOURCE_DELETE = "/cas/role/resource/delete"
	URL_ROLE_RESOURCE_LIST   = "/cas/role/resource/list"
	URL_RESOURCE_ROLE_LIST   = "/cas/resource/role/list"

	URL_ROLE_ACTION_ADD    = "/cas/role/action/add"
	URL_ROLE_ACTION_DELETE = "/cas/role/action/delete"
	URL_ROLE_ACTION_LIST   = "/cas/role/action/list"
	URL_ACTION_ROLE_LIST   = "/cas/action/role/list"

	URL_USER_ROLE_ADD      = "/cas/user/role/add"
	URL_USER_ROLE_DELETE   = "/cas/user/role/delete"
	URL_USER_ROLE_LIST     = "/cas/user/role/list"
	URL_ROLE_USER_LIST     = "/cas/role/user/list"
	URL_USER_RESOURCE_LIST = "/cas/user/resource/list"
	URL_USER_ACTION_LIST   = "/cas/user/action/list"
)
