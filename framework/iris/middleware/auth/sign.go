package auth

import (
	"encoding/base64"
	"encoding/json"
	"gitee.com/itsos/golibs/v2/cerrors"
	"gitee.com/itsos/golibs/v2/config"
	"gitee.com/itsos/golibs/v2/framework/iris/bootstrap"
	"gitee.com/itsos/golibs/v2/framework/iris/caches"
	"gitee.com/itsos/golibs/v2/utils/crypt"
	"gitee.com/itsos/golibs/v2/utils/crypt/rsa"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"sort"
	"strconv"
	"strings"
	"time"
)

// CheckSign 验证签名
func CheckSign(b *bootstrap.Bootstrapper) {
	b.UseGlobal(func(ctx iris.Context) {
		if ctx.GetStatusCode() < iris.StatusInternalServerError && ctx.Method() != "OPTIONS" {
			var mJson []byte
			contentType := ctx.GetContentTypeRequested()
			if contentType == context.ContentFormMultipartHeaderValue ||
				contentType == context.ContentFormHeaderValue {
				mJson, _ = json.Marshal(ctx.FormValues())
			} else {
				body, _ := ctx.GetBody()
				uParams := ctx.URLParams()
				params := map[string]interface{}{}
				if len(body) > 0 {
					err := json.Unmarshal(body, &params)
					if err != nil {
						golog.Error(err)
						return
					}
				}
				for k, v := range uParams {
					params[k] = v
				}
				mJson, _ = json.Marshal(params)
			}
			data := filter(mJson)

			token := ctx.GetHeader("token")
			sign := ctx.GetHeader("sign")
			ts := ctx.GetHeader("ts")
			nonce := ctx.GetHeader("nonce")

			// 300s 五分钟内有效
			tss, _ := strconv.Atoi(ts)
			if time.Now().Unix() > (int64(tss) + 300) {
				forbidden(ctx)
				golog.Errorf("sign fail: time expired.[token:%s]", token)
				return
			}

			dPlain, _ := base64.StdEncoding.DecodeString(nonce)
			priv := []byte(config.Config.GetCryptRsaPriv())
			key, err := rsa.Decrypt(dPlain, priv)
			if err != nil {
				forbidden(ctx)
				golog.Errorf("sign fail: nonce decode fail.[token:%s]", token)
				return
			}

			if crypt.Md5(token+ts+string(key)+data+nonce) != sign {
				forbidden(ctx)
				golog.Errorf("sign fail: not equals.[token:%s]", token)
				return
			}

			if !caches.SignSet(sign) {
				forbidden(ctx)
				golog.Errorf("sign is exists.[token:%s]", token)
				return
			}

			ctx.Next()
		}
	})
}

func forbidden(ctx iris.Context) {
	ctx.StopWithError(iris.StatusForbidden, cerrors.Error("forbidden_access"))
}

func filter(str []byte) string {
	mString := string(str)
	mString = strings.ReplaceAll(mString, "{", "")
	mString = strings.ReplaceAll(mString, "}", "")
	mString = strings.ReplaceAll(mString, "[", "")
	mString = strings.ReplaceAll(mString, "]", "")
	mString = strings.ReplaceAll(mString, ",", "")
	mString = strings.ReplaceAll(mString, "\"", "")
	mString = strings.ReplaceAll(mString, ":", "")
	sString := strings.SplitN(mString, "", len(mString))
	sort.Strings(sString)
	return strings.Join(sString, "")
}
