package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/m4rk0p0p/go-demo/models"
)

type UserCtl struct {
	userIdPattern *regexp.Regexp
}

func newUserCtl() *UserCtl {
	return &UserCtl{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (ctl UserCtl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/users" {
		switch request.Method {
		case http.MethodGet:
			ctl.getAll(writer, request)
		case http.MethodPost:
			ctl.post(writer, request)
		default:
			writer.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := ctl.userIdPattern.FindStringSubmatch(request.URL.Path)
		// Check if regex is matched at all
		if len(matches) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		// Try to get the user id from the appropriate subgroup match
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
		}

		switch request.Method {
		case http.MethodGet:
			ctl.get(id, writer)
		case http.MethodPut:
			ctl.put(id, writer, request)
		case http.MethodDelete:
			ctl.delete(id, writer)
		default:
			writer.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (ctl *UserCtl) getAll(writer http.ResponseWriter, request *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), writer)
}

func (ctl *UserCtl) get(id int, writer http.ResponseWriter) {
	usr, err := models.GetUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(usr, writer)
}

func (ctl *UserCtl) post(writer http.ResponseWriter, request *http.Request) {
	usr, err := ctl.parseRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not parse User object"))
		return
	}

	usr, err = models.AddUser(usr)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(usr, writer)
}

func (ctl *UserCtl) put(id int, writer http.ResponseWriter, request *http.Request) {
	usr, err := ctl.parseRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not parse User object"))
		return
	}
	if id != usr.Id {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("ID of submitted user must match ID in the URL"))
		return
	}

	usr, err = models.UpdateUser(usr)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(usr, writer)
}

func (ctl *UserCtl) delete(id int, writer http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (ctl *UserCtl) parseRequest(request *http.Request) (models.User, error) {
	dec := json.NewDecoder(request.Body)
	var usr models.User
	err := dec.Decode(&usr)
	if err != nil {
		return models.User{}, err
	}
	return usr, nil
}
