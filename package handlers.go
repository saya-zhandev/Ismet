package handlers

import (
	"encoding/json"
	"net/http"
)

// Project represents a simple project structure
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetProject handles GET requests to retrieve projects
func GetProject(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{ID: "1", Name: "Project One"},
		{ID: "2", Name: "Project Two"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// CreateProject handles POST requests to create a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here you would normally add the project to your database
	project.ID = "3" // Example ID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// DeleteProject handles DELETE requests to delete a project
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Here you would normally delete the project from your database
	w.WriteHeader(http.StatusNoContent)
}

// Project represents a simple project structure
type Project struct {
	ID          string `json:"id"`
	Github      string `json:"github"`
	DomainName  string `json:"domainName"`
	Link        string `json:"link"`
	LastUpdated string `json:"lastUpdated"`
	Favicon     string `json:"favicon"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{
			ID:          "1",
			Github:      "https://github.com/user/repo1",
			DomainName:  "example1.com",
			Link:        "https://example1.com",
			LastUpdated: "2025-05-19T14:00:00Z",
			Favicon:     "https://example1.com/favicon.ico",
		},
		{
			ID:          "2",
			Github:      "https://github.com/user/repo2",
			DomainName:  "example2.com",
			Link:        "https://example2.com",
			LastUpdated: "2025-05-20T14:00:00Z",
			Favicon:     "https://example2.com/favicon.ico",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing project ID", http.StatusBadRequest)
		return
	}

	// Here you would normally delete the project from your database
	// For example:
	// err := deleteProjectFromDatabase(id)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }

	w.WriteHeader(http.StatusNoContent)
}
