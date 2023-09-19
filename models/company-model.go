package models

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/db"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nleeper/goment"
)

const database_company_local = configs.DB_tbl_mst_company

func Fetch_companyHome() (helpers.Responsecompany, error) {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	var res helpers.Responsecompany
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcompany ,    
			to_char(COALESCE(startjoincompany,now()), 'YYYY-MM-DD HH24:MI:SS'),
			to_char(COALESCE(endjoincompany,now()), 'YYYY-MM-DD HH24:MI:SS'),
			idcurr, nmcompany, nmowner, phoneowner, emailowner, companyurl,statuscompany,
			createcompany, to_char(COALESCE(createdatecompany,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatecompany, to_char(COALESCE(updatedatecompany,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_company_local + `  
			ORDER BY createdatecompany DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcompany_db, startjoincompany_db, endjoincompany_db                                               string
			idcurr_db, nmcompany_db, nmowner_db, phoneowner_db, emailowner_db, companyurl_db, statuscompany_db string
			createcompany_db, createdatecompany_db, updatecompany_db, updatedatecompany_db                     string
		)

		err = row.Scan(&idcompany_db, &startjoincompany_db, &endjoincompany_db,
			&idcurr_db, &nmcompany_db, &nmowner_db, &phoneowner_db, &emailowner_db, &companyurl_db, &statuscompany_db,
			&createcompany_db, &createdatecompany_db, &updatecompany_db, &updatedatecompany_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcompany_db != "" {
			create = createcompany_db + ", " + createdatecompany_db
		}
		if updatecompany_db != "" {
			update = updatecompany_db + ", " + updatedatecompany_db
		}
		if statuscompany_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if startjoincompany_db == endjoincompany_db {
			endjoincompany_db = ""
		}
		obj.Company_id = idcompany_db
		obj.Company_startjoin = startjoincompany_db
		obj.Company_endjoin = endjoincompany_db
		obj.Company_idcurr = idcurr_db
		obj.Company_name = nmcompany_db
		obj.Company_nmowner = nmowner_db
		obj.Company_phoneowner = phoneowner_db
		obj.Company_emailowner = emailowner_db
		obj.Company_url = companyurl_db
		obj.Company_status = statuscompany_db
		obj.Company_status_css = status_css
		obj.Company_create = create
		obj.Company_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	sql_selectcurr := `SELECT 
			idcurr  
			FROM ` + configs.DB_tbl_mst_curr + ` 
			ORDER BY idcurr ASC    
	`
	rowcurr, errcurr := con.QueryContext(ctx, sql_selectcurr)
	helpers.ErrorCheck(errcurr)
	for rowcurr.Next() {
		var (
			idcurr_db string
		)

		errcurr = rowcurr.Scan(&idcurr_db)

		helpers.ErrorCheck(errcurr)

		objcurr.Curr_id = idcurr_db
		arraobjcurr = append(arraobjcurr, objcurr)
		msg = "Success"
	}
	defer rowcurr.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listcurr = arraobjcurr
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_company(admin, idrecord, idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_company_local, "idcompany", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_company_local + ` (
					idcompany , startjoincompany, endjoincompany,  
					idcurr , nmcompany, nmowner, phoneowner, emailowner, companyurl, statuscompany,  
					createcompany, createdatecompany 
				) values (
					$1, $2, $3,    
					$4, $5, $6, $7, $8, $9, $10, 
					$11, $12   
				)
			`
			startjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_company_local, "INSERT",
				idrecord, startjoin, startjoin,
				idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status,
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
		if status == "Y" {
			sql_update := `
				UPDATE 
				` + database_company_local + `  
				SET nmcompany=$1, nmowner=$2, phoneowner=$3, emailowner=$4, companyurl=$5, statuscompany=$6,   
				updatecompany=$7, updatedatecompany=$8      
				WHERE idcompany=$9     
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
				nmcompany, nmowner, phoneowner, emailowner, url, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		} else {
			endjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			sql_update := `
				UPDATE 
				` + database_company_local + `  
				SET endjoincompany=$1, nmcompany=$2, nmowner=$3, phoneowner=$4, emailowner=$5, companyurl=$6, statuscompany=$7,   
				updatecompany=$8, updatedatecompany=$9       
				WHERE idcompany=$10     
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_company_local, "UPDATE",
				endjoin, nmcompany, nmowner, phoneowner, emailowner, url, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		}

	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
