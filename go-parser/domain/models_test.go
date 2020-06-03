package domain

import (
	"fmt"
	"testing"
)

//TestResource - test resource model
func TestArticles(t *testing.T) {
	resources := []Resource{
		Resource{ID: 1, Payload: "{}", Pattern: ""},
		Resource{ID: 2, Payload: "{}", Pattern: ""},
	}

	for i, tt := range resources {
		testname := fmt.Sprintf("%d", tt.ID)
		t.Run(testname, func(t *testing.T) {
			if resources[i].ID != tt.ID {
				t.Errorf("error")
			}
		})
	}
}
