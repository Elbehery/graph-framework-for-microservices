// Code generated by github.com/vmware-tanzu/graph-framework-for-microservices/src/gqlgen, DO NOT EDIT.

package config

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID          string `json:"id"`
	DatabaseID  int    `json:"databaseId"`
	Description string `json:"text"`
	Done        bool   `json:"done"`
	User        *User  `json:"user"`
}
