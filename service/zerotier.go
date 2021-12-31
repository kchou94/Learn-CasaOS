package service

import (
	"Learn-CasaOS/pkg/config"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	httper2 "oasis/pkg/utils/httper"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

type ZeroTierService interface {
	GetToken(username, pwd string) string
}

type zerotierStruct struct {
}

var client http.Client

// 固定请求head
func GetHead() map[string]string {
	var head = make(map[string]string, 4)
	head["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	head["Accept-Language"] = "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7"
	head["Connection"] = "keep-alive"
	head["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"
	return head
}

// t 1:获取action，2：登录成功后拿session（可能需要验证有了或登录失败） 3:随机生成token 4:注册页面拿action  5:注册成功后拿验证邮箱的地址
func ZeroTierGet(url string, cookies []*http.Cookie, t uint8) (action string, c []*http.Cookie, isExistSession bool) {
	isExistSession = false
	action = ""
	c = []*http.Cookie{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	for k, v := range GetHead() {
		req.Header.Add(k, v)
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	c = resp.Cookies()
	if t == 1 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}
		action, _ = doc.Find("#kc-form-login").Attr("action")
		return
	} else if t == 2 {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "pgx-session" {
				isExistSession = true
				break
			}
		}
		// 判断是否登录成功，如果需要验证邮箱，则返回验证邮箱的地址。
		if resp.StatusCode == http.StatusFound && len(resp.Header.Get("Location")) > 0 {
			action = resp.Header.Get("Location")
		}
		return
	} else if t == 3 {
		// 返回获取到的字符串
		byteArr, _ := ioutil.ReadAll(resp.Body)
		action = string(byteArr)
	} else if t == 4 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return
		}
		action, _ = doc.Find("#kc-register-form").Attr("action")
		return
	} else if t == 5 {
		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		fmt.Println(doc.Html())
		action, _ = doc.Find("#kc-info-wrapper a").Attr("href")
		return
	}

	return
}

// 模拟提交表单
func ZeroTierPost(str bytes.Buffer, action string, cookies []*http.Cookie, isLogin bool) (url, errInfo string, err error) {
	req, err := http.NewRequest(http.MethodPost, action, strings.NewReader(str.String()))
	if err != nil {
		return "", "", errors.New("new request error")
	}
	for k, v := range GetHead() {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", "", errors.New("request error")
	}
	if !isLogin {
		// 注册成功
		if resp.StatusCode == http.StatusFound && len(resp.Header.Get("Location")) > 0 {
			return resp.Header.Get("Location"), "", nil
		} else {
			register, _ := goquery.NewDocumentFromReader(resp.Body)
			firstErr := strings.TrimSpace(register.Find("#input-error-firstname").Text())
			lastErr := strings.TrimSpace(register.Find("#input-error-lastname").Text())
			emailErr := strings.TrimSpace(register.Find("#input-error-email").Text())
			pwdErr := strings.TrimSpace(register.Find("#input-error-password").Text())
			var errD strings.Buffer
			if len(firstErr) > 0 {
				errD.WriteString(firstErr + ",")
			}
			if len(lastErr) > 0 {
				errD.WriteString(lastErr + ",")
			}
			if len(emailErr) > 0 {
				errD.WriteString(emailErr + ",")
			}
			if len(pwdErr) > 0 {
				errD.WriteString(pwdErr + ",")
			}
			return "", errD.String(), nil
		}

	} else {
		if resp.StatusCode == http.StatusFound && len(resp.Header.Get("Location")) > 0 {
			return resp.Header.Get("Location"), "", nil
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return "", "", errors.New("request error")
		}

		errDesc := doc.Find("#input-error").Text()
		if len(errDesc) > 0 {
			return "", strings.TrimSpace(errDesc), nil
		}

	}

	return "", "", nil
}

// 登录并获取 token，会出现账号密码错误，和邮箱未验证情况，目前未出现其他情况
func LoginGetToken(username, pwd string) string {
	// 拿到登录的 action
	var loginUrl = "https://accounts.zerotier.com/auth/realms/zerotier/protocol/openid-connect/auth?client_id=zt-central&redirect_uri=https%3A%2F%2Fmy.zerotier.com%2Fapi%2F_auth%2Foidc%2Fcallback&response_type=code&scope=openid+profile+email+offline_access&state=state"
	action, cookies, _ := ZeroTierGet(loginUrl, nil, 1)
	if len(action) == 0 {
		// 没有拿到action，页面结构变了
		return ""
	}
	// 登录
	var str bytes.Buffer
	str.WriteString("username=")
	str.WriteString(username)
	str.WriteString("&password=")
	str.WriteString(pwd)
	str.WriteString("&credentialID=&login=Log+In")
	url, logingErrInfo, _ := ZeroTierPost(str, action, cookies, true)

	action, cookies, isLoginOk := ZeroTierGet(url, cookies, 2)
	if !isLoginOk {
		// 登录成功，可以继续调用 api
		randomTokenUrl := "https://my.zerotier.com/api/randomtoken"
		json, _, _ := ZeroTierGet(randomTokenUrl, cookies, 3)
		// 获取一个随机 token
		token := gjson.Get(json, "token")

		userInfoUrl := "https://my.zerotier.com/api/status"
		json, _, _ = ZeroTierGet(userInfoUrl, cookies, 3)
		// 拿到用户 id
		userId := gjson.Get(json, "user.id")

		// 设置新 token
		addTokenUrl := "https://my.zerotier.com/api/user/" + userId.String() + "/token"
		data := make(map[string]string)
		rand.Seed(time.Now().UnixNano())
		data["tokenName"] = "oasis-token-" + strconv.Itoa(rand.Intn(1000))
		data["token"] = token.String()
		head := make(map[string]string)
		head["Content-Type"] = "application/json"
		_, statusCode := httper2.ZeroTierPost(addTokenUrl, data, head, cookies)
	}
}

// 登录并获取token
func (z *zerotierStruct) GetToken(username, pwd string) string {
	if len(config.ZerotierInfo.Token) > 0 {
		return config.ZerotierInfo.Token
	} else {

	}
	return ""
}

func NewZeroTierService() ZeroTierService {
	// 初始化 client
	client = http.Client{Timeout: 30 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse // 防止重定向
	},
	}
	return &zerotierStruct{}
}
