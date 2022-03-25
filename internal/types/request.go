package types

type DownloadTemplateRequest struct {
	Template string            `json:"template"`
	Envs     map[string]string `json:"envs"`
}
