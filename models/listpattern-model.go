package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_listpattern_local = configs.DB_tbl_trx_listpattern
const database_listpatterndetail_local = configs.DB_tbl_trx_listpatterndetail

func Fetch_listpatternHome(search, status string, page int) (helpers.Responsepaging, error) {
	var obj entities.Model_listpattern
	var arraobj []entities.Model_listpattern
	var res helpers.Responsepaging
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	perpage := 25
	totalrecord := 0
	offset := page

	sql_selectcount := ""
	sql_selectcount += ""
	sql_selectcount += "SELECT "
	sql_selectcount += "COUNT(idlistpattern) as totallistpattern  "
	sql_selectcount += "FROM " + database_listpattern_local + "  "
	if search != "" {
		sql_selectcount += "WHERE LOWER(idlistpattern) LIKE '%" + strings.ToLower(search) + "%' "
	}
	if status != "" {
		sql_selectcount += "WHERE status_listpattern = '" + status + "' "
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
	sql_select += "idlistpattern , nmlistpattern, status_listpattern,  "
	sql_select += "create_listpattern, to_char(COALESCE(createdate_listpattern,now()), 'YYYY-MM-DD HH24:MI:SS'), "
	sql_select += "update_listpattern, to_char(COALESCE(updatedate_listpattern,now()), 'YYYY-MM-DD HH24:MI:SS') "
	sql_select += "FROM " + database_listpattern_local + " as A  "
	if search == "" {
		if status != "" {
			sql_select += "WHERE status_listpattern = '" + status + "' "
			sql_select += "ORDER BY createdate_listpattern DESC  LIMIT " + strconv.Itoa(perpage)
		} else {
			sql_select += "ORDER BY createdate_listpattern DESC  OFFSET " + strconv.Itoa(offset) + " LIMIT " + strconv.Itoa(perpage)
		}

	} else {
		sql_select += "WHERE LOWER(idlistpattern) LIKE '%" + strings.ToLower(search) + "%' "
		sql_select += "ORDER BY createdate_listpattern DESC  LIMIT " + strconv.Itoa(perpage)
	}
	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idlistpattern_db, nmlistpattern_db, status_listpattern_db                                          string
			create_listpattern_db, createdate_listpattern_db, update_listpattern_db, updatedate_listpattern_db string
		)

		err = row.Scan(&idlistpattern_db, &nmlistpattern_db, &status_listpattern_db,
			&create_listpattern_db, &createdate_listpattern_db, &update_listpattern_db, &updatedate_listpattern_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if create_listpattern_db != "" {
			create = create_listpattern_db + ", " + createdate_listpattern_db
		}
		if update_listpattern_db != "" {
			update = update_listpattern_db + ", " + updatedate_listpattern_db
		}
		if status_listpattern_db == "Y" {
			status_css = configs.STATUS_RUNNING
		}
		obj.Listpattern_id = idlistpattern_db
		obj.Listpattern_nmlistpattern = nmlistpattern_db
		obj.Listpattern_nmpoin = _Get_total_listpatterndetail_poin(idlistpattern_db)
		obj.Listpattern_totallose = _Get_total_listpatterndetail_losewin(idlistpattern_db, "N")
		obj.Listpattern_totalwin = _Get_total_listpatterndetail_losewin(idlistpattern_db, "Y")
		obj.Listpattern_status = status_listpattern_db
		obj.Listpattern_status_css = status_css
		obj.Listpattern_create = create
		obj.Listpattern_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Perpage = perpage
	res.Totalrecord = totalrecord
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_listpatterndetailHome(idlistpattern string) (helpers.Responselistpatterndetail, error) {
	var obj entities.Model_listpatterndetail
	var arraobj []entities.Model_listpatterndetail
	var res helpers.Responselistpatterndetail
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	total_lose := _Get_total_listpatterndetail_losewin(idlistpattern, "N")
	total_win := _Get_total_listpatterndetail_losewin(idlistpattern, "Y")

	sql_select := `SELECT 
			A.idlistpatterndetail , A.status_card, A.idpoin, 
			COALESCE(B.codepoin,'') as codepoin, COALESCE(B.nmpoin,'') as nmpoin, COALESCE(B.poin,0) as poin,   
			A.create_listpatterndetail, to_char(COALESCE(A.createdate_listpatterndetail,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.update_listpatterndetail, to_char(COALESCE(A.updatedate_listpatterndetail,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_listpatterndetail_local + ` as A   
			LEFT JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin    
			WHERE A.idlistpattern=$1 
			ORDER BY A.createdate_listpatterndetail ASC   `

	row, err := con.QueryContext(ctx, sql_select, idlistpattern)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idlistpatterndetail_db, idpoin_db, poin_db                                                                                 int
			status_card_db, codepoin_db, nmpoin_db                                                                                     string
			create_listpatterndetail_db, createdate_listpatterndetail_db, update_listpatterndetail_db, updatedate_listpatterndetail_db string
		)

		err = row.Scan(&idlistpatterndetail_db, &status_card_db, &idpoin_db, &codepoin_db, &nmpoin_db, &poin_db,
			&create_listpatterndetail_db, &createdate_listpatterndetail_db, &update_listpatterndetail_db, &updatedate_listpatterndetail_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if create_listpatterndetail_db != "" {
			create = create_listpatterndetail_db + ", " + createdate_listpatterndetail_db
		}
		if update_listpatterndetail_db != "" {
			update = update_listpatterndetail_db + ", " + updatedate_listpatterndetail_db
		}
		if status_card_db == "Y" {
			status_css = configs.STATUS_RUNNING
		}
		obj.Listpatterndetail_id = idlistpatterndetail_db
		obj.Listpatterndetail_idpoin = idpoin_db
		obj.Listpatterndetail_nmpoin = codepoin_db + "-" + nmpoin_db
		obj.Listpatterndetail_poin = poin_db
		obj.Listpatterndetail_status = status_card_db
		obj.Listpatterndetail_status_css = status_css
		obj.Listpatterndetail_create = create
		obj.Listpatterndetail_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Totallose = total_lose
	res.Totalwin = total_win
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_listpattern(admin, name, status, idrecord, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_listpattern_local, "idlistpattern", idrecord)
		if !flag {
			sql_insert := `
					insert into
					` + database_listpattern_local + ` (
						idlistpattern ,nmlistpattern, status_listpattern, 
						create_listpattern, createdate_listpattern 
					) values (
						$1, $2, $3,      
						$4, $5   
					)
				`

			flag_insert, msg_insert := Exec_SQL(sql_insert, database_listpattern_local, "INSERT",
				idrecord, name, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_listpattern_local + `  
				SET nmlistpattern=$1, status_listpattern=$2,  
				update_listpattern=$3, updatedate_listpattern=$4         
				WHERE idlistpattern=$5       
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_listpattern_local, "UPDATE",
			name, status,
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
func Save_listpatterndetail(admin, idlistpattern, status, sData string, idpoin int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
					insert into
					` + database_listpatterndetail_local + ` (
						idlistpatterndetail ,idlistpattern, status_card, idpoin,  
						create_listpatterndetail, createdate_listpatterndetail 
					) values (
						$1, $2, $3, $4,       
						$5, $6    
					)
				`
		field_column := database_listpatterndetail_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_listpatterndetail_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idlistpattern, status, idpoin,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Delete_listpatterndetail(idrecord int, tipe, idlistpatter string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()

	if tipe == "ALL" {
		sql_delete := `
				DELETE FROM
				` + database_listpatterndetail_local + ` 
				WHERE idlistpattern=$1 
			`

		flag_episode, msg_episode := Exec_SQL(sql_delete, database_listpatterndetail_local, "DELETE", idlistpatter)

		if flag_episode {
			msg = "Succes"
		} else {
			log.Println(msg_episode)
		}
	} else {
		sql_delete := `
				DELETE FROM
				` + database_listpatterndetail_local + ` 
				WHERE idlistpatterndetail=$1 
			`

		flag_episode, msg_episode := Exec_SQL(sql_delete, database_listpatterndetail_local, "DELETE", idrecord)

		if flag_episode {
			msg = "Succes"
		} else {
			log.Println(msg_episode)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _Get_total_listpatterndetail_losewin(idlistpattern, status string) int {
	con := db.CreateCon()
	ctx := context.Background()
	total := 0
	sql_select := `SELECT
			COUNT(idlistpatterndetail) total 
			FROM ` + configs.DB_tbl_trx_listpatterndetail + `  
			WHERE idlistpattern='` + idlistpattern + `'     
			AND status_card='` + status + `'     
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
func _Get_total_listpatterndetail_poin(idlistpattern string) string {
	con := db.CreateCon()
	ctx := context.Background()
	poin := ""
	total := 0

	sql_select_total := `SELECT
			COUNT(idlistpatterndetail) total 
			FROM ` + configs.DB_tbl_trx_listpatterndetail + ` as A   
			JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin    
			WHERE A.idlistpattern='` + idlistpattern + `'    
			AND A.status_card='Y'      
			AND A.idpoin>0    
		`

	row_total := con.QueryRowContext(ctx, sql_select_total)
	switch e := row_total.Scan(&total); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	sql_select := `SELECT
			B.nmpoin  
			FROM ` + configs.DB_tbl_trx_listpatterndetail + ` as A   
			JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin    
			WHERE A.idlistpattern='` + idlistpattern + `'    
			AND A.status_card='Y'      
			AND A.idpoin>0      
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	no := 0
	for row.Next() {
		no = no + 1
		var (
			nmpoin_db string
		)

		err = row.Scan(&nmpoin_db)

		helpers.ErrorCheck(err)
		if no == total {
			poin += nmpoin_db
		} else {
			poin += nmpoin_db + ", "
		}

	}
	defer row.Close()

	return poin
}
