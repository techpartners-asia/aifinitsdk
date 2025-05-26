package aifinitsdk

// isSuccessStatus checks if the status code is in the 2xx range
func isSuccessStatus(status int) bool {
	return status >= 200 && status < 300
}
