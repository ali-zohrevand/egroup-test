package usecase

import (
	"testing"
)

func TestGetLastProjects(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Basic Get All", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGitlabLastProjects, err := GetLastProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotGitlabLastProjects.Names == "" {
				t.Errorf("GetLastProjects() error = %v", "Did not get any name")
				return
			}
		})
	}
}
