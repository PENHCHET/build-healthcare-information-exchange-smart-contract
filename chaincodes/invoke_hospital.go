package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	uuid "github.com/satori/go.uuid"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: initialize the hospital patient medical history data from MySQL Local Database
func initHospitalPatientMedicalHistory(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("======================== Initalize the hospital patient medical history data from MYSQL Local Database =========================")
	db, err := sql.Open("mysql", "user:password@tcp(210.115.182.218:3306)/HEALTHCARE_INFORMATION_EXCHANGE?timeout=1000s&charset=utf8")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tsql := fmt.Sprintf("SELECT YKIHO_ID, PERSON_ID, TRT_ORG_TP, KEY_SEQ, SEQ_NO, RECU_FR_DT, DSBJT_CD, SICK_SYM, SUBSTRING_INDEX(NAME,'.',-1) AS DESCRIPTION FROM HOSPITAL_T40_T20 INNER JOIN LOD_D_Disease ON SICK_SYM = CODE WHERE BLOCKCHAIN = '1'")

	// Execute query
	rows, err := db.Query(tsql)
	if err != nil {
		log.Fatal("Error reading rows: " + err.Error())
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var hospitalPatientMedicalHistory HospitalPatientMedicalHistory
		// Get values from row.
		err := rows.Scan(
			&hospitalPatientMedicalHistory.YKIHO_ID,
			&hospitalPatientMedicalHistory.PERSON_ID,
			&hospitalPatientMedicalHistory.TRT_ORG_TP,
			&hospitalPatientMedicalHistory.KEY_SEQ,
			&hospitalPatientMedicalHistory.SEQ_NO,
			&hospitalPatientMedicalHistory.RECU_FR_DT,
			&hospitalPatientMedicalHistory.DSBJT_CD,
			&hospitalPatientMedicalHistory.SICK_SYM,
			&hospitalPatientMedicalHistory.DESCRIPTION)

		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
		}

		hospitalPatientMedicalHistory.Timestamp = time.Now()

		fmt.Printf("%#v", hospitalPatientMedicalHistory)

		hospitalPatientMedicalHistoryAsBytes, _ := json.Marshal(hospitalPatientMedicalHistory)

		key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY, []string{
			strconv.Itoa(hospitalPatientMedicalHistory.YKIHO_ID),
			strconv.Itoa(hospitalPatientMedicalHistory.PERSON_ID),
			strconv.Itoa(hospitalPatientMedicalHistory.TRT_ORG_TP),
			hospitalPatientMedicalHistory.KEY_SEQ,
			strconv.Itoa(hospitalPatientMedicalHistory.SEQ_NO),
		})
		if err != nil {
			return shim.Error(err.Error())
		}
		APIStub.PutState(key, hospitalPatientMedicalHistoryAsBytes)
		fmt.Println("Added")
	}

	return shim.Success(nil)
}

// TODO: get all patients by hospital id
func getAllPatientByHospitalID(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

// TODO: create new patient request permission
// params: args[0]= {"HOSPITAL_ID": 1, "PATIENT_ID": 10168602}

func createNewPatientRequestPermission(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== Create New Patient Request Permission From Hospital ==================================")
	if len(args) == 1 {
		request := struct {
			HospitalID int    `json:"HOSPITAL_ID"`
			PatientID  int    `json:"PATIENT_ID"`
			Type       string `json:"TYPE"`
		}{}

		err := json.Unmarshal([]byte(args[0]), &request)

		if err != nil {
			return shim.Error(err.Error())
		}

		uuid := uuid.Must(uuid.NewV4())
		uuidStr := fmt.Sprintf("%s", uuid)

		key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_REQUEST_PERMISSION, []string{strconv.Itoa(request.HospitalID), strconv.Itoa(request.PatientID), request.Type})
		if err != nil {
			return shim.Error(err.Error())
		}

		hospitalPatientRequestPermission := HospitalPatientRequestPermission{
			UUID:        uuidStr,
			HospitalID:  request.HospitalID,
			PatientID:   request.PatientID,
			Description: "Request Permission From Hospital " + strconv.Itoa(request.HospitalID) + " To Patient " + strconv.Itoa(request.PatientID),
			Timestamp:   time.Now(),
			Type:        request.Type,
			Status:      Pending,
		}

		hospitalPatientRequestPermissionAsBytes, _ := json.Marshal(hospitalPatientRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = APIStub.PutState(key, hospitalPatientRequestPermissionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("Added")
		return shim.Success(hospitalPatientRequestPermissionAsBytes)
	}
	return shim.Error("Request Data is not found.")
}

func getPatientRequestPermissionByHospitalIdAndPatientId(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== getPatientRequestPermissionByHospitalIdAndPatientId ==================================")

	if len(args) != 3 {
		return shim.Error("Request have to provide the hospitalId ,patientId and type")
	}

	key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_REQUEST_PERMISSION, []string{args[0], args[1], args[2]})

	if err != nil {
		return shim.Error(err.Error())
	}

	hospitalPatientRequestPermissionAsBytes, _ := APIStub.GetState(key)

	if len(hospitalPatientRequestPermissionAsBytes) == 0 {
		return shim.Success(nil)
	}
	return shim.Success(hospitalPatientRequestPermissionAsBytes)
}

func updatePatientRequestPermissionByHospitalIdAndPatientId(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== updatePatientRequestPermissionByHospitalIdAndPatientId ==================================")
	if len(args) == 1 {
		fmt.Println("REQUEST ==> %s", args[0])

		request := struct {
			HospitalID int                     `json:"HOSPITAL_ID"`
			PatientID  int                     `json:"PATIENT_ID"`
			Type       string                  `json:"TYPE"`
			Status     RequestPermissionStatus `json:"STATUS"`
		}{}

		err := json.Unmarshal([]byte(args[0]), &request)
		if err != nil {
			return shim.Error(err.Error())
		}

		key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_REQUEST_PERMISSION, []string{strconv.Itoa(request.HospitalID), strconv.Itoa(request.PatientID), request.Type})
		if err != nil {
			return shim.Error(err.Error())
		}

		hospitalPatientRequestPermissionAsBytes, _ := APIStub.GetState(key)

		response := struct {
			Message string `json:"MESSAGE"`
		}{Message: "Could not find the hospital patient request permission"}

		responseAsBytes, _ := json.Marshal(response)

		if len(hospitalPatientRequestPermissionAsBytes) == 0 {
			return shim.Success(responseAsBytes)
		}

		hospitalPatientRequestPermission := HospitalPatientRequestPermission{}
		err = json.Unmarshal(hospitalPatientRequestPermissionAsBytes, &hospitalPatientRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		hospitalPatientRequestPermission.Status = request.Status

		hospitalPatientRequestPermissionAsBytes, err = json.Marshal(hospitalPatientRequestPermission)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = APIStub.PutState(key, hospitalPatientRequestPermissionAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		fmt.Println("Updated")
		return shim.Success(hospitalPatientRequestPermissionAsBytes)
	}
	return shim.Error("Request Data is not found.")
}

func getAllPatientMedicalHistory(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("=============================== GetAllPatientMedicalHistory ===============================================")

	resultsIterator, err := APIStub.GetStateByPartialCompositeKey(PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY, []string{})

	if len(args) == 2 {
		resultsIterator, err = APIStub.GetStateByPartialCompositeKey(PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY, []string{args[0], args[1]})
	}

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

		hospitalPatientMedicalHistory := HospitalPatientMedicalHistory{}
		err = json.Unmarshal(kvResult.Value, &hospitalPatientMedicalHistory)

		if err != nil {
			return shim.Error(err.Error())
		}

		results = append(results, hospitalPatientMedicalHistory)
		fmt.Printf("%#v", hospitalPatientMedicalHistory)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

func addNewPatientMedicalHistory(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== Create New Patient Medical History ==================================")
	if len(args) == 1 {

		hospitalPatientMedicalHistory := HospitalPatientMedicalHistory{}

		err := json.Unmarshal([]byte(args[0]), &hospitalPatientMedicalHistory)

		if err != nil {
			return shim.Error(err.Error())
		}

		key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY, []string{strconv.Itoa(hospitalPatientMedicalHistory.YKIHO_ID), strconv.Itoa(hospitalPatientMedicalHistory.PERSON_ID)})
		if err != nil {
			return shim.Error(err.Error())
		}
		hospitalPatientMedicalHistory.Timestamp = time.Now()
		hospitalPatientMedicalHistoryAsBytes, _ := json.Marshal(hospitalPatientMedicalHistory)
		if err != nil {
			return shim.Error(err.Error())
		}
		err = APIStub.PutState(key, hospitalPatientMedicalHistoryAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("Added")
	}
	return shim.Success(nil)
}

// TODO: create new personal checkup history
// params: args[0] = array of insurance peronsal checkup histories in json [{""}]
func createNewPatientMedicalHistoryFromLocalDBByPersonId(APIStub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("============================== Create New Person Check History ==================================")
	if len(args) == 1 {

		var hospitalPatientMedicalHistories = []HospitalPatientMedicalHistory{}

		err := json.Unmarshal([]byte(args[0]), &hospitalPatientMedicalHistories)

		if err != nil {
			return shim.Error(err.Error())
		}

		for _, hospitalPatientMedicalHistory := range hospitalPatientMedicalHistories {
			key, err := APIStub.CreateCompositeKey(PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY, []string{
				strconv.Itoa(hospitalPatientMedicalHistory.YKIHO_ID),
				strconv.Itoa(hospitalPatientMedicalHistory.PERSON_ID),
				strconv.Itoa(hospitalPatientMedicalHistory.TRT_ORG_TP),
				hospitalPatientMedicalHistory.KEY_SEQ,
				strconv.Itoa(hospitalPatientMedicalHistory.SEQ_NO),
			})
			if err != nil {
				return shim.Error(err.Error())
			}
			hospitalPatientMedicalHistoryAsBytes, _ := json.Marshal(hospitalPatientMedicalHistory)
			if err != nil {
				return shim.Error(err.Error())
			}
			err = APIStub.PutState(key, hospitalPatientMedicalHistoryAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
			fmt.Println("Added")
		}
	}
	return shim.Success(nil)
}
