package application

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

type Application struct {
	ID                  int    `sql:"id" json:"id"`
	UserId              int    `sql:"user_id" json:"userId"`
	UserUid             string `sql:"user_uid" json:"userUid"`
	JobTitle            string `sql:"job_title" json:"jobTitle"`
	JobDescription      string `sql:"job_description" json:"jobDescription"`
	JobSourceUrl        string `sql:"job_source_url" json:"jobSourceUrl"`
	CompanyId           int    `sql:"company_id" json:"companyId"`
	ResumeId            int    `sql:"resume_id" json:"resumeId"`
	CoverLetterId       int    `sql:"cover_letter_id" json:"coverLetterId"`
	NotesId             int    `sql:"notes_id" json:"notesId"`
	ApplicationDeadline string `sql:"application_deadline" json:"applicationDeadline"`
	ApplicationDate     string `sql:"application_date" json:"applicationDate"`
	InterviewStartDate  string `sql:"interview_start_date" json:"interviewStartDate"`
	InterviewEndDate    string `sql:"interview_end_date" json:"interviewEndDate"`
	RejectionDate       string `sql:"rejection_date" json:"rejectionDate"`
	OfferDate           string `sql:"offer_date" json:"offerDate"`
	AcceptedDate        string `sql:"accepted_date" json:"acceptedDate"`
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindApplicationById(id int) (*Application, error) {
	application := &Application{}
	query := "SELECT * FROM applications WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&application.ID, &application.UserId, &application.JobTitle)
	if err != nil {
		fmt.Println("error with the query", err)
		return nil, err
	}
	return application, nil
}
