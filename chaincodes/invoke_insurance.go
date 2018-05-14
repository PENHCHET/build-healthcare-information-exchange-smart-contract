package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// TODO: initialize the insurance peronsal checkup history data from MySQL Local Database
func initInsurancePersonCheckHistoryLedger(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	db, err := sql.Open("mysql", "user:password@tcp(210.115.182.218:3306)/HEALTHCARE_INFORMATION_EXCHANGE?timeout=1000s&charset=utf8")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tsql := fmt.Sprintf("SELECT " +
		"	HCHK_YEAR, " +
		"	PERSON_ID, " +
		"	YKIHO_GUBUN_CD, " +
		"	HEIGHT, " +
		"	WEIGHT, " +
		"	WAIST, " +
		"	BP_HIGH, " +
		"	BP_LWST, " +
		"	BLDS, " +
		"	TOT_CHOLE, " +
		"	TRIGLYCERIDE, " +
		"	HDL_CHOLE, " +
		"	LDL_CHOLE, " +
		"	HMG, " +
		"	OLIG_PROTE_CD, " +
		"	CREATININE, " +
		"	SGOT_AST, " +
		"	SGPT_ALT, " +
		"	GAMMA_GTP, " +
		"	HCHK_APOP_PMH_YN, " +
		"	HCHK_HDISE_PMH_YN, " +
		"	HCHK_HPRTS_PMH_YN, " +
		"	HCHK_DIABML_PMH_YN, " +
		"	HCHK_HPLPDM_PMH_YN, " +
		"	HCHK_PHSS_PMH_YN, " +
		"	HCHK_ETCDSE_PMH_YN, " +
		"	FMLY_APOP_PATIEN_YN, " +
		"	FMLY_HDISE_PATIEN_YN, " +
		"	FMLY_HPRTS_PATIEN_YN, " +
		"	FMLY_DIABML_PATIEN_YN, " +
		"	FMLY_CANCER_PATIEN_YN, " +
		"	SMK_STAT_TYPE_RSPS_CD, " +
		"	PAST_SMK_TERM_RSPS_CD, " +
		"	PAST_DSQTY_RSPS_CD, " +
		"	CUR_SMK_TERM_RSPS_CD, " +
		"	CUR_DSQTY_RSPS_CD, " +
		"	DRNK_HABIT_RSPS_CD, " +
		"	TM1_DRKQTY_RSPS_CD, " +
		"	MOV20_WEK_FREQ_ID, " +
		"	MOV30_WEK_FREQ_ID, " +
		"	WLK30_WEK_FREQ_ID, " +
		"	GLY_CD, " +
		"	OLIG_OCCU_CD, " +
		"	OLIG_PH, " +
		"	HCHK_PMH_CD1, " +
		"	HCHK_PMH_CD2, " +
		"	HCHK_PMH_CD3, " +
		"	FMLY_LIVER_DISE_PATIEN_YN, " +
		"	SMK_TERM_RSPS_CD, " +
		"	DSQTY_RSPS_CD, " +
		"	DRNK_HABIT_RSPS_CD_2008, " +
		"	TM1_DRKQTY_RSPS_CD_2008, " +
		"	EXERCI_FREQ_RSPS_CD " +
		"FROM F_GJ " +
		"WHERE BLOCKCHAIN='1'")

	// Execute query
	rows, err := db.Query(tsql)
	if err != nil {
		log.Fatal("Error reading rows: " + err.Error())
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var insurancePersonCheckupHistory InsurancePersonCheckupHistory
		// Get values from row.
		err := rows.Scan(
			&insurancePersonCheckupHistory.HCHK_YEAR,
			&insurancePersonCheckupHistory.PERSON_ID,
			&insurancePersonCheckupHistory.YKIHO_GUBUN_CD,
			&insurancePersonCheckupHistory.HEIGHT,
			&insurancePersonCheckupHistory.WEIGHT,
			&insurancePersonCheckupHistory.WAIST,
			&insurancePersonCheckupHistory.BP_HIGH,
			&insurancePersonCheckupHistory.BP_LWST,
			&insurancePersonCheckupHistory.BLDS,
			&insurancePersonCheckupHistory.TOT_CHOLE,
			&insurancePersonCheckupHistory.TRIGLYCERIDE,
			&insurancePersonCheckupHistory.HDL_CHOLE,
			&insurancePersonCheckupHistory.LDL_CHOLE,
			&insurancePersonCheckupHistory.HMG,
			&insurancePersonCheckupHistory.OLIG_PROTE_CD,
			&insurancePersonCheckupHistory.CREATININE,
			&insurancePersonCheckupHistory.SGOT_AST,
			&insurancePersonCheckupHistory.SGPT_ALT,
			&insurancePersonCheckupHistory.GAMMA_GTP,
			&insurancePersonCheckupHistory.HCHK_APOP_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_HDISE_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_HPRTS_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_DIABML_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_HPLPDM_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_PHSS_PMH_YN,
			&insurancePersonCheckupHistory.HCHK_ETCDSE_PMH_YN,
			&insurancePersonCheckupHistory.FMLY_APOP_PATIEN_YN,
			&insurancePersonCheckupHistory.FMLY_HDISE_PATIEN_YN,
			&insurancePersonCheckupHistory.FMLY_HPRTS_PATIEN_YN,
			&insurancePersonCheckupHistory.FMLY_DIABML_PATIEN_YN,
			&insurancePersonCheckupHistory.FMLY_CANCER_PATIEN_YN,
			&insurancePersonCheckupHistory.SMK_STAT_TYPE_RSPS_CD,
			&insurancePersonCheckupHistory.PAST_SMK_TERM_RSPS_CD,
			&insurancePersonCheckupHistory.PAST_DSQTY_RSPS_CD,
			&insurancePersonCheckupHistory.CUR_SMK_TERM_RSPS_CD,
			&insurancePersonCheckupHistory.CUR_DSQTY_RSPS_CD,
			&insurancePersonCheckupHistory.DRNK_HABIT_RSPS_CD,
			&insurancePersonCheckupHistory.TM1_DRKQTY_RSPS_CD,
			&insurancePersonCheckupHistory.MOV20_WEK_FREQ_ID,
			&insurancePersonCheckupHistory.MOV30_WEK_FREQ_ID,
			&insurancePersonCheckupHistory.WLK30_WEK_FREQ_ID,
			&insurancePersonCheckupHistory.GLY_CD,
			&insurancePersonCheckupHistory.OLIG_OCCU_CD,
			&insurancePersonCheckupHistory.OLIG_PH,
			&insurancePersonCheckupHistory.HCHK_PMH_CD1,
			&insurancePersonCheckupHistory.HCHK_PMH_CD2,
			&insurancePersonCheckupHistory.HCHK_PMH_CD3,
			&insurancePersonCheckupHistory.FMLY_LIVER_DISE_PATIEN_YN,
			&insurancePersonCheckupHistory.SMK_TERM_RSPS_CD,
			&insurancePersonCheckupHistory.DSQTY_RSPS_CD,
			&insurancePersonCheckupHistory.DRNK_HABIT_RSPS_CD_2008,
			&insurancePersonCheckupHistory.TM1_DRKQTY_RSPS_CD_2008,
			&insurancePersonCheckupHistory.EXERCI_FREQ_RSPS_CD,
		)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
		}

		fmt.Printf("%#v", insurancePersonCheckupHistory)

		insurancePersonCheckupHistoryAsBytes, _ := json.Marshal(insurancePersonCheckupHistory)

		// "INSURANCE_PERSON_CHECK_HISTORY"+PERSON_ID+HCHK_YEAR+YKIHO_GUBUN_CD
		key, err := APIStub.CreateCompositeKey(PREFIX_INSURANCE_PERSON_CHECK_HISTORY, []string{
			strconv.Itoa(insurancePersonCheckupHistory.PERSON_ID),
			strconv.Itoa(insurancePersonCheckupHistory.HCHK_YEAR),
			strconv.Itoa(insurancePersonCheckupHistory.YKIHO_GUBUN_CD)})
		if err != nil {
			return shim.Error(err.Error())
		}
		// APIStub.PutState("key", "value")
		APIStub.PutState(key, insurancePersonCheckupHistoryAsBytes)
		fmt.Println("Added")
	}
	return shim.Success(nil)
}

// TODO: create new personal checkup history
// params: args[0] = array of insurance peronsal checkup histories in json [{""}]
func createNewPersonCheckHistory(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== Create New Person Check History ==================================")
	if len(args) == 1 {

		var insurancePersonCheckupHistories = []InsurancePersonCheckupHistory{}

		err := json.Unmarshal([]byte(args[0]), &insurancePersonCheckupHistories)

		if err != nil {
			return shim.Error(err.Error())
		}

		for _, insurancePersonCheckupHistory := range insurancePersonCheckupHistories {
			key, err := APIStub.CreateCompositeKey(PREFIX_INSURANCE_PERSON_CHECK_HISTORY, []string{strconv.Itoa(insurancePersonCheckupHistory.PERSON_ID), strconv.Itoa(insurancePersonCheckupHistory.HCHK_YEAR), strconv.Itoa(insurancePersonCheckupHistory.YKIHO_GUBUN_CD)})
			if err != nil {
				return shim.Error(err.Error())
			}
			insurancePersonCheckupHistoryAsBytes, _ := json.Marshal(insurancePersonCheckupHistory)
			if err != nil {
				return shim.Error(err.Error())
			}
			err = APIStub.PutState(key, insurancePersonCheckupHistoryAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
			fmt.Println("Added")
		}
	}
	return shim.Success(nil)
}

// TODO: get all personal checkup histories
func getAllPersonCheckHistory(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("=============================== GetAllPersonCheck ===============================================")
	resultsIterator, err := APIStub.GetStateByPartialCompositeKey(PREFIX_INSURANCE_PERSON_CHECK_HISTORY, []string{})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonCheckupHistory := InsurancePersonCheckupHistory{}
		err = json.Unmarshal(kvResult.Value, &insurancePersonCheckupHistory)

		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, insurancePersonCheckupHistory)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

// TODO: get all peronsal checkup history by person id
func getAllPersonCheckHistoryByPersonId(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("=============================== Get All Personal Checkup History By Perons Id = %d ===============================\n", args[0])

	resultsIterator, err := APIStub.GetStateByPartialCompositeKey(PREFIX_INSURANCE_PERSON_CHECK_HISTORY, []string{args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonCheckupHistory := InsurancePersonCheckupHistory{}
		err = json.Unmarshal(kvResult.Value, &insurancePersonCheckupHistory)

		if err != nil {
			return shim.Error(err.Error())
		}
		results = append(results, insurancePersonCheckupHistory)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

// TODO: get all personal checkup history by person id and hospital id
func getAllPersonCheckHistoryByPersonIdAndCheckHospitalPermission(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("=============================== Get All Personal Checkup History By Perons Id = %d ===============================\n", args[0])

	if len(args) != 2 {
		return shim.Error("Request have to provide the hospitalId and patientId")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_REQUEST_PERMISSION, []string{args[0], args[1], "1"})

	if err != nil {
		return shim.Error(err.Error())
	}

	hospitalPatientRequestPermission := HospitalPatientRequestPermission{}

	hospitalPatientRequestPermissionAsBytes, _ := APIStub.GetState(key)

	if len(hospitalPatientRequestPermissionAsBytes) == 0 {
		return shim.Error("Could not find the hospital patient request permission yet please create a new request permission")
	}

	err = json.Unmarshal(hospitalPatientRequestPermissionAsBytes, &hospitalPatientRequestPermission)
	if err != nil {
		return shim.Error(err.Error())
	}

	if hospitalPatientRequestPermission.Status != Allowed {
		return shim.Error("The insurance information doesn't allowed")
	}

	resultsIterator, _ := APIStub.GetStateByPartialCompositeKey(PREFIX_INSURANCE_PERSON_CHECK_HISTORY, []string{args[1]})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonCheckupHistory := InsurancePersonCheckupHistory{}
		err = json.Unmarshal(kvResult.Value, &insurancePersonCheckupHistory)

		if err != nil {
			return shim.Error(err.Error())
		}
		results = append(results, insurancePersonCheckupHistory)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

// TODO: get all personal checkup history by query couchdb
func getAllPersonCheckHistoryByQueryCouchdb(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("========================== GET QUERY RESULT ================================")
	sql := "{\"selector\": {},\"pagesize\": 10}"

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", sql)

	resultsIterator, err := APIStub.GetQueryResult(sql)

	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonCheckupHistory := InsurancePersonCheckupHistory{}
		err = json.Unmarshal(kvResult.Value, &insurancePersonCheckupHistory)

		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, insurancePersonCheckupHistory)

		fmt.Printf("%#v", insurancePersonCheckupHistory)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func createNewPersonRequestPermission(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== Create New Person Request Permission From Insurance ==================================")
	if len(args) == 1 {
		request := struct {
			PersonID int    `json:"PERSON_ID"`
			Type     string `json:"TYPE"`
		}{}

		err := json.Unmarshal([]byte(args[0]), &request)

		if err != nil {
			return shim.Error(err.Error())
		}

		uuid := uuid.Must(uuid.NewV4())
		uuidStr := fmt.Sprintf("%s", uuid)

		key, err := APIStub.CreateCompositeKey(PREFIX_INSURANCE_PERSON_REQUEST_PERMISSION, []string{strconv.Itoa(request.PersonID), request.Type})
		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonRequestPermission := InsurancePersonRequestPermission{
			UUID:        uuidStr,
			PersonID:    request.PersonID,
			Description: "Request Permission To Patient " + strconv.Itoa(request.PersonID),
			Timestamp:   time.Now(),
			Type:        request.Type,
			Status:      Pending,
		}

		insurancePersonRequestPermissionAsBytes, _ := json.Marshal(insurancePersonRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = APIStub.PutState(key, insurancePersonRequestPermissionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("Added")
		return shim.Success(insurancePersonRequestPermissionAsBytes)
	}
	return shim.Error("Request Data is not found.")
}

func getPersonRequestPermission(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== getPersonRequestPermission ==================================")

	if len(args) != 2 {
		return shim.Error("Request have to provide the patientId and type")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_INSURANCE_PERSON_REQUEST_PERMISSION, []string{args[0], args[1]})

	if err != nil {
		return shim.Error(err.Error())
	}

	insurancePersonRequestPermissionAsBytes, err := APIStub.GetState(key)
	if len(insurancePersonRequestPermissionAsBytes) == 0 {
		return shim.Error("Insurance Person Request Permssion Is Not Found.")
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(insurancePersonRequestPermissionAsBytes)
}

func updatePersonRequestPermission(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== updatePersonRequestPermission ==================================")
	if len(args) == 1 {
		fmt.Println("REQUEST ==> %s", args[0])

		request := struct {
			PersonId int                     `json:"PERSON_ID"`
			Type     string                  `json:"TYPE"`
			Status   RequestPermissionStatus `json:"STATUS"`
		}{}

		err := json.Unmarshal([]byte(args[0]), &request)
		if err != nil {
			return shim.Error(err.Error())
		}

		key, err := APIStub.CreateCompositeKey(PREFIX_INSURANCE_PERSON_REQUEST_PERMISSION, []string{strconv.Itoa(request.PersonId), request.Type})
		if err != nil {
			return shim.Error(err.Error())
		}

		insurancePersonRequestPermissionAsBytes, _ := APIStub.GetState(key)

		response := struct {
			Message string `json:"MESSAGE"`
		}{Message: "Could not find the insurance person request permission"}

		responseAsBytes, _ := json.Marshal(response)

		if len(insurancePersonRequestPermissionAsBytes) == 0 {
			return shim.Success(responseAsBytes)
		}

		insurancePersonRequestPermission := InsurancePersonRequestPermission{}
		err = json.Unmarshal(insurancePersonRequestPermissionAsBytes, &insurancePersonRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		insurancePersonRequestPermission.Status = request.Status

		insurancePersonRequestPermissionAsBytes, err = json.Marshal(insurancePersonRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = APIStub.PutState(key, insurancePersonRequestPermissionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println("Updated")
		return shim.Success(insurancePersonRequestPermissionAsBytes)
	}
	return shim.Error("Request Data is not found.")
}
