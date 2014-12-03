// +build go1.2

package models

type PresentationConfig struct {
	Repo    string       `json:"repository"`
	Commits []ConfigMeta `json:"commits"`
}

type ConfigMeta struct {
	Sha     string `json:"sha"`
	Message string `json:"message"`
}
