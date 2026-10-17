
// Import data from excel file

package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	// authdom "github.com/sghol/lms/internal/auth/domain"
	// authrepo "github.com/sghol/lms/internal/auth/repository"
	// "github.com/sghol/lms/internal/school/repository"
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
	studentSvc *StudentService
	tenantID   int
}

func NewImportService(tenantID int, studentSvc *StudentService) *ImportService {
	return &ImportService{studentSvc: studentSvc, tenantID: tenantID}
}

// 1211  12: grade, 1: Math major 1: class #1
// 7004  70 : grade 7  0: No major  4: class #4
func (s *ImportService) parseGradeCode(gradeCode string) (grade, major, class int, err error) {
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

// getRoleName get role name for role code
// func getRoleName(roleCode string) (string, error) {
// 	roleMap := map[string]string{
// 		"d": "student",
// 		"m": "teacher",
// 		"p": "staff",
// 	}
//
// 	if role, exists := roleMap[roleCode]; exists {
// 		return role, nil
// 	}
// 	return "", fmt.Errorf("unknown role code: %s", roleCode)
// }

// func getGradeName(grade int) (string, error) {
// 	gradeMap := map[int]string{
// 		1:  "اول",
// 		2:  "دوم",
// 		3:  "سوم",
// 		4:  "چهارم",
// 		5:  "پنجم",
// 		6:  "ششم",
// 		7:  "هفتم",
// 		8:  "هشتم",
// 		9:  "نهم",
// 		10: "دهم",
// 		11: "یازدهم",
// 		12: "دوازدهم",
// 	}
// 	if gradeName, exists := gradeMap[grade]; exists {
// 		return gradeName, nil
// 	}
// 	return "", fmt.Errorf("unknown grade: %d", grade)
// }
//
// func getMajorName(majorCode int) (string, error) {
// 	majorMap := map[int]string{
// 		0: "",
// 		1: "ریاضی",
// 		2: "تجربی",
// 		3: "انسانی",
// 	}
//
// 	majorName, exists := majorMap[majorCode]
// 	if !exists {
// 		return "", nil
// 	}
// 	return majorName, fmt.Errorf("Unknow manjor %d", majorCode)
// }
//
// func (s *ImportService) processRow(row ImportRow) error {
// 1. Create user and assign role
// 2. Crrate classroom
// 3. Assign every student that belong to the class.
// 4. If there is no class for student just create class.
//
// 	// Start transaction
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()
//
// 	// Check if user exists
// 	user, err := s.userRepo.FindByNationalID(row.NationalID)
// 	if err != nil {
// 		return err
// 	}
//
// 	// 1. Create user if doesn't exist
// 	if user == nil {
//
// 		user = &authdom.User{
// 			NationalID: row.NationalID,
// 			Phone:      row.Phone,
// 			BirthDate:  row.BirthDate,
// 		}
// 		firstName, lastName := user.GetFirstLastName(row.FullName)
// 		user.FirstName = firstName
// 		user.LastName = lastName
// 		password, _ := s.passwordSvc.HashPassword(user.NationalID)
// 		userID, err := s.userRepo.Create(user, password)
// 		if err != nil {
// 			return err
// 		}
// 		user.ID = userID
// 	}
// 	// 2. Create role for the user
// 	// Get role ID from code
// 	roleName, err := getRoleName(row.RoleCode)
// 	if err != nil {
// 		return err
// 	}
//
// 	role, err := s.roleRepo.FindByName(roleName)
// 	if err != nil {
// 		return err
// 	}
// 	s.roleRepo.AssignUserRole(user.ID, s.tenantID, role.ID)
//
// 	// 3. parrse the grade_code (class_code)
// 	grade, major, number, err := parseGradeCode(row.GradeCode)
// 	if err != nil {
// 		return err
// 	}
// 	gradeName, err := getGradeName(grade)
// 	if err != nil {
// 		return err
// 	}
// 	classroomName := fmt.Sprintf("%s %d %d", gradeName, major, number)
// 	fmt.Println(classroomName)
//
// 	// Create classroom
//
// 	// enroll student to the classroom
//
// 	return tx.Commit()
// }
//
// func (s *ImportService) OpenAndParseExcelFile(filePath string) {
// 	// open excel file
// 	filePath = "/home/saeed/dow/efaf2.xlsx"
// 	f, err := excelize.OpenFile(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
//
// 	// Parse excel file
// 	// first sheet
// 	sheets := f.GetSheetList()
// 	firstSheet := sheets[0]
//
// 	rows, err := f.GetRows(firstSheet)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	firstRow := rows[1]
//
// 	row := ImportRow{
// 		NationalID: firstRow[0],
// 		FullName:   firstRow[1],
// 		FatherName: firstRow[2],
// 		Phone:      firstRow[3],
// 		BirthDate:  firstRow[4],
// 		GradeCode:  firstRow[5],
// 		RoleCode:   firstRow[6],
// 	}
// 	fmt.Println(row)
// }

func (s *ImportService) ImportUserFromExcel(filePath string) {
	// 1. Create user and assign role
	// 2. Crrate classroom
	// 3. Assign every student that belong to the class.
	// 4. If there is no class for student just create class.

	// Read and create struct from excel row
	// Open excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parse excel file
	// first sheet
	sheets := f.GetSheetList()
	firstSheet := sheets[0]

	rows, err := f.GetRows(firstSheet)
	if err != nil {
		log.Fatal(err)
	}

	firstRow := rows[1]

	row := UserRow{
		NationalID: firstRow[0],
		FullName:   firstRow[1],
		FatherName: firstRow[2],
		Phone:      firstRow[3],
		BirthDate:  firstRow[4],
		GradeCode:  firstRow[5],
		RoleCode:   firstRow[6],
	}

	fmt.Println(row)
}




package school

import (
	"database/sql"
	"net/http"

	authmw "github.com/sghol/lms/internal/auth/middleware"
	authrepo "github.com/sghol/lms/internal/auth/repository"
	authsvc "github.com/sghol/lms/internal/auth/service"
	"github.com/sghol/lms/internal/school/handler"
	"github.com/sghol/lms/internal/school/repository"
	"github.com/sghol/lms/internal/school/service"
)

type Module struct {
	StudentRepo    *repository.StudentRepository
	StudentService *service.StudentService
	StudentHandler *handler.StudentHandler
}

func NewModule(
	db *sql.DB,
	userRepo *authrepo.UserRepository,
	roleRepo *authrepo.RoleRepository,
	passwordSvc interface {
		HashPassword(password string) (string, error)
	},
) *Module {
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo, userRepo, roleRepo, passwordSvc)
	studentHandler := handler.NewStudentHandler(studentService)
	return &Module{
		StudentRepo:    studentRepo,
		StudentService: studentService,
		StudentHandler: studentHandler,
	}
}

// RegisterAPIRoutes adds all school routes under /school prefix.
// Uses Go 1.22+ method-based routing.
func (m *Module) RegisterAPIRoutes(router *http.ServeMux, permSvc *authsvc.PermissionService) {
	// Helper for permission + tenant middleware
	require := authmw.RequireTenantAndPermission

	// Collection routes – method-specific patterns
	router.Handle("POST /students",
		require(permSvc, "student:create")(
			http.HandlerFunc(m.StudentHandler.CreateStudent),
		),
	)
	router.Handle("GET /students",
		require(permSvc, "student:read")(
			http.HandlerFunc(m.StudentHandler.GetAllStudents),
		),
	)

	// Item routes with path variable – method-specific patterns
	router.Handle("GET /students/{id}",
		require(permSvc, "student:read")(
			http.HandlerFunc(m.StudentHandler.GetStudent),
		),
	)

}

func (m *Module) ImportExcel(w http.ResponseWriter, r *http.Request) {
		handler.ImportExcel(w, r)
	
}
