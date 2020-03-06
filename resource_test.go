package cas_client

import (
	"cas-client/models"
	"testing"
)

func TestCas_AddResource(t *testing.T) {
	id := TestNew().AddResource(1, "resource-101", "resource-name", "333333333")
	LogV(id)
}

func TestCas_UpdateResource(t *testing.T) {
	var data = "==========="
	res := TestNew().UpdateResource(78, models.Resource{
		Data: &data,
	})
	LogV(res)
}

func TestCas_DeleteResource(t *testing.T) {
	res := TestNew().DeleteResource(78)
	LogV(res)
}

func TestCas_GetResource(t *testing.T) {
	res := TestNew().GetResource(79)
	LogV(res)
}

func TestCas_GetResourceList(t *testing.T) {
	total, list := TestNew().GetResourceList(1, 1, "", 1, 1)
	LogV(total)
	LogV(list)
}
