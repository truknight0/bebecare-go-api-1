package constants

const SUCCESS = 200
const ERR_FAIL = 400

const ERR_MSG_INVALID_PARAMETER = 9000

// 데이터베이스 관련 실패 코드
const ERR_DB_NODATA = 1000
const ERR_DB_DUPLICATION_DATA = 1001
const ERR_DB_UPDATE_DATA = 1002
const ERR_DB_DELETE_DATA = 1003
const ERR_DB_INSERT_DATA = 1004
const ERR_DB_EXIST_DATA = 1005
const ERR_DB_LOCK = 1100

// 로그인 관련 실패 코드
const ERR_LOGIN_UNAUTHORIZED_TOKEN = 2000

func GetResponseMsg(code int) string {
	//status := ""
	message := ""
	switch code {
	case SUCCESS:
		message = "api system error"
	case ERR_FAIL:
		message = "success"
	case ERR_MSG_INVALID_PARAMETER:
		message = "필수 입력항목이 누락되었습니다."
	case ERR_LOGIN_UNAUTHORIZED_TOKEN:
		message = ""
	case ERR_DB_NODATA:
		message = ""
	case ERR_DB_DUPLICATION_DATA:
		message = ""
	case ERR_DB_UPDATE_DATA:
		message = ""
	case ERR_DB_DELETE_DATA:
		message = ""
	case ERR_DB_INSERT_DATA:
		message = ""
	case ERR_DB_EXIST_DATA:
		message = ""
	default:
		message = "api system error"
	}

	return message
}
