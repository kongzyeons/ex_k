package models

type AdminGetQueryRequest struct {
	ParentID   int  `json:"parent_id"`
	IconID     int  `json:"icon_id"`
	IsChildren bool `json:"is_children"`
}

type AdminGetQueryResponse struct {
	ParentID   int        `json:"id"`
	Name       string     `json:"name"`
	Route      string     `json:"route"`
	Icon       string     `json:"icon"`
	IsChildren bool       `json:"is_children"`
	Children   []Children `json:"children"`
}

type Children struct {
	IconID   int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Route    string `json:"route"`
}
