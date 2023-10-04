package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_pattern_local = configs.DB_tbl_trx_pattern

func Fetch_patternHome() (helpers.Response, error) {
	var obj entities.Model_pattern
	var arraobj []entities.Model_pattern
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idpattern , idcard, idpoin, status_pattern,   
			create_pattern, to_char(COALESCE(createdate_pattern,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_pattern, to_char(COALESCE(updatedate_pattern,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_pattern_local + `  
			ORDER BY createdate_pattern DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpoin_db                                                                          int
			idpattern_db, idcard_db, status_pattern_db                                         string
			create_pattern_db, createdate_pattern_db, update_pattern_db, updatedate_pattern_db string
		)

		err = row.Scan(&idpattern_db, &idcard_db, &idpoin_db, &status_pattern_db,
			&create_pattern_db, &createdate_pattern_db, &update_pattern_db, &updatedate_pattern_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if create_pattern_db != "" {
			create = create_pattern_db + ", " + createdate_pattern_db
		}
		if update_pattern_db != "" {
			update = update_pattern_db + ", " + updatedate_pattern_db
		}
		if status_pattern_db == "Y" {
			status_css = configs.STATUS_RUNNING
		}
		obj.Pattern_id = idpattern_db
		obj.Pattern_idcard = idcard_db
		obj.Pattern_idpoin = idpoin_db
		obj.Pattern_nmpoin = _Get_infomasterpoint(idpoin_db)
		obj.Pattern_status = status_pattern_db
		obj.Pattern_status_css = status_css
		obj.Pattern_create = create
		obj.Pattern_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_pattern(admin, listpattern, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		log.Println(listpattern)
		json := []byte(listpattern)
		jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			idpattern, _ := jsonparser.GetString(value, "idpattern")
			idcard, _ := jsonparser.GetString(value, "idcard")
			point, _ := jsonparser.GetString(value, "point")
			status, _ := jsonparser.GetString(value, "status")
			log.Println("IDPATTERN :" + idpattern)
			flag = CheckDB(database_pattern_local, "idpattern", idpattern)
			if !flag {
				sql_insert := `
					insert into
					` + database_pattern_local + ` (
						idpattern ,idcard, idpoin, status_pattern,  
						create_pattern, createdate_pattern 
					) values (
						$1, $2, $3, $4,    
						$5, $6 
					)
				`

				flag_insert, msg_insert := Exec_SQL(sql_insert, database_pattern_local, "INSERT",
					idpattern, idcard, _Get_infomasterpointByCode(point), status,
					admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

				if flag_insert {
					msg = "Succes"
				} else {
					fmt.Println(msg_insert)
				}
			} else {
				msg = "Duplicate Entry"
			}
		})

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _Get_infomasterpointByCode(code string) int {
	con := db.CreateCon()
	ctx := context.Background()
	idpoin := 0
	sql_select := `SELECT
			idpoin    
			FROM ` + configs.DB_tbl_mst_listpoint + `  
			WHERE codepoin='` + code + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&idpoin); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return idpoin
}
