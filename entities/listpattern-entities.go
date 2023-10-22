package entities

type Model_listpattern struct {
	Listpattern_id            string `json:"listpattern_id"`
	Listpattern_nmlistpattern string `json:"listpattern_nmlistpattern"`
	Listpattern_nmpoin        string `json:"listpattern_nmpoin"`
	Listpattern_totallose     int    `json:"listpattern_totallose"`
	Listpattern_totalwin      int    `json:"listpattern_totalwin"`
	Listpattern_status        string `json:"listpattern_status"`
	Listpattern_status_css    string `json:"listpattern_status_css"`
	Listpattern_create        string `json:"listpattern_create"`
	Listpattern_update        string `json:"listpattern_update"`
}
type Model_listpatterndetail struct {
	Listpatterndetail_id         int    `json:"listpatterndetail_id"`
	Listpatterndetail_idpoin     int    `json:"listpatterndetail_idpoin"`
	Listpatterndetail_nmpoin     string `json:"listpatterndetail_nmpoin"`
	Listpatterndetail_poin       int    `json:"listpatterndetail_poin"`
	Listpatterndetail_status     string `json:"listpatterndetail_status"`
	Listpatterndetail_status_css string `json:"listpatterndetail_status_css"`
	Listpatterndetail_create     string `json:"listpatterndetail_create"`
	Listpatterndetail_update     string `json:"listpatterndetail_update"`
}

type Controller_listpattern struct {
	Listpattern_search        string `json:"listpattern_search"`
	Listpattern_search_status string `json:"listpattern_search_status"`
	Listpattern_page          int    `json:"listpattern_page"`
}
type Controller_listpatterndetail struct {
	Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern"`
}
type Controller_listpatternsave struct {
	Page                      string `json:"page" validate:"required"`
	Sdata                     string `json:"sdata" validate:"required"`
	Listpattern_search        string `json:"listpattern_search"`
	Listpattern_search_status string `json:"listpattern_search_status"`
	Listpattern_page          int    `json:"listpattern_page"`
	Listpattern_id            string `json:"listpattern_id"`
	Listpattern_nmlistpattern string `json:"listpattern_nmlistpattern" validate:"required"`
	Listpattern_status        string `json:"listpattern_status" alidate:"required"`
}
type Controller_listpatterndetailsave struct {
	Page                            string `json:"page" validate:"required"`
	Sdata                           string `json:"sdata" validate:"required"`
	Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern" validate:"required"`
	Listpatterndetail_status        string `json:"listpatterndetail_status" validate:"required"`
	Listpatterndetail_idpoin        int    `json:"listpatterndetail_idpoin" `
}
type Controller_listpatterndetaildelete struct {
	Page                            string `json:"page" validate:"required"`
	Sdata                           string `json:"sdata" validate:"required"`
	Listpatterndetail_id            int    `json:"listpatterndetail_id"`
	Listpatterndetail_idlistpattern string `json:"listpatterndetail_idlistpattern"`
	Listpatterndetail_tipe          string `json:"listpatterndetail_tipe"`
}
