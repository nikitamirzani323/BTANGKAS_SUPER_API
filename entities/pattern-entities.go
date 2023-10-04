package entities

type Model_pattern struct {
	Pattern_id         string `json:"pattern_id"`
	Pattern_idcard     string `json:"pattern_idcard"`
	Pattern_idpoin     int    `json:"pattern_idpoin"`
	Pattern_nmpoin     string `json:"pattern_nmpoin"`
	Pattern_status     string `json:"pattern_status"`
	Pattern_status_css string `json:"pattern_status_css"`
	Pattern_create     string `json:"pattern_create"`
	Pattern_update     string `json:"pattern_update"`
}

type Controller_patternsave struct {
	Page         string `json:"page" validate:"required"`
	Sdata        string `json:"sdata" validate:"required"`
	Pattern_List string `json:"pattern_list" `
}
