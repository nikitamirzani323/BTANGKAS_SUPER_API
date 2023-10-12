package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldpattern_home_redis = "LISTPATTERN_BACKEND"
const Fieldpattern_home_client_redis = "LISTPATTERN_FRONTEND"

func Patternhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pattern)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	fmt.Println(client.Pattern_page)
	if client.Pattern_search != "" {
		val_pattern := helpers.DeleteRedis(Fieldpattern_home_redis + "_" + strconv.Itoa(client.Pattern_page) + "_" + client.Pattern_search + "_" + client.Pattern_search_status)
		fmt.Printf("Redis Delete BACKEND PATTERN : %d", val_pattern)
	}
	var obj entities.Model_pattern
	var arraobj []entities.Model_pattern
	var objlistpoint entities.Model_patternlistpoint
	var arraobjlistpoint []entities.Model_patternlistpoint
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpattern_home_redis + "_" + strconv.Itoa(client.Pattern_page) + "_" + client.Pattern_search + "_" + client.Pattern_search_status)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalwin_RD, _ := jsonparser.GetInt(jsonredis, "totalwin")
	totallose_RD, _ := jsonparser.GetInt(jsonredis, "totallose")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listpoint_RD, _, _, _ := jsonparser.Get(jsonredis, "listpoint")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pattern_id, _ := jsonparser.GetString(value, "pattern_id")
		pattern_idcard, _ := jsonparser.GetString(value, "pattern_idcard")
		pattern_codepoin, _ := jsonparser.GetString(value, "pattern_codepoin")
		pattern_nmpoin, _ := jsonparser.GetString(value, "pattern_nmpoin")
		pattern_resultcardwin, _ := jsonparser.GetString(value, "pattern_resultcardwin")
		pattern_status, _ := jsonparser.GetString(value, "pattern_status")
		pattern_status_css, _ := jsonparser.GetString(value, "pattern_status_css")
		pattern_create, _ := jsonparser.GetString(value, "pattern_create")
		pattern_update, _ := jsonparser.GetString(value, "pattern_update")

		obj.Pattern_id = pattern_id
		obj.Pattern_idcard = pattern_idcard
		obj.Pattern_codepoin = pattern_codepoin
		obj.Pattern_nmpoin = pattern_nmpoin
		obj.Pattern_resultcardwin = pattern_resultcardwin
		obj.Pattern_status = pattern_status
		obj.Pattern_status_css = pattern_status_css
		obj.Pattern_create = pattern_create
		obj.Pattern_update = pattern_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listpoint_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		patternlistpoint_id, _ := jsonparser.GetInt(value, "patternlistpoint_id")
		patternlistpoint_codepoin, _ := jsonparser.GetString(value, "patternlistpoint_codepoin")
		patternlistpoint_nmpoin, _ := jsonparser.GetString(value, "patternlistpoint_nmpoin")
		patternlistpoint_total, _ := jsonparser.GetInt(value, "patternlistpoint_total")

		objlistpoint.Patternlistpoint_id = int(patternlistpoint_id)
		objlistpoint.Patternlistpoint_codepoin = patternlistpoint_codepoin
		objlistpoint.Patternlistpoint_nmpoin = patternlistpoint_nmpoin
		objlistpoint.Patternlistpoint_total = int(patternlistpoint_total)
		arraobjlistpoint = append(arraobjlistpoint, objlistpoint)
	})
	if !flag {
		//search string, page int
		result, err := models.Fetch_patternHome(client.Pattern_search, client.Pattern_search_status, client.Pattern_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpattern_home_redis+"_"+strconv.Itoa(client.Pattern_page)+"_"+client.Pattern_search+"_"+client.Pattern_search_status, result, 60*time.Minute)
		fmt.Println("PATTERN MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PATTERN CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"totalwin":    totalwin_RD,
			"totallose":   totallose_RD,
			"listpoint":   arraobjlistpoint,
			"time":        time.Since(render_page).String(),
		})
	}
}
func PatternByPoin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_patternbycode)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	var obj entities.Model_pattern
	var arraobj []entities.Model_pattern
	var objlistpoint entities.Model_patternlistpoint
	var arraobjlistpoint []entities.Model_patternlistpoint
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpattern_home_redis + "_" + strconv.Itoa(client.Pattern_page) + "_" + client.Pattern_code)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalwin_RD, _ := jsonparser.GetInt(jsonredis, "totalwin")
	totallose_RD, _ := jsonparser.GetInt(jsonredis, "totallose")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listpoint_RD, _, _, _ := jsonparser.Get(jsonredis, "listpoint")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pattern_id, _ := jsonparser.GetString(value, "pattern_id")
		pattern_idcard, _ := jsonparser.GetString(value, "pattern_idcard")
		pattern_codepoin, _ := jsonparser.GetString(value, "pattern_codepoin")
		pattern_nmpoin, _ := jsonparser.GetString(value, "pattern_nmpoin")
		pattern_resultcardwin, _ := jsonparser.GetString(value, "pattern_resultcardwin")
		pattern_status, _ := jsonparser.GetString(value, "pattern_status")
		pattern_status_css, _ := jsonparser.GetString(value, "pattern_status_css")
		pattern_create, _ := jsonparser.GetString(value, "pattern_create")
		pattern_update, _ := jsonparser.GetString(value, "pattern_update")

		obj.Pattern_id = pattern_id
		obj.Pattern_idcard = pattern_idcard
		obj.Pattern_codepoin = pattern_codepoin
		obj.Pattern_nmpoin = pattern_nmpoin
		obj.Pattern_resultcardwin = pattern_resultcardwin
		obj.Pattern_status = pattern_status
		obj.Pattern_status_css = pattern_status_css
		obj.Pattern_create = pattern_create
		obj.Pattern_update = pattern_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listpoint_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		patternlistpoint_id, _ := jsonparser.GetInt(value, "patternlistpoint_id")
		patternlistpoint_codepoin, _ := jsonparser.GetString(value, "patternlistpoint_codepoin")
		patternlistpoint_nmpoin, _ := jsonparser.GetString(value, "patternlistpoint_nmpoin")
		patternlistpoint_total, _ := jsonparser.GetInt(value, "patternlistpoint_total")

		objlistpoint.Patternlistpoint_id = int(patternlistpoint_id)
		objlistpoint.Patternlistpoint_codepoin = patternlistpoint_codepoin
		objlistpoint.Patternlistpoint_nmpoin = patternlistpoint_nmpoin
		objlistpoint.Patternlistpoint_total = int(patternlistpoint_total)
		arraobjlistpoint = append(arraobjlistpoint, objlistpoint)
	})
	if !flag {
		//search string, page int
		result, err := models.Fetch_patternByPoin(client.Pattern_code, client.Pattern_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpattern_home_redis+"_"+strconv.Itoa(client.Pattern_page)+"_"+client.Pattern_code, result, 60*time.Minute)
		fmt.Println("PATTERN BY CODE POINT MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PATTERN BY CODE POINT")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"totalwin":    totalwin_RD,
			"totallose":   totallose_RD,
			"listpoint":   arraobjlistpoint,
			"time":        time.Since(render_page).String(),
		})
	}
}
func PatternSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_patternsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, listpattern, resultcardwin, codepoin, status, idrecord, sData string
	result, err := models.Save_pattern(
		client_admin,
		client.Pattern_List, client.Pattern_resultcardwin, client.Pattern_codepoin, client.Pattern_status, client.Pattern_id, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_pattern(client.Pattern_search, client.Pattern_search_status, client.Pattern_page)
	return c.JSON(result)
}
func PatternSavemanual(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_patternsavemanual)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idpattern, idcard, codepoin, resultwin, status, sData string
	result, err := models.Save_patternmanual(
		client_admin,
		client.Pattern_id, client.Pattern_idcard, client.Pattern_codepoin,
		client.Pattern_resultcardwin, client.Pattern_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_pattern(client.Pattern_search, client.Pattern_search_status, client.Pattern_page)
	return c.JSON(result)
}
func _deleteredis_pattern(search, status string, page int) {
	val_master := helpers.DeleteRedis(Fieldpattern_home_redis + "_" + strconv.Itoa(page) + "_" + search + "_" + status)
	fmt.Printf("Redis Delete BACKEND PATTERN : %d\n", val_master)

	val_client := helpers.DeleteRedis(Fieldpattern_home_client_redis)
	fmt.Printf("Redis Delete CLIENT PATTERN : %d", val_client)

}
