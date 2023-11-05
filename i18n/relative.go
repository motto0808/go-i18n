/*
 * @Author       : Motto motto@hortorgames.com
 * @Description  :
 * @Date         : 2023-06-06 18:00:16
 * @LastEditors  : Motto motto@hortorgames.com
 * @LastEditTime : 2023-06-06 20:05:00
 * @FilePath     : /go-i18n/i18n/relative.go
 */
package i18n

var timeStrMap = map[RelativeTimeFormatUnit]string{
	RelativeTimeWeek:   "week",
	RelativeTimeDay:    "day",
	RelativeTimeHour:   "hour",
	RelativeTimeMinute: "minute",
	RelativeTimeSecond: "second",
}

func (t *translator) FmtRelativeTime(value int64, unit RelativeTimeFormatUnit) string {
	key := timeStrMap[unit]
	if value == 0 {
		return t.T("this-" + key)
	}
	if value == 1 {
		s := t.T("next-" + key)
		if s != "" {
			return s
		}
	}
	if value == -1 {
		s := t.T("last-" + key)
		if s != "" {
			return s
		}
	}

	field := key + "s-num"
	key += "-relative"
	if value > 0 {
		return t.T(key, t.C(field, float64(value), 0, t.FmtNumber(float64(value), 0)))
	}
	key += "-neg"
	return t.T(key, t.C(field, float64(-value), 0, t.FmtNumber(float64(-value), 0)))
}
