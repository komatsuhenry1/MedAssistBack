package dto
type DocumentInfoResponse struct {
	Name        string `json:"name"`        // Um nome amigável, ex: "Documento de Licença (COREN)"
	Type        string `json:"type"`        // Um identificador, ex: "license_document"
	DownloadURL string `json:"download_url"` // A URL para baixar o arquivo
}