package projects

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.redberry.api/domain/entities"
	"com.redberry.api/domain/services/projects"
)

// Project handles project-related endpoints.
type Project struct {
	service *projects.Service
}

// New creates a new instance of the Project controller.
func New() *Project {
	return &Project{
		service: projects.New(),
	}
}

// Create godoc
// @Summary      Create Project
// @Description  Creates a new project.
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        body body      entities.Project true "Project data"
// @Success      201  {object}  entities.Project "Created project"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/projects [post]
func (controller Project) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody entities.Project

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := controller.service.Create(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// GetByID godoc
// @Summary      Get Project by ID
// @Description  Retrieves a project by its ID.
// @Tags         projects
// @Produce      json
// @Param        project_id path      int true "Project ID"
// @Success      200  {object}  entities.Project "Project data"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/projects/{project_id} [get]
func (controller Project) GetByID(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(r.PathValue("project_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := controller.service.GetByID(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Update godoc
// @Summary      Update Project
// @Description  Updates an existing project.
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        project_id path      int true "Project ID"
// @Param        body body      entities.Project true "Updated project data"
// @Success      200  {object}  entities.Project "Updated project"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/projects/{project_id} [put]
func (controller Project) Update(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(r.PathValue("project_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody entities.Project

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := controller.service.Update(projectID, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Delete godoc
// @Summary      Delete Project
// @Description  Deletes a project by its ID.
// @Tags         projects
// @Produce      json
// @Param        project_id path      int true "Project ID"
// @Success      200  {object}  entities.Project "Deleted project"
// @failure      400  {string}  string "Bad Request"
// @Failure      500  {string}  string "Internal Server Error"
// @Router       /api/v1/projects/{project_id} [delete]
func (controller Project) Delete(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(r.PathValue("project_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	project, err := controller.service.Delete(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
