// +build db

package storage

import (
	"testing"
)

func TestCategoryRepository_GetAll(t *testing.T) {
	cr := NewCategoryRepository(getTestDB())
	res, err := cr.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(res) == 0 {
		t.Errorf("nil len")
	}
}
