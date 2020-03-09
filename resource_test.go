package cas_client

import (
	"testing"
)

func TestCas_AddResource(t *testing.T) {
	id := TestNew().AddResource(1, "resource-101", "resource-name", "333333333")
	LogV(id)
}

func TestCas_UpdateResource(t *testing.T) {
	var data = "==========="
	res := TestNew().UpdateResource(78, Resource{
		Data: &data,
	})
	LogV(res)
}

func TestCas_UpdateResourceByCode(t *testing.T) {
	var code = "resource-101"
	var name = "---1--"
	res := TestNew().UpdateResourceByCode(code, Resource{
		Name: &name,
	})
	LogV(res)
}

func TestCas_DeleteResource(t *testing.T) {
	res := TestNew().DeleteResource(78)
	LogV(res)
}

func TestCas_DeleteResourceByCode(t *testing.T) {
	res := TestNew().DeleteResourceByCode("--1--")
	LogV(res)
}

func TestCas_GetResource(t *testing.T) {
	res := TestNew().GetResource(79)
	LogV(res)
}

func TestCas_GetResourceByCode(t *testing.T) {
	res := TestNew().GetResourceByCode("---1--")
	LogV(res)
}

func TestCas_GetResourceList(t *testing.T) {
	total, list := TestNew().GetResourceList(1, 1, "", 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_AddResourceAction(t *testing.T) {
	res := TestNew().AddResourceAction([]int{80, 81}, []int{8, 9})
	LogV(res)
}

func TestCas_DeleteResourceAction(t *testing.T) {
	res := TestNew().DeleteResourceAction([]int{80, 81}, []int{8, 9})
	LogV(res)
}

func TestCas_GetResourceAction(t *testing.T) {
	total, list := TestNew().GetResourceAction(80, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetActionResource(t *testing.T) {
	total, list := TestNew().GetActionResource(8, 1, 1)
	LogV(total)
	LogV(list)
}

func TestCas_GetResourceTree(t *testing.T) {
	tree := TestNew().GetResourceTree(1, 0)
	LogV(tree)
}
