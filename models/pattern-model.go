package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
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

func Fetch_patternHome(search string, page int) (helpers.Responsepattern, error) {
	var obj entities.Model_pattern
	var arraobj []entities.Model_pattern
	var res helpers.Responsepattern
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	total_lose := _Get_pattern_losewin("N")
	total_win := _Get_pattern_losewin("Y")
	perpage := 50
	totalrecord := 0
	offset := page

	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(idpattern) as totalpattern  "
	sql_selectcount += "FROM " + database_pattern_local + "  "
	if search != "" {
		sql_selectcount += "AND LOWER(idpattern) LIKE '%" + strings.ToLower(search) + "%' "
	}

	row_selectcount := con.QueryRowContext(ctx, sql_selectcount)
	switch e_selectcount := row_selectcount.Scan(&totalrecord); e_selectcount {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e_selectcount)
	}

	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "idpattern , idcard, idpoin, resultcardwin, status_pattern, "
	sql_select += "create_pattern, to_char(COALESCE(createdate_pattern,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "update_pattern, to_char(COALESCE(updatedate_pattern,now()), 'YYYY-MM-DD HH24:MI:SS') "
	sql_select += "FROM " + database_pattern_local + "  "
	if search == "" {
		sql_select += "ORDER BY createdate_pattern DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)
	} else {
		sql_select += "WHERE LOWER(idpattern) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdate_pattern DESC  LIMIT " + strconv.Itoa(perpage)
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpoin_db                                                                          int
			idpattern_db, idcard_db, resultcardwin_db, status_pattern_db                       string
			create_pattern_db, createdate_pattern_db, update_pattern_db, updatedate_pattern_db string
		)

		err = row.Scan(&idpattern_db, &idcard_db, &idpoin_db, &resultcardwin_db, &status_pattern_db,
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
		obj.Pattern_resultcardwin = resultcardwin_db
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
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Totallose = total_lose
	res.Totalwin = total_win
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_pattern(admin, listpattern, resultcardwin, idrecord, sData string) (helpers.Response, error) {
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
			resultwin, _ := jsonparser.GetString(value, "resultwin")
			status, _ := jsonparser.GetString(value, "status")
			log.Println("IDPATTERN :" + idpattern)
			flag = CheckDB(database_pattern_local, "idpattern", idpattern)
			if !flag {
				sql_insert := `
					insert into
					` + database_pattern_local + ` (
						idpattern ,idcard, idpoin, status_pattern,  resultcardwin, 
						create_pattern, createdate_pattern 
					) values (
						$1, $2, $3, $4, $5,     
						$6, $7  
					)
				`

				flag_insert, msg_insert := Exec_SQL(sql_insert, database_pattern_local, "INSERT",
					idpattern, idcard, _Get_infomasterpointByCode(point), status, resultwin,
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
	} else {
		sql_update := `
				UPDATE 
				` + database_pattern_local + `  
				SET resultcardwin=$1, 
				update_pattern=$2, updatedate_pattern=$3       
				WHERE idpattern=$4     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_listbet_local, "UPDATE",
			resultcardwin,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
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
			COUNTidpoin    
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
func _Get_pattern_losewin(status string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total := 0
	sql_select := `SELECT
			COUNT(idpattern) total 
			FROM ` + configs.DB_tbl_trx_pattern + `  
			WHERE status_pattern='` + status + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&total); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return total
}
