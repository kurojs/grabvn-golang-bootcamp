package returncode

const (
	Unknown = iota
	OK
	NotFound
	DuplicateFeedback
)

var messages map[int]string

// Message ...
func Message(code int) string {
	if messages != nil {
		return messages[code]
	}

	messages = map[int]string{
		0: "Unknow error",
		1: "OK",
		2: "Not found",
		3: "Duplicate feedback",
	}

	return messages[code]
}
