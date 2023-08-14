package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var Data map[string]string = make(map[string]string)

func query(word string, sign string) {
	client := &http.Client{}

	Data["from"] = "en"
	Data["to"] = "zh"
	Data["query"] = word
	Data["transtype"] = "realtime"
	Data["simple_means_flag"] = "3"
	Data["sign"] = sign
	Data["token"] = "a28f71253ec1bea6d4ab1d9b330ce8f7"
	Data["domain"] = "common"
	Data["ts"] = "1691901515065"

	DataUrlVal := url.Values{}
	for key, val := range Data {
		DataUrlVal.Add(key, val)
	}

	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/v2transapi?from=en&to=zh", strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Acs-Token", "1691901449301_1691901515081_UWl0Y0w1v6xoDGSqyLzp+tk2HY5wVmDaU8wf8Gw60sfAY0uubU4s785HQhsxDQNuSXUBN2xyspP9uM6wjhg5Xpp/OW8TAvLhY2oocZc7KYusjI4H7MNQP+Pc0tOp9QZchqFx1ICVtwLSM0oOOTcZ8K8ZEKs/PQuuUeX7LJLHgKTIaOERfrgJnpsIvHkOKBLdbvMytnCPTUIeo4eiWhc6nRLmhSOqE/fW1JueP6yqMVJYoMchQv3cVlgAmG1/ZwdwlTq6lnSew+o5iaL/ddbxSDteeqDoIn5tZcUj7Vj8hFCnH5f0Wvk3QY6TZ1CoSM/v5KG4AJuCj0rW6xEWpH78wCyfxrHCGK0AZN5gPKDjHmrSgN7F7ec0/l1FzMOT+gRVi/M1KSdxqjS1x3k/itbSjeuJFygK2vREDGeku1u7li8MuIuEq2lwPi5YkZUpxaRVQ4H6TcwHtZpt5MW1zU0bBnXHIOMcM2mB8zopYTZaJSDodjrbn7xiGGfP+i8o2uB5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=300D84227A88EAAB7DFF7BD2266C8E8A; PSTM=1658741149; HISTORY_SWITCH=1; REALTIME_TRANS_SWITCH=1; SOUND_PREFER_SWITCH=1; SOUND_SPD_SWITCH=1; FANYI_WORD_SWITCH=1; BDUSS=g4Z05WeVBacFhxTW1tMGtycjU5dDRYdmxKalIzMzd-c0Z6OUdEQVZOMjAyaDFqSVFBQUFBJCQAAAAAAAAAAAEAAAA40GI3xr23sbuowuS-oQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALRN9mK0TfZiZl; BDUSS_BFESS=g4Z05WeVBacFhxTW1tMGtycjU5dDRYdmxKalIzMzd-c0Z6OUdEQVZOMjAyaDFqSVFBQUFBJCQAAAAAAAAAAAEAAAA40GI3xr23sbuowuS-oQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALRN9mK0TfZiZl; BAIDUID=BE546934C07F41B344C7FF111709EED1:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BA_HECTOR=24052h0haha4252h21842l8v1idepsj1p; ZFY=1jfTQWDDNK0R3a80FZY7j0yFOv2GTfrOZl:Auxl:Bx:BC8:C; BAIDUID_BFESS=BE546934C07F41B344C7FF111709EED1:FG=1; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; PSINO=7; delPer=0; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1689411682,1690620849,1691897537; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1691897537; ab_sr=1.0.1_OTMxNzc0ZTAxNDU5YTVjZjMxMDdjYjY1MWU3Zjk2ZTJmOWIyOTM0N2JjYTE2Y2E0Y2EwOGM5N2NjMTY5ZDhmYzcxMWZhYjFmYjk4ZTAwNmVlMTg5MDBhNDQwZmYzZjI4Y2EyYTZlMmNlZmQ1M2UyOTg5NzRhOGQ4ZWUzNGJjMTRmYzBhNzJjN2Q4MWNjNTY2MWIzYmJjMWE0ZjY3OWRmMzYzYmRkODJiZWQxODUwYmQ1ZDQzZDlkYjk4OTExNTNh; BCLID=9047179323281517891; BCLID_BFESS=9047179323281517891; BDSFRCVID=t9COJeC62l10L_ofdbuihM3WbKDpfojTH6aopYqdAQwgA-wl0ALKEG0PIf8g0K4MyP7UogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; BDSFRCVID_BFESS=t9COJeC62l10L_ofdbuihM3WbKDpfojTH6aopYqdAQwgA-wl0ALKEG0PIf8g0K4MyP7UogKK3gOTH4PF_2uxOjjg8UtVJeC6EG0Ptf8g0f5; H_BDCLCKID_SF=tbKHoD8htC83eRjk2tu_2t40bMQyetJyaR3JVlOvWJ5TMCo1Dj3dBP4EbM5ta4vutavaahbM5J78ShPC-tn2WJ_OLJQfKM_LbKtj5lRc3l02VhTIe-t2yT0I2J0fL4RMW20e0h7mWIbmsxA45J7cM4IseboJLfT-0bc4KKJxbnLWeIJIjj6jK4JKja_JtjjP; H_BDCLCKID_SF_BFESS=tbKHoD8htC83eRjk2tu_2t40bMQyetJyaR3JVlOvWJ5TMCo1Dj3dBP4EbM5ta4vutavaahbM5J78ShPC-tn2WJ_OLJQfKM_LbKtj5lRc3l02VhTIe-t2yT0I2J0fL4RMW20e0h7mWIbmsxA45J7cM4IseboJLfT-0bc4KKJxbnLWeIJIjj6jK4JKja_JtjjP; H_PS_PSSID=36561_39109_39115_39097_39038_38917_26350_39138_39137_39101")
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

func main() {
	word := "present"
	sign := "612796.899725"
	query(word, sign)
}
