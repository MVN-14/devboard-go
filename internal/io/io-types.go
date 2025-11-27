package io

type ApplicationData struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
