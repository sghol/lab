package main


import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
)




type RoleName string

const (
	RoleAdmin   RoleName = "admin"
	RoleManager RoleName = "manager"
	RoleTeacher RoleName = "teacher"
	RoleStaff   RoleName = "staff"
	RoleStudent RoleName = "student"
	RoleParent  RoleName = "parent"

)

func GetFirstLastName(fullName string) (firstName, lastName string) {
	fullName = strings.TrimSpace(fullName)
	if fullName == "" {
		return "", ""
	}
	// Fields method split string by space
	parts := strings.Fields(fullName)
	if len(parts) == 1 {
		return parts[0], ""
	}
	firstName = parts[0]
	lastName = strings.Join(parts[1:], " ")
	return firstName, lastName
}



type GradeParts struct {
	Grade       int
	Major       int // 0=none, 1=math, 2=science, 3=literature
	ClassNumber int
}

// gradeNames maps grade number to Persian ordinal name
var gradeNames = map[int]string{
	1:  "اول",
	2:  "دوم",
	3:  "سوم",
	4:  "چهارم",
	5:  "پنجم",
	6:  "ششم",
	7:  "هفتم",
	8:  "هشتم",
	9:  "نهم",
	10: "دهم",
	11: "یازدهم",
	12: "دوازدهم",
}

// majorNames maps major code to Persian name
var majorNames = map[int]string{
	1: "ریاضی",
	2: "تجربی",
	3: "انسانی",
}

// parseGradeCode extracet grade, major and class number from a gradeCode
func ParseGradeCode(gradeCode string) (GradeParts, error) { // Input validation
	if gradeCode == "" {
		return GradeParts{}, fmt.Errorf("gradeCode cannot be empty.")
	}

	// check minimu length
	if len(gradeCode) != 4 {
		return GradeParts{}, fmt.Errorf("invalid grade level gradeCode: %s", gradeCode)
	}

	// Validate that all characters are digits
	for i, char := range gradeCode {
		if !unicode.IsDigit(char) {
			return GradeParts{}, fmt.Errorf("gradeCode contains non-digit character at position %d: %c in %s", i, char, gradeCode)
		}
	}

	// major is 3rd digit and class number is 4th
	major, _ := strconv.Atoi(gradeCode[2:3])
	classNumber, _ := strconv.Atoi(gradeCode[3:])
	// two first digit as grade if student is in high school.
	grade, _ := strconv.Atoi(gradeCode[0:2])

	// if major is 0 it means first digit is grade level
	if gradeCode[2:3] == "0" {
		grade, _ = strconv.Atoi(gradeCode[0:1])

	}

	return GradeParts{
		Grade:       grade,
		Major:       major,
		ClassNumber: classNumber,
	}, nil
}

func BuildClassName(p GradeParts) (string, error) {
	gradeName, ok := gradeNames[p.Grade]
	if !ok {
		return "", fmt.Errorf("no Persian name for grade %d", p.Grade)
	}

	// convert class number to Persian digit
	classNumPersian := toPersianDigit(p.ClassNumber)

	if p.Major == 0 {
		// no major — grades 1-9 typically
		return fmt.Sprintf("%s-%s", gradeName, classNumPersian), nil
	}

	majorName, ok := majorNames[p.Major]
	if !ok {
		return "", fmt.Errorf("unknown major code %d", p.Major)
	}

	return fmt.Sprintf("%s-%s-%s", gradeName, majorName, classNumPersian), nil
}

// toPersianDigit converts an integer (1-9) to its Persian/Arabic-Indic digit string
func toPersianDigit(n int) string {
	persianDigits := []rune("۰۱۲۳۴۵۶۷۸۹")
	if n < 0 || n > 9 {
		return strconv.Itoa(n) // fallback to ASCII
	}
	return string(persianDigits[n])
}

func main() {
	p := GradeParts{Grade: 11, Major:1, ClassNumber:4}
	res, _ := BuildClassName(p)
	fmt.Println(res)
}
