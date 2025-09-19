package dto
type DocumentInfoResponse struct {
	Name        string `json:"name"`        // Um nome amigável, ex: "Documento de Licença (COREN)"
	Type        string `json:"type"`        // Um identificador, ex: "license_document"
	DownloadURL string `json:"download_url"` // A URL para baixar o arquivo
}

type DashboardAdminDataResponse struct {
	TotalNurses int `json:"total_nurses"`
	TotalPatients int `json:"total_patients"`
	VisitsToday int `json:"visits_today"`
	NumberVisits int `json:"number_visits"`
	PendentApprovations int `json:"pendent_approvations"`
	NursesIDsPendentApprovations []string `json:"nurses_ids_pendent_approvations"`
}