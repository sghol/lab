package main

import (
	"fmt"
	"strconv"
)

// Represent a user information in the execel file
type UserRow struct {
	NationalID string
	FullName   string
	FatherName string
	Phone      string
	BirthDate  string
	GradeCode  string
	RoleCode   string
}

// ImportService handles the legacy data import.
type ImportService struct {
	tenantID int
}

func NewImportService(tenantID int) *ImportService {
	return &ImportService{tenantID: tenantID}
}

func (s *ImportService) parseGradeCode(gradeCode string) (grade, major, class int, err error) {
	// 1211  12: grade, 1: Math major 2: class #2
	// 7001  70 : grade 7  0: No major  1: class #2
	if len(gradeCode) < 4 {
		return 0, 0, 0, fmt.Errorf("invalid grade level gradeCode: %s", gradeCode)
	}

	major, _ = strconv.Atoi(gradeCode[2:3])
	class, _ = strconv.Atoi(gradeCode[3:])
	// No major means elementry or middle shcool
	if gradeCode[2:3] == "0" {
		grade, _ = strconv.Atoi(gradeCode[0:1])
	} else {
		grade, _ = strconv.Atoi(gradeCode[0:2])
	}

	return
}

func main() {
	importSvc := NewImportService(1)
	grade, major, class, err := importSvc.parseGradeCode("5004")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(grade, major, class)
}
