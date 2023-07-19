package constants

// log level
type LOG_LEVEL int

const (
	LOG_LEVEL_ERROR LOG_LEVEL = 1 + iota
	LOG_LEVEL_INFO
	LOG_LEVEL_DEBUG
)

// 예약 상태
const RESERVE_STATUS_RESERVE = 10        // 호출
const RESERVE_STATUS_RESERVE_CANCEL = 11 // 호출 취소

const RESERVE_STATUS_MATCHING = 20        // 매칭
const RESERVE_STATUS_MATCHING_CANCEL = 21 // 매칭 취소
const RESERVE_STATUS_NOT_MATCHING = 22    // 매칭 실패
const RESERVE_STATUS_HAND_OVER = 30       // 차량인계  = 유저 -> 링커 주차장 이동

const RESERVE_STATUS_DEPOSIT_KEY = 35      // 차량 인계 후 사진 업로드 시작, 메모등록, 키 번호, 키 비밀번호 부여
const RESERVE_STATUS_ADDITION_READY = 36   // 부가서비스 준비중
const RESERVE_STATUS_ADDITION_START = 37   // 부가서비스 시작
const RESERVE_STATUS_ADDITION_ONGOING = 38 // 부가서비스 진행중
const RESERVE_STATUS_ADDITION_END = 39     // 부가서비스 완료

const RESERVE_STATUS_PARKING_COMPLETION = 40   // 주차완료
const RESERVE_STATUS_KEEP_KEY = 50             // 키 보관
const RESERVE_STATUS_CALL_VEHICLE = 60         // 출차요청  = 잇차부르기 내차부르기
const RESERVE_STATUS_MATCHING_VEHICLE = 65     // 출차매칭  = 잇차링커 매칭 내차링커 매칭
const RESERVE_STATUS_MOVING = 70               // 잇차-주차장으로이동 내차-유저에게이동
const RESERVE_STATUS_VALET_DESTINATION = 75    // 목척지 도착 내차-유저에게이동
const RESERVE_STATUS_VALET_AUTH = 80           // 링커 = 잇차에게 키 수령
const RESERVE_STATUS_FAIL_PAYMENT = 82         // 결제 실패
const RESERVE_STATUS_DIRECT_VALET_PAYMENT = 85 // 직접출차 결제 완료
const RESERVE_STATUS_NO_PAYMENT = 89           // 미 결제
const RESERVE_STATUS_SERVICE_COMPLETE = 90     // 하차 또는 운행 완료를 통한 서비스 종료

// 인스타 워시 예약상태
const INSTAWASH_BOOKING_COMPLETE = 1013 // 예약 완료
const INSTAWASH_MOVING = 1020           // 세차 위치 이동중
const INSTAWASH_WASH_ONGOING = 1021     // 세차중
const INSTAWASH_WASH_COMPLETE = 1022    // 세차 완료
const INSTAWASH_BOOKING_CANCEL = 1030   // 예약 취소

// 자스민 관련 예약 정보값
const RESERVE_PARKING_INTERNATIONAL = 0 // 국제선
const RESERVE_PARKING_DOMESTIC = 1      // 국내선
const RESERVE_PARKING_OUTSIDE = 1       // 실외
const RESERVE_PARKING_INSIDE = 2        // 실내

const APP_INSTAWASH = "instawash"
const APP_TMAP = "tmap"
const APP_JASMINE = "jasmine"
