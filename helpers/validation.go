package helpers

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`@`)
	return emailRegex.MatchString(email)
}

func IsValidPassword(password string) bool {
	return len(password) >= 8
}

func ValidateUserInput(email string, password string) (bool, string) {

    if !IsValidEmail(email) {
		return false, "이메일 형식이 올바르지 않습니다. 다시 입력해주세요."
	}

	if !IsValidPassword(password) {
		return false, "비밀번호는 최소 8자 이상이어야 합니다."
	}

	return true, ""
}