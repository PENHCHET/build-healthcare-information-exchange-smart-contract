package main

import (
	"encoding/json"
	"strings"
	"time"

	null "gopkg.in/guregu/null.v3"
)

const (
	PREFIX_INSURANCE_PERSON_CHECK_HISTORY      = "INSURANCE_PERSON_CHECK_HISTORY"
	PREFIX_HOSPITAL_PATIENT_MEDICAL_HISTORY    = "HOSPITAL_PATIENT_MEDICAL_HISTORY"
	PREFIX_HOSPITAL_PATIENT                    = "HOSPITAL_PATIENT"
	PREFIX_HOSPITAL_PATIENT_REQUEST_PERMISSION = "HOSPITAL_PATIENT_REQUEST_PERMISSION"
	PREFIX_INSURANCE_PERSON_REQUEST_PERMISSION = "INSURANCE_PERSON_REQUEST_PERMISSION"
)

type InsurancePersonCheckupHistory struct {
	HCHK_YEAR                 int         `json:"HCHK_YEAR"`
	PERSON_ID                 int         `json:"PERSON_ID"`
	YKIHO_GUBUN_CD            int         `json:"YKIHO_GUBUN_CD"`
	HEIGHT                    null.Int    `json:"HEIGHT"`
	WEIGHT                    null.Int    `json:"WEIGHT"`
	WAIST                     null.Int    `json:"WAIST"`
	BP_HIGH                   null.Int    `json:"BP_HIGH"`
	BP_LWST                   null.Int    `json:"BP_LWST"`
	BLDS                      null.Int    `json:"BLDS"`
	TOT_CHOLE                 null.Int    `json:"TOT_CHOLE"`
	TRIGLYCERIDE              null.Int    `json:"TRIGLYCERIDE"`
	HDL_CHOLE                 null.Int    `json:"HDL_CHOLE"`
	LDL_CHOLE                 null.Int    `json:"LDL_CHOLE"`
	HMG                       null.Float  `json:"HMG"`
	OLIG_PROTE_CD             null.String `json:"OLIG_PROTE_CD"`
	CREATININE                null.Float  `json:"CREATININE"`
	SGOT_AST                  null.Int    `json:"SGOT_AST"`
	SGPT_ALT                  null.Int    `json:"SGPT_ALT"`
	GAMMA_GTP                 null.Int    `json:"GAMMA_GTP"`
	HCHK_APOP_PMH_YN          null.String `json:"HCHK_APOP_PMH_YN"`
	HCHK_HDISE_PMH_YN         null.String `json:"HCHK_HDISE_PMH_YN"`
	HCHK_HPRTS_PMH_YN         null.String `json:"HCHK_HPRTS_PMH_YN"`
	HCHK_DIABML_PMH_YN        null.String `json:"HCHK_DIABML_PMH_YN"`
	HCHK_HPLPDM_PMH_YN        null.String `json:"HCHK_HPLPDM_PMH_YN"`
	HCHK_PHSS_PMH_YN          null.String `json:"HCHK_PHSS_PMH_YN"`
	HCHK_ETCDSE_PMH_YN        null.String `json:"HCHK_ETCDSE_PMH_YN"`
	FMLY_APOP_PATIEN_YN       null.String `json:"FMLY_APOP_PATIEN_YN"`
	FMLY_HDISE_PATIEN_YN      null.String `json:"FMLY_HDISE_PATIEN_YN"`
	FMLY_HPRTS_PATIEN_YN      null.String `json:"FMLY_HPRTS_PATIEN_YN"`
	FMLY_DIABML_PATIEN_YN     null.String `json:"FMLY_DIABML_PATIEN_YN"`
	FMLY_CANCER_PATIEN_YN     null.String `json:"FMLY_CANCER_PATIEN_YN"`
	SMK_STAT_TYPE_RSPS_CD     null.String `json:"SMK_STAT_TYPE_RSPS_CD"`
	PAST_SMK_TERM_RSPS_CD     null.Int    `json:"PAST_SMK_TERM_RSPS_CD"`
	PAST_DSQTY_RSPS_CD        null.Int    `json:"PAST_DSQTY_RSPS_CD"`
	CUR_SMK_TERM_RSPS_CD      null.Int    `json:"CUR_SMK_TERM_RSPS_CD"`
	CUR_DSQTY_RSPS_CD         null.Int    `json:"CUR_DSQTY_RSPS_CD"`
	DRNK_HABIT_RSPS_CD        null.String `json:"DRNK_HABIT_RSPS_CD"`
	TM1_DRKQTY_RSPS_CD        null.Int    `json:"TM1_DRKQTY_RSPS_CD"`
	MOV20_WEK_FREQ_ID         null.String `json:"MOV20_WEK_FREQ_ID"`
	MOV30_WEK_FREQ_ID         null.String `json:"MOV30_WEK_FREQ_ID"`
	WLK30_WEK_FREQ_ID         null.String `json:"WLK30_WEK_FREQ_ID"`
	GLY_CD                    null.String `json:"GLY_CD"`
	OLIG_OCCU_CD              null.String `json:"OLIG_OCCU_CD"`
	OLIG_PH                   null.Float  `json:"OLIG_PH"`
	HCHK_PMH_CD1              null.String `json:"HCHK_PMH_CD1"`
	HCHK_PMH_CD2              null.String `json:"HCHK_PMH_CD2"`
	HCHK_PMH_CD3              null.String `json:"HCHK_PMH_CD3"`
	FMLY_LIVER_DISE_PATIEN_YN null.String `json:"FMLY_LIVER_DISE_PATIEN_YN"`
	SMK_TERM_RSPS_CD          null.String `json:"SMK_TERM_RSPS_CD"`
	DSQTY_RSPS_CD             null.String `json:"DSQTY_RSPS_CD"`
	DRNK_HABIT_RSPS_CD_2008   null.String `json:"DRNK_HABIT_RSPS_CD_2008"`
	TM1_DRKQTY_RSPS_CD_2008   null.String `json:"TM1_DRKQTY_RSPS_CD_2008"`
	EXERCI_FREQ_RSPS_CD       null.String `json:"EXERCI_FREQ_RSPS_CD"`
}

type HospitalPatientMedicalHistory struct {
	YKIHO_ID    int       `json:"YKIHO_ID"`
	PERSON_ID   int       `json:"PERSON_ID"`
	TRT_ORG_TP  int       `json:"TRT_ORG_TP"`
	KEY_SEQ     string    `json:"KEY_SEQ"`
	SEQ_NO      int       `json:"SEQ_NO"`
	RECU_FR_DT  string    `json:"RECU_FR_DT"`
	DSBJT_CD    string    `json:"DSBJT_CD"`
	SICK_SYM    string    `json:"SICK_SYM"`
	DESCRIPTION string    `json:"DESCRIPTION"`
	Timestamp   time.Time `json:"TIMESTAMP"`
}

type HospitalPatientRequestPermission struct {
	UUID        string                  `json:"UUID"`
	HospitalID  int                     `json:"HOSPITAL_ID"`
	PatientID   int                     `json:"PATIENT_ID"`
	Timestamp   time.Time               `json:"TIMESTAMP"`
	Description string                  `json:"DESCRIPTION"`
	Type        string                  `json:"TYPE"` // '1' REQUEST TO VIEW THE INSURANCE HISTORY, '2' REQUEST TO INSERT NEW EMR DATA
	Status      RequestPermissionStatus `json:"STATUS"`
}

type InsurancePersonRequestPermission struct {
	UUID        string                  `json:"UUID"`
	PersonID    int                     `json:"PERSON_ID"`
	Timestamp   time.Time               `json:"TIMESTAMP"`
	Description string                  `json:"DESCRIPTION"`
	Type        string                  `json:"TYPE"` // '1' REQUEST TO ADD THE PERSONAL CHECK HISTORY TO BLOCKCHAIN, '2' REQUEST FROM PERSON TO ADD THEIR DATA TO BLOCKCHAIN
	Status      RequestPermissionStatus `json:"STATUS"`
}

///////////////////////////////////////////////////////////
//////// CONSTANT DECLARATION
///////////////////////////////////////////////////////////
// TODO: Alias int8 to RequestPermissionStatus
type RequestPermissionStatus int8

// TODO: Create Constant Pending, Allowed, Rejected
const (
	//The request Permission is waiting
	Pending RequestPermissionStatus = iota
	// The request Permission is allowed
	Allowed
	// The request Permission is denied
	Rejected
)

func (s *RequestPermissionStatus) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	switch strings.ToUpper(value) {
	default:
		*s = Pending
	case "P":
		*s = Pending
	case "A":
		*s = Allowed
	case "R":
		*s = Rejected
	}
	return nil
}

func (s RequestPermissionStatus) MarshalJSON() ([]byte, error) {
	var value string

	switch s {
	default:
		value = "P"
	case Pending:
		value = "P"
	case Allowed:
		value = "A"
	case Rejected:
		value = "R"
	}

	return json.Marshal(value)
}
