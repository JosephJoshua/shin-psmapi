package utils

// If we put the UserRole type in the models package,
// then use it in the user form, and then use the user form in the user model,
// it will cause UserRole to be an invalid type.

import "database/sql/driver"

type UserRole string

const (
	AdminUserRole           UserRole = "admin"
	CustomerServiceUserRole UserRole = "customer_service"
)

func IsValidUserRole(r string) bool {
	return r == string(AdminUserRole) || r == string(CustomerServiceUserRole)
}

func (r *UserRole) Scan(value interface{}) error {
	*r = UserRole(value.([]byte))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}
