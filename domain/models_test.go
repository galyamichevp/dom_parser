package domain

import (
	"fmt"
	"testing"
)

//TestArticles - test article model
func TestArticles(t *testing.T) {
	articles := []article{
		article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
		article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	}

	for i, tt := range articles {
		testname := fmt.Sprintf("%d", tt.ID)
		t.Run(testname, func(t *testing.T) {
			if articles[i].ID != tt.ID {
				t.Errorf("error")
			}
		})
	}
}
