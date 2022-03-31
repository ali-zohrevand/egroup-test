package service

import "testing"

func TestGetLastProjects(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Get Basic", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames, gotNumberOfForks, err := GetLastProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNames == "" || gotNumberOfForks == 0 {
				t.Errorf("GetLastProjects() error = %v, ", "Did not get any Node")
				return
			}
		})
	}
}
