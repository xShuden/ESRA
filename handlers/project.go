package handlers

import (
	"esra/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowAddProject renders the form to add a new project
func ShowAddProject(c *gin.Context) {
	c.HTML(http.StatusOK, "add_project.html", gin.H{})
}

// AddProject handles the submission of a new project
func AddProject(c *gin.Context) {
	projectName := c.PostForm("project_name")
	projectPlatform := c.PostForm("project_platform")

	if projectName == "" || projectPlatform == "" {
		c.HTML(http.StatusBadRequest, "add_project.html", gin.H{
			"error": "Both Project Name and Platform are required!",
		})
		return
	}

	if err := models.CreateProject(projectName, projectPlatform); err != nil {
		c.HTML(http.StatusInternalServerError, "add_project.html", gin.H{
			"error": "Failed to add project to database!",
		})
		return
	}

	c.HTML(http.StatusOK, "add_project.html", gin.H{
		"message": "Project added successfully!",
	})
}

// ShowProjects displays a list of all projects
func ShowProjects(c *gin.Context) {
	projects, err := models.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	formattedProjects := make([]map[string]string, len(projects))
	for i, project := range projects {
		formattedProjects[i] = map[string]string{
			"id":       strconv.Itoa(project.ID),
			"name":     project.Name,
			"platform": project.Platform,
		}
	}

	c.HTML(http.StatusOK, "projects.html", gin.H{
		"title":    "Project List",
		"projects": formattedProjects,
	})
}

// ShowProjectDetails displays detailed information about a specific project
func ShowProjectDetails(c *gin.Context) {
	projectID := c.Param("id")
	project, err := models.GetProjectByID(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	c.HTML(http.StatusOK, "project_detail.html", gin.H{
		"title":   "Project Details",
		"project": project,
	})
}
