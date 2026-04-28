package models

type MetadataSearchRequest struct {
	Metadata                         []MetadataListItemInput      `json:"metadata"`
	Permissions                      []PermissionsInput           `json:"permissions"`
	CreatedByUserIdentifiers         []string                     `json:"created_by_user_identifiers,omitempty"`
	DependentObjectVersion           string                       `json:"dependent_object_version,omitempty"`
	ExcludeObjects                   ExcludeMetadataListItemInput `json:"exclude_objects,omitempty"`
	FavoriteObjectOptions            bool                         `json:"favorite_object_options,omitempty"`
	IncludeAutoCreatedObjects        bool                         `json:"include_auto_created_objects,omitempty"`
	IncludeDependentObjects          bool                         `json:"include_dependent_objects,omitempty"`
	DependentObjectRecordSize        int64                        `json:"dependent_objects_record_size,omitempty"`
	IncludeDetails                   bool                         `json:"include_details,omitempty"`
	IncludeHeaders                   bool                         `json:"include_headers,omitempty"`
	IncludeIncompleteObjects         bool                         `json:"include_incomplete_objects,omitempty"`
	IncludeVisualizationHeaders      bool                         `json:"include_visualization_headers,omitempty"`
	IncludeWorksheetSearchAssistData bool                         `json:"include_worksheet_search_assist_data,omitempty"`
	ModifiedByUserIdentifiers        []string                     `json:"modified_by_user_identifiers,omitempty"`
	RecordOffset                     int64                        `json:"record_offset,omitempty"`
	RecordSize                       int64                        `json:"record_size,omitempty"`
	SortOptions                      MetadataSearchSortOptions    `json:"sort_options,omitempty"`
	TagIdentifiers                   []string                     `json:"tag_identifiers,omitempty"`
	IncludeStats                     bool                         `json:"include_stats,omitempty"`
	IncludeDiscoverableObjects       bool                         `json:"include_discoverable_objects,omitempty"`
	ShowResolvedParameters           bool                         `json:"show_resolved_parameters,omitempty"`
	LiveboardResponseVersion         string                       `json:"liveboard_response_version,omitempty"`
	IncludeOnlyPublishedObjects      bool                         `json:"include_only_published_objects,omitempty"`
}
