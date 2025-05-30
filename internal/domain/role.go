package domain

type Role string

const (
	RoleAdmin    Role = "ADMIN"
	RoleCustomer Role = "CUSTOMER"
)

var validRoles = []Role{
	RoleAdmin,
	RoleCustomer,
}

func IsValidRole(role Role) bool {
	for _, v := range validRoles {
		if v == role {
			return true
		}
	}
	return false
}