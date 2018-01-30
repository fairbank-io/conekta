package conekta

// When creating any call to our service, using the API, you will be notified
// in the event of any errors and provided with all the respective error information.
// https://developers.conekta.com/api?language=bash#errors
type APIError struct {
	// Contains the error type and error code
	Type string `json:"type"`

	// The id of the http log of the request which was performed
	LogID string `json:"log_id"`

	// Detailed list of the errors
	Details []ErrorDetails `json:"details"`
}

// // Detailed information of the errors
type ErrorDetails struct {
	// Human-readable message which provides more details about the error. This message
	// is meant to be shown to purchasers and it is available in English and Spanish.
	// For card charges, the message can be displayed to the user.
	Message string `json:"message"`

	// Human-readable message which provides more details about the error. This message
	// is meant to be used for internal debugging and it is only available in English.
	// For card charges, the message can be displayed to the user.
	DebugMessage string `json:"debug_message"`

	// The parameter to which the error is related. You can use this to highlight
	// erroneous form fields.
	Params string `json:"params"`

	// A short, specific error code to elaborate on processing_error
	Code string `json:"code"`
}

func (e *APIError) Error() string {
	return e.Type
}
