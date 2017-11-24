package entity

import "testing"

func TestUserRole_Value(t *testing.T) {
	r := RoleUser
	if str, err := r.Value(); str.(string) != "user" || err != nil {
		t.Errorf("wrong value: %v, %v", str, err)
	}
}
