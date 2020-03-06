package cas_client

import (
	"cas-client/models"
	"fmt"
	"testing"
)

func TestCas_AddAction(t *testing.T) {
	res := TestNew().AddAction("action-12", "action")
	fmt.Printf("%v", res)
}

func TestCas_UpdateAction(t *testing.T) {
	var code = "action-123"
	var name = "namename"
	res := TestNew().UpdateAction(2, models.Action{
		Code: &code,
		Name: &name,
	})
	fmt.Printf("%v", res)
}

func TestCas_DeleteAction(t *testing.T) {
	res := TestNew().DeleteAction(2)
	fmt.Printf("%v", res)
}

func TestCas_GetAction(t *testing.T) {
	res := TestNew().GetAction(3)
	fmt.Printf("%v", res)
}

func TestCas_GetActionList(t *testing.T) {
	total, list := TestNew().GetActionList("====", 1, 1)
	fmt.Printf("%v", total)
	fmt.Printf("%v", list)
}
