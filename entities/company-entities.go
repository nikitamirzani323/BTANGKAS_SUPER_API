package entities

type Model_company struct {
	Company_id         string `json:"company_id"`
	Company_startjoin  string `json:"company_startjoin"`
	Company_endjoin    string `json:"company_endjoin"`
	Company_name       string `json:"company_name"`
	Company_idcurr     string `json:"company_idcurr"`
	Company_nmowner    string `json:"company_nmowner"`
	Company_phoneowner string `json:"company_phoneowner"`
	Company_emailowner string `json:"company_emailowner"`
	Company_url        string `json:"company_url"`
	Company_status     string `json:"company_status"`
	Company_status_css string `json:"company_status_css"`
	Company_create     string `json:"company_create"`
	Company_update     string `json:"company_update"`
}
type Model_companyshare struct {
	Company_id   string `json:"company_id"`
	Company_name string `json:"company_name"`
}

type Controller_companysave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Company_id         string `json:"company_id"`
	Company_name       string `json:"company_name" validate:"required"`
	Company_idcurr     string `json:"company_idcurr" validate:"required"`
	Company_nmowner    string `json:"company_nmowner" validate:"required"`
	Company_phoneowner string `json:"company_phoneowner" validate:"required"`
	Company_emailowner string `json:"company_emailowner" `
	Company_url        string `json:"company_url" validate:"required"`
	Company_status     string `json:"company_status" validate:"required"`
}
