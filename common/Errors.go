package common

import "errors"

var (
	ERR_LOCK_ALREDAY_REQUIRED = errors.New("锁被占用")
	ERR_NO_LOCK_IP_FOUND = errors.New("没有找到网卡IP")

)
