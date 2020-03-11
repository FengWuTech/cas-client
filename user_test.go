package cas_client

import "testing"

func TestCas_AddUser(t *testing.T) {
	res := TestNew().AddUser(71)
	LogV(res)
}

func TestCas_DeleteUser(t *testing.T) {
	res := TestNew().DeleteUser(71)
	LogV(res)
}

func TestCas_GetUserList(t *testing.T) {
	total, list := TestNew().GetUserList(1, "", 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_AddUserRole(t *testing.T) {
	res := TestNew().AddUserRole([]int{17}, []int{4})
	LogV(res)
}

func TestCas_DeleteUserRole(t *testing.T) {
	res := TestNew().DeleteUserRole([]int{17}, []int{3})
	LogV(res)
}

func TestCas_GetUserRole(t *testing.T) {
	total, list := TestNew().GetUserRole(17, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetRoleUser(t *testing.T) {
	total, list := TestNew().GetRoleUser(3, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetUserResource(t *testing.T) {
	total, list := TestNew().GetUserResource(17, 1, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetUserAction(t *testing.T) {
	total, list := TestNew().GetUserAction(17, 1, 1)
	LogV(total)
	LogV(list)
}
