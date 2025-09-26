package models

type Tag struct {
	Name                     string  `json:"name"`
	Id                       string  `json:"id"`
	Color                    string  `json:"color,omitempty"`
	Deleted                  bool    `json:"deleted,omitempty"`
	Hidden                   bool    `json:"hidden,omitempty"`
	External                 bool    `json:"external,omitempty"`
	Deprecated               bool    `json:"deprecated,omitempty"`
	CreationTimeInMillis     float32 `json:"creation_time_in_millis,omitempty"`
	ModificationTimeInMillis float32 `json:"modification_time_in_millis,omitempty"`
	AuthorId                 string  `json:"author_id,omitempty"`
	ModifierId               string  `json:"modifier_id,omitempty"`
	OwnerId                  string  `json:"owner_id,omitempty"`
}
