package controller

import (
	"dalas98/golang-lesson/Session-9/httpclient"
	"dalas98/golang-lesson/Session-9/params"
	"encoding/json"
	"net/http"
)

type PostController struct {
	client *httpclient.Client
}

func NewPostController(client *httpclient.Client) *PostController {
	return &PostController{
		client: client,
	}
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := p.client.Get("/")
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL_SERVER_ERROR",
			Error:   err.Error(),
		})
		return
	}

	var posts []Post
	err = json.Unmarshal(resp, &posts)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL_SERVER_ERROR",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(w, &Response{
		Status:  http.StatusOK,
		Message: "GET_POSTS_SUCCESS",
		Payload: posts,
	})
}

func (p *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req params.PostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  400,
			Message: "BAD_REQUEST",
			Error:   err.Error(),
		})
		return
	}

	bytePayload, err := json.Marshal(req)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  400,
			Message: "BAD_REQUEST",
		})
		return
	}

	resp, err := p.client.Post("/", bytePayload)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  400,
			Message: "FAILED_TO_CREATE_NEW_POST",
		})
		return
	}

	var dataInterface interface{}
	err = json.Unmarshal(resp, &dataInterface)
	if err != nil {
		WriteJsonResponse(w, &Response{
			Status:  400,
			Message: "BAD_REQUEST",
		})
		return
	}

	WriteJsonResponse(w, &Response{
		Status:  http.StatusAccepted,
		Message: "CREATE_POST_SUCCESS",
		Payload: dataInterface,
	})
}
