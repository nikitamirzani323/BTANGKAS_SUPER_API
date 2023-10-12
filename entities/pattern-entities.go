package entities

type Model_pattern struct {
	Pattern_id            string `json:"pattern_id"`
	Pattern_idcard        string `json:"pattern_idcard"`
	Pattern_codepoin      string `json:"pattern_codepoin"`
	Pattern_nmpoin        string `json:"pattern_nmpoin"`
	Pattern_resultcardwin string `json:"pattern_resultcardwin"`
	Pattern_status        string `json:"pattern_status"`
	Pattern_status_css    string `json:"pattern_status_css"`
	Pattern_create        string `json:"pattern_create"`
	Pattern_update        string `json:"pattern_update"`
}
type Model_patternlistpoint struct {
	Patternlistpoint_id       int    `json:"patternlistpoint_id"`
	Patternlistpoint_codepoin string `json:"patternlistpoint_codepoin"`
	Patternlistpoint_nmpoin   string `json:"patternlistpoint_nmpoin"`
	Patternlistpoint_total    int    `json:"patternlistpoint_total"`
}

type Controller_pattern struct {
	Pattern_search        string `json:"pattern_search"`
	Pattern_search_status string `json:"pattern_search_status"`
	Pattern_page          int    `json:"pattern_page"`
}
type Controller_patternbycode struct {
	Pattern_code string `json:"pattern_code"`
	Pattern_page int    `json:"pattern_page"`
}
type Controller_patternsave struct {
	Page                  string `json:"page" validate:"required"`
	Sdata                 string `json:"sdata" validate:"required"`
	Pattern_search        string `json:"pattern_search"`
	Pattern_search_status string `json:"pattern_search_status"`
	Pattern_page          int    `json:"pattern_page"`
	Pattern_id            string `json:"pattern_id"`
	Pattern_codepoin      string `json:"pattern_codepoin"`
	Pattern_status        string `json:"pattern_status"`
	Pattern_resultcardwin string `json:"pattern_resultcardwin"`
	Pattern_List          string `json:"pattern_list" `
}
type Controller_patternsavemanual struct {
	Page                  string `json:"page" validate:"required"`
	Sdata                 string `json:"sdata" validate:"required"`
	Pattern_search        string `json:"pattern_search"`
	Pattern_search_status string `json:"pattern_search_status"`
	Pattern_page          int    `json:"pattern_page"`
	Pattern_id            string `json:"pattern_id"`
	Pattern_idcard        string `json:"pattern_idcard"`
	Pattern_codepoin      string `json:"pattern_codepoin"`
	Pattern_resultcardwin string `json:"pattern_resultcardwin"`
	Pattern_status        string `json:"pattern_status"`
}
