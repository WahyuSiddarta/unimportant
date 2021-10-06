package helper

import (
	"database/sql"
	"encoding/json"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx/types"
)

// StringToInt :
func StringToInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

//TimeToString : null time to string
func TimeToString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

//TimeToTimestamp : null time to timestamp
func TimeToTimestamp(t time.Time) int64 {
	return t.UTC().Unix()
}

//NullTimeToString : null time to string
func NullTimeToString(t sql.NullTime) string {
	if t.Valid {
		return t.Time.UTC().Format(time.RFC3339)
	}
	return ""
}

//NullTimeToStringMySQL : null time to string MySQL
func NullTimeToStringMySQL(t mysql.NullTime) string {
	if t.Valid {
		return t.Time.UTC().Format(time.RFC3339)
	}
	return ""
}

//NullTimeToTimestamp : null time to timestamp
func NullTimeToTimestamp(t sql.NullTime) int64 {
	if t.Valid {
		return t.Time.UTC().Unix()
	}
	return 0
}

//RemoveWhiteSpace :
func RemoveWhiteSpace(t string) string {
	return strings.Replace(t, " ", "", -1)
}

//NullStringToString :
func NullStringToString(t sql.NullString) string {
	if t.Valid {
		return t.String
	}
	return ""
}

//NullIntToInt :
func NullIntToInt(t sql.NullInt32) int32 {
	if t.Valid {
		return t.Int32
	}
	return 0
}

// NullInt64ToInt64
func NullInt64ToInt64(t sql.NullInt64) int64 {
	if t.Valid {
		return t.Int64
	}
	return 0
}

//NullFloat64ToFloat64 :
func NullFloat64ToFloat64(t sql.NullFloat64) float64 {
	if t.Valid {
		return t.Float64
	}
	return 0
}

//NullJSONTextToJSONText :
func NullJSONTextToJSONText(t types.NullJSONText) types.JSONText {
	if t.Valid {
		return t.JSONText
	}
	return []byte{}
}

//NullStringToStringNoWhiteSpace :
func NullStringToStringNoWhiteSpace(t sql.NullString) string {
	if t.Valid {
		return RemoveWhiteSpace(t.String)
	}
	return ""
}

//SplitToArray : Split string to array by delimiter
func SplitToArray(str, delim string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, delim)
}

//EmptyStringToString :
func EmptyStringToString(t *string) string {
	value := ""
	if t != nil {
		value = *t
	} else {
		value = ""
	}
	return value
}

//DateFormatter :
func DateFormatter(dateString string) (time.Time, error) {
	return time.Parse("01/02/2006", dateString)
}

// CheckZero :
func CheckZero(val int64) (int64, int) {
	zero := 0
	for val >= 10000 {
		val = val / 1000
		zero = zero + 3
	}
	return val, zero
}

// CheckZeroRoundUp :
func CheckZeroRoundUp(val int64) (float64, int) {
	f := float64(val)
	zero := 0
	for f >= 1000 {
		f = f / 1000
		zero = zero + 3
	}
	divider := 0.0
	if zero > 0 {
		divider = math.Pow10((zero))
		f = (float64(val) / divider)

	} else {
		f = float64(val)
	}

	return f, zero
}

// IsPhoneNumberFormat :
func IsPhoneNumberFormat(phoneNumber string) (bool, error) {
	isFormat, err := regexp.MatchString(`^([1-9|\+]{1})([0-9]{1,4})[.]([1-9]{1})([0-9]{6,13})$`, phoneNumber)
	return isFormat, err
}

//IsCountryCodeFormat :
func IsCountryCodeFormat(countryCode string) (bool, error) {
	isFormat, err := regexp.MatchString(`^[A-Z]{3}$`, countryCode)
	return isFormat, err
}

//IsRegencyIDFormat :
func IsRegencyIDFormat(regencyID string) (bool, error) {
	isFormat, err := regexp.MatchString(`^[A-Z]{3}\-[A-Z]{2}\-[0-9]{2}$`, regencyID)
	return isFormat, err
}

// CheckInt :
func CheckInt(number string) bool {
	if _, err := strconv.ParseInt(number, 10, 64); err == nil {
		return true
	}
	return false
}

// IsJSON :
func IsJSON(str string) bool {
	var j map[string]interface{}
	k := json.Unmarshal([]byte(str), &j)
	return k == nil
}

// IsJSONString :
func IsJSONString(str string) bool {
	var j string
	return json.Unmarshal([]byte(str), &j) == nil
}

//Int64ToThousand :
func Int64ToThousand(nominal int64, separator string) string {

	n := strconv.FormatInt(nominal, 10)
	if len(n) < 4 {
		return n
	}

	var res string
	i := 0
	for j := len(n) - 1; j >= 0; j-- {
		i++
		res = string(n[j]) + res
		if i > 1 && i%3 == 0 {
			res = separator + res
		}
	}
	return strings.Trim(res, separator)
}

func StringToArrayString(a []string) string {
	return strings.Join(a, ",")
}
