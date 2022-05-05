package business

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePwd(t *testing.T) {
	type generatePwdTestCase struct {
		description  string
		hasError     bool
		minLength    int
		specialChars int
		numbers      int
	}
	for _, tc := range []generatePwdTestCase{
		{
			"error due to minLength < specialChars + numbers",
			true,
			5,
			5,
			1,
		},
		{
			"check correctness of specified params",
			false,
			5,
			3,
			1,
		},
		{
			"check correctness of specified params",
			false,
			8,
			2,
			4,
		},
		{
			"check correctness of specified params",
			false,
			15,
			2,
			2,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			pwd, err := generatePwd(tc.minLength, tc.specialChars, tc.numbers)
			if tc.hasError {
				assert.NotNil(t, err)
				return
			}
			pwdStr := string(pwd)

			digitsCount := countDigits(pwdStr)
			assert.Equal(t, tc.numbers, digitsCount, "digits count")
			specialCharsCount := countSpecialChars(pwdStr)
			assert.Equal(t, tc.specialChars, specialCharsCount, "special chars count")
			assert.Equal(t, tc.minLength, len(pwdStr), "characters count")
		})
	}
}

func TestDigitsCount(t *testing.T) {
	type digitsCountTestCase struct {
		in     string
		digits int
	}

	for _, tc := range []digitsCountTestCase{
		{
			"67890",
			5,
		},
		{
			"abc5def6ghi7",
			3,
		},
		{
			"asdf1234567890",
			10,
		},
		{
			"098765ttttt890",
			9,
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			digitsCount := countDigits(tc.in)
			assert.Equal(t, tc.digits, digitsCount)
		})

	}
}

func TestSpecialChatsCount(t *testing.T) {
	type specialCharsCountTestCase struct {
		in           string
		specialChars int
	}

	for _, tc := range []specialCharsCountTestCase{
		{
			")H*AU*f&k",
			4,
		},
		{
			"rVEPVO+#Q",
			2,
		},
		{
			"asdf1^^^^^234567890",
			5,
		},
		{
			"sbk-+eA*-",
			4,
		},
	} {
		t.Run(tc.in, func(t *testing.T) {
			specialCharsCount := countSpecialChars(tc.in)
			assert.Equal(t, tc.specialChars, specialCharsCount)
		})

	}
}
