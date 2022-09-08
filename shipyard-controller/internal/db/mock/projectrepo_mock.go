// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db_mock

import (
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	mvmodels "github.com/keptn/keptn/shipyard-controller/internal/db/models/projects_mv"
	"sync"
)

// ProjectRepoMock is a mock implementation of db.ProjectRepo.
//
// 	func TestSomethingThatUsesProjectRepo(t *testing.T) {
//
// 		// make and configure a mocked db.ProjectRepo
// 		mockedProjectRepo := &ProjectRepoMock{
// 			CreateProjectFunc: func(project *apimodels.ExpandedProject) error {
// 				panic("mock out the CreateProject method")
// 			},
// 			DeleteProjectFunc: func(projectName string) error {
// 				panic("mock out the DeleteProject method")
// 			},
// 			GetProjectFunc: func(projectName string) (*apimodels.ExpandedProject, error) {
// 				panic("mock out the GetProject method")
// 			},
// 			GetProjectsFunc: func() ([]*apimodels.ExpandedProject, error) {
// 				panic("mock out the GetProjects method")
// 			},
// 			UpdateProjectFunc: func(project *apimodels.ExpandedProject) error {
// 				panic("mock out the UpdateProject method")
// 			},
// 			UpdateProjectServiceFunc: func(projectName string, stageName string, serviceName string, properties mvmodels.ServiceUpdate) error {
// 				panic("mock out the UpdateProjectService method")
// 			},
// 			UpdateProjectUpstreamFunc: func(projectName string, uri string, user string) error {
// 				panic("mock out the UpdateProjectUpstream method")
// 			},
// 		}
//
// 		// use mockedProjectRepo in code that requires db.ProjectRepo
// 		// and then make assertions.
//
// 	}
type ProjectRepoMock struct {
	// CreateProjectFunc mocks the CreateProject method.
	CreateProjectFunc func(project *apimodels.ExpandedProject) error

	// DeleteProjectFunc mocks the DeleteProject method.
	DeleteProjectFunc func(projectName string) error

	// GetProjectFunc mocks the GetProject method.
	GetProjectFunc func(projectName string) (*apimodels.ExpandedProject, error)

	// GetProjectsFunc mocks the GetProjects method.
	GetProjectsFunc func() ([]*apimodels.ExpandedProject, error)

	// UpdateProjectFunc mocks the UpdateProject method.
	UpdateProjectFunc func(project *apimodels.ExpandedProject) error

	// UpdateProjectServiceFunc mocks the UpdateProjectService method.
	UpdateProjectServiceFunc func(projectName string, stageName string, serviceName string, properties mvmodels.ServiceUpdate) error

	// UpdateProjectUpstreamFunc mocks the UpdateProjectUpstream method.
	UpdateProjectUpstreamFunc func(projectName string, uri string, user string) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateProject holds details about calls to the CreateProject method.
		CreateProject []struct {
			// Project is the project argument value.
			Project *apimodels.ExpandedProject
		}
		// DeleteProject holds details about calls to the DeleteProject method.
		DeleteProject []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
		// GetProject holds details about calls to the GetProject method.
		GetProject []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
		// GetProjects holds details about calls to the GetProjects method.
		GetProjects []struct {
		}
		// UpdateProject holds details about calls to the UpdateProject method.
		UpdateProject []struct {
			// Project is the project argument value.
			Project *apimodels.ExpandedProject
		}
		// UpdateProjectService holds details about calls to the UpdateProjectService method.
		UpdateProjectService []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// StageName is the stageName argument value.
			StageName string
			// ServiceName is the serviceName argument value.
			ServiceName string
			// Properties is the properties argument value.
			Properties mvmodels.ServiceUpdate
		}
		// UpdateProjectUpstream holds details about calls to the UpdateProjectUpstream method.
		UpdateProjectUpstream []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
			// URI is the uri argument value.
			URI string
			// User is the user argument value.
			User string
		}
	}
	lockCreateProject         sync.RWMutex
	lockDeleteProject         sync.RWMutex
	lockGetProject            sync.RWMutex
	lockGetProjects           sync.RWMutex
	lockUpdateProject         sync.RWMutex
	lockUpdateProjectService  sync.RWMutex
	lockUpdateProjectUpstream sync.RWMutex
}

// CreateProject calls CreateProjectFunc.
func (mock *ProjectRepoMock) CreateProject(project *apimodels.ExpandedProject) error {
	if mock.CreateProjectFunc == nil {
		panic("ProjectRepoMock.CreateProjectFunc: method is nil but ProjectRepo.CreateProject was just called")
	}
	callInfo := struct {
		Project *apimodels.ExpandedProject
	}{
		Project: project,
	}
	mock.lockCreateProject.Lock()
	mock.calls.CreateProject = append(mock.calls.CreateProject, callInfo)
	mock.lockCreateProject.Unlock()
	return mock.CreateProjectFunc(project)
}

// CreateProjectCalls gets all the calls that were made to CreateProject.
// Check the length with:
//     len(mockedProjectRepo.CreateProjectCalls())
func (mock *ProjectRepoMock) CreateProjectCalls() []struct {
	Project *apimodels.ExpandedProject
} {
	var calls []struct {
		Project *apimodels.ExpandedProject
	}
	mock.lockCreateProject.RLock()
	calls = mock.calls.CreateProject
	mock.lockCreateProject.RUnlock()
	return calls
}

// DeleteProject calls DeleteProjectFunc.
func (mock *ProjectRepoMock) DeleteProject(projectName string) error {
	if mock.DeleteProjectFunc == nil {
		panic("ProjectRepoMock.DeleteProjectFunc: method is nil but ProjectRepo.DeleteProject was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockDeleteProject.Lock()
	mock.calls.DeleteProject = append(mock.calls.DeleteProject, callInfo)
	mock.lockDeleteProject.Unlock()
	return mock.DeleteProjectFunc(projectName)
}

// DeleteProjectCalls gets all the calls that were made to DeleteProject.
// Check the length with:
//     len(mockedProjectRepo.DeleteProjectCalls())
func (mock *ProjectRepoMock) DeleteProjectCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockDeleteProject.RLock()
	calls = mock.calls.DeleteProject
	mock.lockDeleteProject.RUnlock()
	return calls
}

// GetProject calls GetProjectFunc.
func (mock *ProjectRepoMock) GetProject(projectName string) (*apimodels.ExpandedProject, error) {
	if mock.GetProjectFunc == nil {
		panic("ProjectRepoMock.GetProjectFunc: method is nil but ProjectRepo.GetProject was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockGetProject.Lock()
	mock.calls.GetProject = append(mock.calls.GetProject, callInfo)
	mock.lockGetProject.Unlock()
	return mock.GetProjectFunc(projectName)
}

// GetProjectCalls gets all the calls that were made to GetProject.
// Check the length with:
//     len(mockedProjectRepo.GetProjectCalls())
func (mock *ProjectRepoMock) GetProjectCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockGetProject.RLock()
	calls = mock.calls.GetProject
	mock.lockGetProject.RUnlock()
	return calls
}

// GetProjects calls GetProjectsFunc.
func (mock *ProjectRepoMock) GetProjects() ([]*apimodels.ExpandedProject, error) {
	if mock.GetProjectsFunc == nil {
		panic("ProjectRepoMock.GetProjectsFunc: method is nil but ProjectRepo.GetProjects was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetProjects.Lock()
	mock.calls.GetProjects = append(mock.calls.GetProjects, callInfo)
	mock.lockGetProjects.Unlock()
	return mock.GetProjectsFunc()
}

// GetProjectsCalls gets all the calls that were made to GetProjects.
// Check the length with:
//     len(mockedProjectRepo.GetProjectsCalls())
func (mock *ProjectRepoMock) GetProjectsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetProjects.RLock()
	calls = mock.calls.GetProjects
	mock.lockGetProjects.RUnlock()
	return calls
}

// UpdateProject calls UpdateProjectFunc.
func (mock *ProjectRepoMock) UpdateProject(project *apimodels.ExpandedProject) error {
	if mock.UpdateProjectFunc == nil {
		panic("ProjectRepoMock.UpdateProjectFunc: method is nil but ProjectRepo.UpdateProject was just called")
	}
	callInfo := struct {
		Project *apimodels.ExpandedProject
	}{
		Project: project,
	}
	mock.lockUpdateProject.Lock()
	mock.calls.UpdateProject = append(mock.calls.UpdateProject, callInfo)
	mock.lockUpdateProject.Unlock()
	return mock.UpdateProjectFunc(project)
}

// UpdateProjectCalls gets all the calls that were made to UpdateProject.
// Check the length with:
//     len(mockedProjectRepo.UpdateProjectCalls())
func (mock *ProjectRepoMock) UpdateProjectCalls() []struct {
	Project *apimodels.ExpandedProject
} {
	var calls []struct {
		Project *apimodels.ExpandedProject
	}
	mock.lockUpdateProject.RLock()
	calls = mock.calls.UpdateProject
	mock.lockUpdateProject.RUnlock()
	return calls
}

// UpdateProjectService calls UpdateProjectServiceFunc.
func (mock *ProjectRepoMock) UpdateProjectService(projectName string, stageName string, serviceName string, properties mvmodels.ServiceUpdate) error {
	if mock.UpdateProjectServiceFunc == nil {
		panic("ProjectRepoMock.UpdateProjectServiceFunc: method is nil but ProjectRepo.UpdateProjectService was just called")
	}
	callInfo := struct {
		ProjectName string
		StageName   string
		ServiceName string
		Properties  mvmodels.ServiceUpdate
	}{
		ProjectName: projectName,
		StageName:   stageName,
		ServiceName: serviceName,
		Properties:  properties,
	}
	mock.lockUpdateProjectService.Lock()
	mock.calls.UpdateProjectService = append(mock.calls.UpdateProjectService, callInfo)
	mock.lockUpdateProjectService.Unlock()
	return mock.UpdateProjectServiceFunc(projectName, stageName, serviceName, properties)
}

// UpdateProjectServiceCalls gets all the calls that were made to UpdateProjectService.
// Check the length with:
//     len(mockedProjectRepo.UpdateProjectServiceCalls())
func (mock *ProjectRepoMock) UpdateProjectServiceCalls() []struct {
	ProjectName string
	StageName   string
	ServiceName string
	Properties  mvmodels.ServiceUpdate
} {
	var calls []struct {
		ProjectName string
		StageName   string
		ServiceName string
		Properties  mvmodels.ServiceUpdate
	}
	mock.lockUpdateProjectService.RLock()
	calls = mock.calls.UpdateProjectService
	mock.lockUpdateProjectService.RUnlock()
	return calls
}

// UpdateProjectUpstream calls UpdateProjectUpstreamFunc.
func (mock *ProjectRepoMock) UpdateProjectUpstream(projectName string, uri string, user string) error {
	if mock.UpdateProjectUpstreamFunc == nil {
		panic("ProjectRepoMock.UpdateProjectUpstreamFunc: method is nil but ProjectRepo.UpdateProjectUpstream was just called")
	}
	callInfo := struct {
		ProjectName string
		URI         string
		User        string
	}{
		ProjectName: projectName,
		URI:         uri,
		User:        user,
	}
	mock.lockUpdateProjectUpstream.Lock()
	mock.calls.UpdateProjectUpstream = append(mock.calls.UpdateProjectUpstream, callInfo)
	mock.lockUpdateProjectUpstream.Unlock()
	return mock.UpdateProjectUpstreamFunc(projectName, uri, user)
}

// UpdateProjectUpstreamCalls gets all the calls that were made to UpdateProjectUpstream.
// Check the length with:
//     len(mockedProjectRepo.UpdateProjectUpstreamCalls())
func (mock *ProjectRepoMock) UpdateProjectUpstreamCalls() []struct {
	ProjectName string
	URI         string
	User        string
} {
	var calls []struct {
		ProjectName string
		URI         string
		User        string
	}
	mock.lockUpdateProjectUpstream.RLock()
	calls = mock.calls.UpdateProjectUpstream
	mock.lockUpdateProjectUpstream.RUnlock()
	return calls
}
