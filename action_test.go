package cas_client

import (
	"testing"
)

func TestCas_AddAction(t *testing.T) {
	res := TestNew().AddAction("action-12", "action")
	LogV(res)
}

func TestCas_UpdateAction(t *testing.T) {
	var code = "action-123"
	var name = "namename"
	res := TestNew().UpdateAction(2, Action{
		Code: &code,
		Name: &name,
	})
	LogV(res)
}

func TestCas_DeleteAction(t *testing.T) {
	res := TestNew().DeleteAction(2)
	LogV(res)
}

func TestCas_GetAction(t *testing.T) {
	res := TestNew().GetAction(3)
	LogV(res)
}

func TestCas_GetActionList(t *testing.T) {
	total, list := TestNew().GetActionList("", 1, 1)
	LogV(total)
	LogV(list)
}
