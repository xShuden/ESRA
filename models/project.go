package models

import (
	"esra/database"
	"time"
)

// Project represents a project record
type Project struct {
	ID        int
	Name      string
	Platform  string
	CreatedAt time.Time
}

// CreateProject inserts a new project record into the database
func CreateProject(name, platform string) error {
	_, err := database.DB.Exec("INSERT INTO projects(name, platform) VALUES($1, $2)", name, platform)
	return err
}

// GetAllProjects retrieves all project records from the database
func GetAllProjects() ([]Project, error) {
	rows, err := database.DB.Query("SELECT id, name, platform, created_at FROM projects ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Platform, &project.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// GetProjectByID retrieves a single project record from the database by ID
func GetProjectByID(id string) (Project, error) {
	var project Project
	err := database.DB.QueryRow("SELECT id, name, platform, created_at FROM projects WHERE id=$1", id).Scan(
		&project.ID, &project.Name, &project.Platform, &project.CreatedAt)
	return project, err
}
