package main

import (
	"fmt"

	"github.com/myself659/i18napi"
)

func main() {
	i18napi.AddTranLang("zh", "locate/zh.all.json")
	i18napi.AddTranLang("en", "locate/en.all.json")
	i18napi.SetDefalut("zh")
	// output: 新年快乐
	fmt.Println(i18napi.TranID("zh", "new_year"))
	// output: Happy New Year
	fmt.Println(i18napi.TranID("en", "new_year"))
	// output: 新年快乐
	fmt.Println(i18napi.TranID("en-US", "new_year"))
	// output:  3 天

	fmt.Println(i18napi.TranMap("zh", "d_days", map[string]interface{}{"Count": 3}))
	// output :3 days
	fmt.Println(i18napi.TranMap("en", "d_days", map[string]interface{}{"Count": 3}))
	// output : 30 days
	fmt.Println(i18napi.TranOne("en", "d_days", 30))
	// output: 10 days 10 hours 10 minutes 1 second
	fmt.Println(i18napi.TranMap("en", "d_h_m_s", map[string]interface{}{
		"Days":    i18napi.TranOne("en", "d_days", 10),
		"Hours":   i18napi.TranOne("en", "h_hours", 10),
		"Minutes": i18napi.TranOne("en", "m_minutes", 10),
		"Seconds": i18napi.TranOne("en", "s_seconds", 1),
	}))
	// output: 1 day 1 hour 1 minute 1 second
	fmt.Println(i18napi.TranMap("en", "d_h_m_s", map[string]interface{}{
		"Days":    i18napi.TranOne("en", "d_days", 1),
		"Hours":   i18napi.TranOne("en", "h_hours", 1),
		"Minutes": i18napi.TranOne("en", "m_minutes", 1),
		"Seconds": i18napi.TranOne("en", "s_seconds", 1),
	}))

}
