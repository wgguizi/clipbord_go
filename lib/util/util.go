package util

import (
	"strconv"
	"strings"
)

// 判断是否手机UA
func IsMobileUserAgent(ua string) bool {
	mobileUAs := [...]string{"iphone", "android", "phone", "mobile", "wap", "netfront", "java", "opera mobi", "opera mini",
		"ucweb", "windows ce", "symbian", "series", "webos", "sony", "blackberry", "dopod", "nokia", "samsung",
		"palmsource", "xda", "pieplus", "meizu", "midp", "cldc", "motorola", "foma", "docomo", "up.browser",
		"up.link", "blazer", "helio", "hosin", "huawei", "novarra", "coolpad", "webos", "techfaith", "palmsource",
		"alcatel", "amoi", "ktouch", "nexian", "ericsson", "philips", "sagem", "wellcom", "bunjalloo", "maui", "smartphone",
		"iemobile", "spice", "bird", "zte-", "longcos", "pantech", "gionee", "portalmmm", "jig browser", "hiptop",
		"benq", "haier", "^lct", "320x320", "240x320", "176x220"}
	ua = strings.ToLower(ua)
	for i := 0; i < len(mobileUAs); i++ {
		if strings.Contains(ua, mobileUAs[i]) {
			return true
		}
	}
	return false
}

// IP 转数值
func Ip2Int32(ipv4 string) int {
	sArr := strings.Split(ipv4, ".")
	if len(sArr) < 4 {
		return -1
	}
	i2Arr := [4]int{}
	for i := 0; i < 4; i++ {
		i2, _ := strconv.ParseUint(sArr[i], 10, 64)
		i2Arr[i] = int(i2)
	}
	r := i2Arr[3] | i2Arr[2]<<8 | i2Arr[1]<<16 | i2Arr[0]<<24
	if r > (1 << 31) {
		r = -1 * (1<<32 - r) //32位补码保存
	}
	return r
}
