package controller

import (
	"Cih2001/ActionsOcean/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ProjectHandler(c echo.Context) error {
	projectID := c.Param("id")
	if projectID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": "null project id"})
	}

	project := new(model.ProjectModel)
	if err := project.Find(projectID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}

	// this is done. now we return the response.
	// however, we can find owener and participant names and include that in the result.
	// for this simple project, we just leave this idea behind.
	return c.JSON(http.StatusOK, project)
}

func ProjectsHandler(c echo.Context) error {
	results, err := model.GetAllProjects()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	return c.JSON(http.StatusInternalServerError, results)
}

type CreateProjectRequest struct {
	Name            string   `json:"name" validate:"required"`
	OwenerID        string   `json:"owener_id" validate:"required,uuid4"`
	State           string   `json:"state" validate:"required,oneof=planned active done failed"`
	Progress        uint     `json:"progress" validate:"min=0,max=100"`
	ParticipantsIDs []string `json:"participants_ids" validate:"unique,dive,uuid4"`
}

func CreateProjectHandler(c echo.Context) error {
	// Restrictions to check:
	//		1. Name, OwenerID, and State should not be empty. (however, they are not
	//			 required to be unique.)
	//		2. Only managers can be oweners.
	//		3. Participants must be in the same department with the owener.
	request := new(CreateProjectRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}

	owener := model.EmployeeModel{}
	if err := owener.Find(request.OwenerID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}

	if owener.Role != "manager" {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": "owener is not a manager."})
	}

	if request.State != "active" {
		request.Progress = 0
	}

	participants, err := model.FindEmployees(request.ParticipantsIDs)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}

	for _, participant := range participants {
		if participant.Department != owener.Department {
			errMsg := fmt.Sprintf("participant %s not in department %s", participant.ID, owener.Department)
			return c.JSON(http.StatusBadRequest, map[string]string{"err": errMsg})
		}
	}

	project := model.ProjectModel{
		Name:            request.Name,
		OwenerID:        request.OwenerID,
		State:           request.State,
		Progress:        request.Progress,
		ParticipantsIDs: request.ParticipantsIDs,
	}

	id, err := project.Create()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"msg": "Project created.",
		"id":  id.Hex(),
	})
}
