package cas_client

import "testing"

func TestCas_AddRole(t *testing.T) {
	res := TestNew().AddRole("testing-role-2", "testing-role", "---1--")
	LogV(res)
}

func TestCas_UpdateRole(t *testing.T) {
	var code = "role-update-2"
	res := TestNew().UpdateRole(1, Role{
		Code: &code,
	})
	LogV(res)
}

func TestCas_UpdateRoleByCode(t *testing.T) {
	var code = "role-update-2"
	var name = "rolename"
	res := TestNew().UpdateRoleByCode(code, Role{
		Name: &name,
	})
	LogV(res)
}

func TestCas_DeleteRole(t *testing.T) {
	res := TestNew().DeleteRole(1)
	LogV(res)
}

func TestCas_DeleteRoleByCode(t *testing.T) {
	res := TestNew().DeleteRoleByCode("--1--")
	LogV(res)
}

func TestCas_GetRole(t *testing.T) {
	res := TestNew().GetRole(3)
	LogV(res)
}

func TestCas_GetRoleByCode(t *testing.T) {
	res := TestNew().GetRoleByCode("--1--")
	LogV(res)
}

func TestCas_GetRoleList(t *testing.T) {
	total, list := TestNew().GetRoleList("", 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_AddRoleResource(t *testing.T) {
	res := TestNew().AddRoleResource([]int{1, 2, 3}, []int{4, 5, 6})
	LogV(res)
}

func TestCas_DeleteRoleResource(t *testing.T) {
	res := TestNew().DeleteRoleResource([]int{1, 2, 3}, []int{4, 5, 6})
	LogV(res)
}

func TestCas_GetRoleResource(t *testing.T) {
	total, list := TestNew().GetRoleResource(3, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetResourceRole(t *testing.T) {
	total, list := TestNew().GetResourceRole(1, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_AddRoleAction(t *testing.T) {
	res := TestNew().AddRoleAction([]int{3}, []int{1, 2, 3})
	LogV(res)
}

func TestCas_DeleteRoleAction(t *testing.T) {
	res := TestNew().DeleteRoleAction([]int{3}, []int{1, 2, 3})
	LogV(res)
}

func TestCas_GetRoleAction(t *testing.T) {
	total, list := TestNew().GetRoleAction(3, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetActionRole(t *testing.T) {
	total, list := TestNew().GetActionRole(1, 1, 1)
	LogV(total)
	LogV(list)
}
