package structs

type MediaFile struct {
	ID          int64  `json:"id,omitempty"`
	FileName    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	Path        string `json:"path"`
	UserID      string `json:"user_id,omitempty"`
	HTTPJson    string `json:"http_json,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}