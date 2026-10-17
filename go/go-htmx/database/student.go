package database

type Student struct {
	ID    int
	Name  string
	Email string
	Grade string
}

func GetStudents() ([]Student, error) {
	rows, err := DB.Query("SELECT id, name, email, grade FROM students ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.Grade); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, rows.Err()
}
