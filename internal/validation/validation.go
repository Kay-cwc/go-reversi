package valiadation

import (
	"strconv"
	"strings"
)

func IsUintString(v string) (uint, string, bool) {
	intVal, err := strconv.ParseUint(v, 10, 64)
	return uint(intVal), "please input a positive integer", err != nil
}

func ValidateUserMovePrompt(v string) ([]uint, string, bool) {
	targets := strings.Split(v, ",")
	output := make([]uint, 2)
	if len(targets) != 2 {
		return output, "Invalid input", true
	}
	var hasErr bool
	var errMsg string
	output[0], errMsg, hasErr = IsUintString(targets[0])
	if hasErr {
		return output, errMsg, hasErr
	}
	output[1], errMsg, hasErr = IsUintString(targets[1])
	if hasErr {
		return output, errMsg, hasErr
	}
	// the value should be between 1-8
	// should be dynamic later for the chessboard dimension
	// maybe return a closure here
	if output[0] == 0 || output[0] > 8 {
		return output, "x must be between 1-8", true
	}
	if output[1] == 0 || output[1] > 8 {
		return output, "y must be between 1-8", true
	}

	return output, "", false
}
