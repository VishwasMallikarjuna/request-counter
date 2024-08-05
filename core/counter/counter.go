package counter

// RegisterAccess defines methods for registering and accessing timestamps.
type RegisterAccess interface {
	Access(int64)
	Register() int64
}
