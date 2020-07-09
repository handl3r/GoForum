package formaterror

import "strings"

var errorMessages = make(map[string]string)

func FormatError(errString string) map[string]string {
	if strings.Contains(errString, "username") {
		errorMessages["taken_username"] = "UserName Already Taken"
	}
	if strings.Contains(errString, "email") {
		errorMessages["taken_email"] = "Email Already Taken"
	}
	if strings.Contains(errString, "title") {
		errorMessages["taken_title"] = "Title Already Taken"
	}
	if strings.Contains(errString, "hashedPassword") {
		errorMessages["incorrect_password"] = "Incorrect Password"
	}
	if strings.Contains(errString, "record not found") {
		errorMessages["no_record"] = "No Record Found"
	}
	if strings.Contains(errString, "double like") {
		errorMessages["double_like"] = "You can not like this post twice"
	}
	if len(errorMessages) > 0 {
		return errorMessages
	}
	if len(errorMessages) == 0 {
		errorMessages["incorrect_details"] = "incorrect details"
		return errorMessages
	}
	return nil
}
