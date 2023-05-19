package validator

type ErrorsMap map[string][]string

type ValidationErrors struct {
	error
	errors ErrorsMap
}

func newValidationErrors(err error) *ValidationErrors {
	return &ValidationErrors{
		error:  err,
		errors: make(ErrorsMap),
	}
}

func (e *ValidationErrors) Add(field, message string) {
	e.errors[field] = append(e.errors[field], message)
}

func (e *ValidationErrors) Error() string {
	return e.error.Error()
}

func (e *ValidationErrors) Errors() ErrorsMap {
	return e.errors
}
