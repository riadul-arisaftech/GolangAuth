package util

const (
	ACTIVE   = "active"
	INACTIVE = "inactive"
	DISABLED = "disabled"
)

func IsSupportedStatus(status string) bool {
	switch status {
	case ACTIVE, INACTIVE, DISABLED:
		return true
	}
	return false
}
