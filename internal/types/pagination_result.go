package types

type PaginationResult struct {
	Results []InstallationTemplate `json:"results"`
	Count   int                    `json:"count"`
}
