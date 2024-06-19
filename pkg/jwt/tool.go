package jwt

import "time"

func GetAccessExpTime(t time.Time) (AccessExpTime time.Time, AccessExpString string) {
	layout := "2006/01/02 15:04:05"
	AccessExpTime = t.Truncate(time.Second).Add(AccessTokenExpiredDuration)
	return AccessExpTime, AccessExpTime.Format(layout)
}
