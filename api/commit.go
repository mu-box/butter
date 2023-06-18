package api

import (
	"net/http"
	"strconv"

	"github.com/mu-box/butter/repo"
)

func showCommits(rw http.ResponseWriter, req *http.Request) {
	page, _ := strconv.Atoi(req.FormValue("page"))
	branch := req.FormValue("branch")
	if branch == "" {
		branch = "main"
	}
	commits, err := repo.ListCommits(branch, page)
	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeBody(commits, rw, http.StatusOK)
}

func showCommitDetails(rw http.ResponseWriter, req *http.Request) {
	commit, err := repo.GetCommit(req.URL.Query().Get(":commit"))
	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeBody(commit, rw, http.StatusOK)
}
