package exception

import "fmt"

type AdminCreating struct {
	AdminID string
}

func (e *AdminCreating) Error() string {
	return fmt.Sprintf("adminID: %s failed to create", e.AdminID)
}
