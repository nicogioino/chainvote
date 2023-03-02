package enums

type ElectionStatus int

const (
	PENDING = 0
	ACTIVE = 1
	ENDED = 2
	CANCELLED = 3
)

func (es ElectionStatus) String () string {
	switch es {
	case PENDING:
		return "PENDING"
	case ACTIVE:
		return "ACTIVE"
	case ENDED:
		return "ENDED"
	case CANCELLED:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}

func FromString (s string) ElectionStatus {
	switch s {
	case "PENDING":
		return PENDING
	case "ACTIVE":
		return ACTIVE
	case "ENDED":
		return ENDED
	case "CANCELLED":
		return CANCELLED
	default:
		return PENDING
	}
}