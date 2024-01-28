package errors

type ConfigParsingError struct {
	Message string
}

func (e ConfigParsingError) Error() string {
	return e.Message
}

func CreateConfigParserError(message string) error {
	return ConfigParsingError{Message: message}
}
