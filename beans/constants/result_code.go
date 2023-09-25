package constants

const SUCCESS = 200
const ERR_FAIL = 400

const ERR_MSG_INVALID_PARAMETER = 9000

const ERR_PROCESS_USER_LEVEL = 9900

// 데이터베이스 관련 실패 코드
const ERR_DB_NODATA = 1000
const ERR_DB_DUPLICATION_DATA = 1001
const ERR_DB_UPDATE_DATA = 1002
const ERR_DB_DELETE_DATA = 1003
const ERR_DB_INSERT_DATA = 1004
const ERR_DB_EXIST_DATA = 1005
const ERR_DB_LOCK = 1100

// 로그인 관련 실패 코드
const ERR_LOGIN_NOTHING_TOKEN = 2000
const ERR_LOGIN_UNAUTHORIZED_TOKEN = 2001

// 초대
const ERR_INVITE_CODE = 3000
const ERR_INVITE_CODE_MAKER_JOIN = 3001

func GetResponseMsg(code int) string {
	//status := ""
	message := ""
	switch code {
	case SUCCESS:
		message = "success"
	case ERR_FAIL:
		message = "api system error"
	case ERR_MSG_INVALID_PARAMETER:
		message = "필수 입력항목이 누락되었습니다."
	case ERR_PROCESS_USER_LEVEL:
		message = "실행 권한이 없습니다."
	case ERR_LOGIN_NOTHING_TOKEN:
		message = "인증값이 존재하지 않습니다."
	case ERR_LOGIN_UNAUTHORIZED_TOKEN:
		message = "사용자 인증이 유효하지 않습니다."
	case ERR_DB_NODATA:
		message = "데이터가 존재하지 않습니다."
	case ERR_DB_DUPLICATION_DATA:
		message = "ERR_DB_DUPLICATION_DATA"
	case ERR_DB_UPDATE_DATA:
		message = "ERR_DB_UPDATE_DATA"
	case ERR_DB_DELETE_DATA:
		message = "ERR_DB_DELETE_DATA"
	case ERR_DB_INSERT_DATA:
		message = "ERR_DB_INSERT_DATA"
	case ERR_DB_EXIST_DATA:
		message = "ERR_DB_EXIST_DATA"
	case ERR_INVITE_CODE:
		message = "올바르지 않은 초대코드 입니다."
	case ERR_INVITE_CODE_MAKER_JOIN:
		message = "초대코드 생성자는 초대코드로 가입할 수 없습니다."
	default:
		message = "api system error"
	}

	return message
}
