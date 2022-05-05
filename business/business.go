package business

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	specialCharsSet = []string{
		"!", "@", "#", "$", "%", "^", "&", "*",
		"(", ")", "-", "_", "=", "+",
	}
	alphabetSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Params struct {
	MinLength         int `json:"min_length"`
	SpecialCharsCount int `json:"special_chars_count"`
	NumbersCount      int `json:"numbers_count"`
	PwdsToCreate      int `json:"pwds_to_create"`
}

type Password string

func ParseParams(rc io.ReadCloser) (Params, error) {
	var p Params

	err := json.NewDecoder(rc).Decode(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func generatePwds(minLength, specialChars, numbers, pwdsToCreate int) ([]Password, error) {
	var pwdsList []Password

	for idx := 0; idx < pwdsToCreate; idx++ {
		newPwd, err := generatePwd(minLength, specialChars, numbers)
		if err != nil {
			return pwdsList, err
		}
		pwdsList = append(pwdsList, newPwd)
	}

	return pwdsList, nil
}

func generatePwd(minLength, specialChars, numbers int) (Password, error) {
	//TODO: Validation: minLength !> specialChars + numbers
	if minLength < (specialChars + numbers) {
		return "", fmt.Errorf("minLength: %d is invalid", minLength)
	}

	var b strings.Builder

	// TODO: Increase efficiency by combining the 2 loops into one
	for idx := 0; idx < specialChars; idx++ {
		sCharIdx := rand.Intn(len(specialCharsSet))
		fmt.Fprintf(&b, "%s", specialCharsSet[sCharIdx])
	}

	for idx := 0; idx < numbers; idx++ {
		fmt.Fprintf(&b, "%d", rand.Intn(10))
	}

	for idx := b.Len(); idx < minLength; idx++ {
		sCharIdx := rand.Intn(len(alphabetSet))
		fmt.Fprintf(&b, "%c", alphabetSet[sCharIdx])
	}

	pwdRune := []rune(b.String())
	rand.Shuffle(len(pwdRune), func(i, j int) {
		pwdRune[i], pwdRune[j] = pwdRune[j], pwdRune[i]
	})

	pwdStr := string(pwdRune)

	return (Password)(pwdStr), nil
}

func GetPwds(minLength, specialChars, numbers, pwdsToCreate int) ([]string, error) {
	pwdsList, err := generatePwds(minLength, specialChars, numbers, pwdsToCreate)
	if err != nil {
		return []string{}, err
	}

	var pwdsStrList []string
	for _, pwd := range pwdsList {
		pwdsStrList = append(pwdsStrList, (string)(pwd))
	}

	return pwdsStrList, nil
}

func countDigits(in string) int {
	re := regexp.MustCompile("[0-9]+")
	digitGroups := re.FindAllString(in, -1)
	digits := strings.Join(digitGroups, "")

	return len(digits)
}

func countSpecialChars(in string) int {
	specialCharsCount := 0
	for _, char := range in {
		if slices.Contains(specialCharsSet, string(char)) {
			specialCharsCount++
		}
	}

	return specialCharsCount
}
