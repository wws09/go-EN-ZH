package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`from=en&to=zh&query=hello&transtype=realtime&simple_means_flag=3&sign=54706.276099&token=a28f71253ec1bea6d4ab1d9b330ce8f7&domain=common&ts=1692004656367`)
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/v2transapi?from=en&to=zh", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Acs-Token", "1692004644272_1692004656387_b0Ygky8a+GI5m9kIbPi/BjmJ4giyjqjH57KxopH5Uzd/vfCgWHLDh6mal1Xr1dEAClPEfm9D8OJZP0Zw7PkDLyHDXVSyd7SIlZi/WV2lGL8ueLVaiq8z9DesPQId+TwElzfTswMLxGYUfOneP9Ge5qygTGvgDNDlnQUx9bZ9gZzoKbL/eGlqNpK5EO38AbT+Y9HixVzuLgTB/CrR81VTzqdlxmNmu2LprLhYdeMypGcMUdrvjtpmwpwOjeWOnrHTW159Y8MeoqcvaH7uWC0vjf7HLjLFouDrYcrK6noo8DaAjjdCICk2i+cjQXNLTipp6IqDdekxCWzdHVO9Qh9s/QiGui39m/V0VLKb5VFqVHFb7fw6PHBQ09zdaJ1nmGQ5aRApLMyeuB5j6PUOV5lj1/H0wcH64jkLCrLmfEDDd8otW6M896qnsgzp2ZNjoykgfOrjCAtkkJ2nk0JApPKIryHekypQvLGl/v954g5tWboMpaJyK/gHIDvOznz8G/u4")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=300D84227A88EAAB7DFF7BD2266C8E8A; PSTM=1658741149; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; REALTIME_TRANS_SWITCH=1; HISTORY_SWITCH=1; FANYI_WORD_SWITCH=1; BDUSS=g4Z05WeVBacFhxTW1tMGtycjU5dDRYdmxKalIzMzd-c0Z6OUdEQVZOMjAyaDFqSVFBQUFBJCQAAAAAAAAAAAEAAAA40GI3xr23sbuowuS-oQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALRN9mK0TfZiZl; BDUSS_BFESS=g4Z05WeVBacFhxTW1tMGtycjU5dDRYdmxKalIzMzd-c0Z6OUdEQVZOMjAyaDFqSVFBQUFBJCQAAAAAAAAAAAEAAAA40GI3xr23sbuowuS-oQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALRN9mK0TfZiZl; BAIDUID=BE546934C07F41B344C7FF111709EED1:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BA_HECTOR=802h21852h210g2k8la58g271idjq5f1p; BAIDUID_BFESS=BE546934C07F41B344C7FF111709EED1:FG=1; BDRCVFR[feWj1Vr5u3D]=mk3SLVN4HKm; delPer=0; PSINO=6; H_PS_PSSID=36561_39109_39115_39097_39038_38917_26350_39138_39137_39101; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1691897537,1691933100,1691988048,1692004646; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1692004646; ab_sr=1.0.1_ZWEwOTBjZTZkYjZjOTZjMWM4YjEyYTA1ZWRjNjFhOTFiYThhNTRmZDg1NTgxYjA5NjJmYjNiOWNlYzAxOTFlMzAzMTBjNDVlNThiYjBjYzE2MzlkZmVmYmVlMTYyMjE5OTVlZjBlMTRmNWVkYzdlMjFjZmIxOTBhODYxZjFkYjljOTlmZDAwM2I3OTdmNDhjMDE4YzdmNjc5MDQ4MTU5ZTgyYTAyZTA1ZmMxMThhZjJhYzlhYzAyZmU1NWU0ZmY4")
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
