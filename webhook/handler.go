// Package webhook defines payload format of GitLab webhook
//
// It also provides a simple wrapper for you to easily create
// applications handle webhooks.
package webhook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Handler handles webhook requests and dispatch events via channel.
// You just set channels you want to receive, and pass the channel to your
// application.
//
//     func handlePushEvent(ch chan PushEvent) {
//         for e := range ch {
//             log.Printf(
//                 "%s pushed %d commits to %s",
//                 e.UserName,
//                 e.TotalCommitCount,
//                 e.Project.PathWithNamespace,
//             )
//         }
//     }
//
//     handler := &Handler{Push: make(chan PushEvent)}
//     go handlePushEvent(handler.Push)
//     http.Handle("/webhook", handler)
//     http.ListenAndServe(":80", nil)
type Handler struct {
	Push         chan PushEvent
	Tag          chan TagEvent
	Issue        chan IssueEvent
	Comment      chan CommentEvent
	MergeRequest chan MergeRequestEvent
	Wiki         chan WikiEvent
	Pipeline     chan PipelineEvent
}

func jsonHelper(w http.ResponseWriter, r *http.Request, data interface{}) (err error) {
	defer r.Body.Close()
	defer func() {
		if err == nil {
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}()

	var buf []byte
	buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, data)
	return err
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := strings.TrimSpace(r.Header.Get("X-Gitlab-Event"))
	switch t {
	case "Push Hook":
		var data PushEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Push != nil {
			h.Push <- data
		}
	case "Tag Push Hook":
		var data TagEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Tag != nil {
			h.Tag <- data
		}
	case "Issue Hook":
		var data IssueEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Issue != nil {
			h.Issue <- data
		}
	case "Note Hook":
		var data CommentEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Comment != nil {
			h.Comment <- data
		}
	case "Merge Request Hook":
		var data MergeRequestEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.MergeRequest != nil {
			h.MergeRequest <- data
		}
	case "Wiki Page Hook":
		var data WikiEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Wiki != nil {
			h.Wiki <- data
		}
	case "Pipeline Hook":
		var data PipelineEvent
		if err := jsonHelper(w, r, &data); err != nil {
			return
		}
		if h.Pipeline != nil {
			h.Pipeline <- data
		}
	}
}
