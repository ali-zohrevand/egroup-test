package repo

import (
	"testing"
)

func TestGetLastProjects(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Basic test", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGitlabResponse, err := GetLastProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotGitlabResponse.Projects.Nodes) == 0 {
				t.Errorf("Did not recive any Node")
				return
			}
		})
	}
}
