package teamcity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
)

// Project is the model for project entities in TeamCity
type Project struct {

	// archived
	Archived *bool `json:"archived,omitempty" xml:"archived"`

	// build types
	// BuildTypes *BuildTypes `json:"buildTypes,omitempty"`

	// // default template
	// DefaultTemplate *BuildType `json:"defaultTemplate,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// internal Id
	InternalID string `json:"internalId,omitempty" xml:"internalId"`

	// links
	// Links *Links `json:"links,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// Parameters for the project. Read-only, only useful when retrieving project details
	Parameters *Properties `json:"parameters,omitempty"` //TODO: Encapsulate field

	// parent project
	ParentProject *Project `json:"parentProject,omitempty"`

	// parent project Id
	ParentProjectID string `json:"parentProjectId,omitempty" xml:"parentProjectId"`

	// parent project internal Id
	ParentProjectInternalID string `json:"parentProjectInternalId,omitempty" xml:"parentProjectInternalId"`

	// parent project name
	ParentProjectName string `json:"parentProjectName,omitempty" xml:"parentProjectName"`

	// project features
	// ProjectFeatures *ProjectFeatures `json:"projectFeatures,omitempty"`

	// projects
	// Projects *Projects `json:"projects,omitempty"`

	// // read only UI
	// ReadOnlyUI *StateField `json:"readOnlyUI,omitempty"`

	// // templates
	// Templates *BuildTypes `json:"templates,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty" xml:"uuid"`

	// vcs roots
	// VcsRoots *VcsRoots `json:"vcsRoots,omitempty"`

	// web Url
	WebURL string `json:"webUrl,omitempty" xml:"webUrl"`
}

// ProjectReference contains basic information, usually enough to use as a type for relationships.
// In addition to that, TeamCity does not return the full detailed representation when creating objects, thus the need for a reference.
type ProjectReference struct {
	// id
	ID string `json:"id,omitempty" xml:"id"`
	// name
	Name string `json:"name,omitempty" xml:"name"`
}

// ProjectService has operations for handling projects
type ProjectService struct {
	sling *sling.Sling
}

func newProjectService(base *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: base.Path("projects/"),
	}
}

// Create creates a new project at root project level
func (s *ProjectService) Create(project *Project) (*ProjectReference, error) {
	var created ProjectReference
	if err := project.Validate(); err != nil {
		return nil, err
	}

	response, err := s.sling.New().BodyJSON(project).Post("").ReceiveSuccess(&created)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 400 {
		return nil, fmt.Errorf("A project with name '%s' already exists", project.Name)
	}

	return &created, nil
}

// GetById Retrieves a project resource by ID
func (s *ProjectService) GetById(id string) (*Project, error) {
	var out Project

	resp, err := s.sling.New().Get(LocatorId(id).String()).ReceiveSuccess(&out)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error when retrieving Project id = '%s', status: %d", id, resp.StatusCode)
	}

	return &out, err
}

//Delete - Deletes a project
func (s *ProjectService) Delete(id string) error {
	request, _ := s.sling.New().Delete(id).Request()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode == 204 {
		return nil
	}

	if response.StatusCode != 200 && response.StatusCode != 204 {
		respData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Error '%d' when deleting project: %s", response.StatusCode, string(respData))
	}

	return nil
}

// Validate validates this project
func (m *Project) Validate() error {
	//var res []error

	if len(m.Name) <= 0 {
		return errors.New("Project must have a name")
	}

	return nil
}
