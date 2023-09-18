package protocols

//go:generate easytags $GOFILE db:snake,json:snake

type BaseRequest struct {
	AppVer    string `json:"appVer" db:"app_ver"`
	AppOs     string `json:"appOs" db:"app_os"`
	AdminId   string `json:"adminId" db:"admin_id"`
	UserId    string `json:"userId" db:"user_id"`
	DriverId  string `json:"driverId" db:"driver_id"`
	SessionId string `json:"sessionId" db:"session_id"`
	ZoneIdx   int    `json:"zoneIdx" db:"zone_idx"`
	AppType   string `json:"appType" db:"app_type"`
}

type SaveAffiliatedRequest struct {
	CardList []CardInfo `json:"card_list" db:"card_list"`
}

type CardInfo struct {
	/* 일련번호 */
	Uid string `json:"uid" db:"uid"`
	/* 제휴카드 일련번호 */
	CardAffiliatedUid string `json:"card_affiliated_uid" db:"card_affiliated_uid"`
	/*BIN 번호*/
	BinNum string `json:"bin_num" db:"bin_num"`
	/*브랜드*/
	Brand string `json:"brand" db:"brand"`
	/*카드사명*/
	Company string `json:"company" db:"company"`
	/*카드명*/
	Name string `json:"name" db:"name"`
	/*사용 여부*/
	OnService string `json:"on_service" db:"on_service"`
	/*이용 가능 횟수*/
	UsableCount string `json:"usable_count" db:"usable_count"`
	/*이용권 안내 문구*/
	VoucherIntro string `json:"voucher_intro" db:"voucher_intro"`
	/*카드 이미지(大) 주소*/
	CardImgBigUrl string `json:"card_img_big_url" db:"card_img_big_url"`
	/*카드 이미지(小) 주소*/
	CardImgSmallUrl string `json:"card_img_small_url" db:"card_img_small_url"`
	/*카드 등록자 수*/
	CardUserCount string `json:"card_user_count" db:"card_user_count"`
	/*쿠폰(이용권) 발급자 수*/
	CouponIssueCount string `json:"coupon_issue_count" db:"coupon_issue_count"`
	/*카드 입력 상태 (1: 추가, 2: 갱신)*/
	Status string `db:"status" json:"status"`
	/*등록일시*/
	RegDate string `json:"reg_date" db:"reg_date"`
	/*수정일시*/
	UpdateDate string `json:"update_date" db:"update_date"`
}
