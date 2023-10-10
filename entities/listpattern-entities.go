package entities

type Model_listpattern struct {
	Listpattern_id            string `json:"listpattern_id"`
	Listpattern_nmlistpattern string `json:"listpattern_nmlistpattern"`
	Listpattern_status        string `json:"listpattern_status"`
	Listpattern_status_css    string `json:"listpattern_status_css"`
	Listpattern_create        string `json:"listpattern_create"`
	Listpattern_update        string `json:"listpattern_update"`
}
type Model_listpatterndetail struct {
	Listpatterndetail_id         int    `json:"listpatterndetail_id"`
	Listpatterndetail_idpattern  string `json:"listpatterndetail_idpattern"`
	Listpatterndetail_nmpoin     string `json:"listpatterndetail_nmpoin"`
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
	Listpatterndetail_idpattern     string `json:"listpatterndetail_idpattern" validate:"required"`
}
type Controller_listpatterndetaildelete struct {
	Page                 string `json:"page" validate:"required"`
	Sdata                string `json:"sdata" validate:"required"`
	Listpatterndetail_id string `json:"listpatterndetail_id"`
}
