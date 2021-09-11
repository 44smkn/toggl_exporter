package toggl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/model"
)

const (
	projectURI = "/projects"
)

// Project is the object representing project in toggl world.
// It is bound with Get project data API Response of toggl.
// See: https://github.com/toggl/toggl_api_docs/blob/master/chapters/projects.md#get-project-data
type Project struct {
	Data struct {
		ID        int       `json:"id"`
		Wid       int       `json:"wid"`
		Cid       int       `json:"cid"`
		Name      string    `json:"name"`
		Billable  bool      `json:"billable"`
		IsPrivate bool      `json:"is_private"`
		Active    bool      `json:"active"`
		At        time.Time `json:"at"`
		Template  bool      `json:"template"`
		Color     string    `json:"color"`
	} `json:"data"`
}

type ProjectRepository struct {
	*Client
}

// GetProject returns the object representing project in toggl_exporter.
// It retrive the project detail using passed pid and create peculiar toggl project object.
func (r *ProjectRepository) GetProject(ctx context.Context, pid string) (*model.Project, error) {
	uri := fmt.Sprintf("%s/%s", projectURI, pid)
	req, err := r.newRequest(ctx, http.MethodGet, uri, nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusForbidden:
		return nil, errors.New(fmt.Sprintf("APIKey may be not valid. status is %v", res.Status))
	}

	var rawProject *Project
	if err := decodeBody(res, &rawProject); err != nil {
		return nil, err
	}

	project := model.Project{
		ID:   rawProject.Data.ID,
		Name: rawProject.Data.Name,
	}
	return &project, nil

}
