package handler

import (
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/view"
	"html/template"
	"net/http"
)

type Templates struct {
	templates       map[string]*template.Template
	currentTemplate *template.Template
}

func jobViewHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("pageView").Parse(view.PageView))
	t = template.Must(t.New("jobView").Parse(view.JobView))

	q := job.FromContext(r.Context())
	jobStatusMap := q.GetJobStatuses()

	jobs := []job.Job{}
	jobStatusMap.Range(func(key, value interface{}) bool {
		jobs = append(jobs, *value.(*job.Job))
		return true
	})

	if r.Header.Get("HX-Request") == "true" {
		err := t.ExecuteTemplate(w, "jobView", jobs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := t.ExecuteTemplate(w, "pageView", jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
