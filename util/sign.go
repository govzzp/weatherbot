package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strconv"
	"strings"
)

func Sign(appKey, appSecret, method, path, nonce string, timestamp int64, query map[string]string) string {
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var queryStr strings.Builder
	for i, k := range keys {
		if i > 0 {
			queryStr.WriteString("&")
		}
		queryStr.WriteString(k)
		queryStr.WriteString("=")
		queryStr.WriteString(query[k])
	}
	stringToSign := strings.Join([]string{
		method, path, queryStr.String(), appKey, nonce, strconv.FormatInt(timestamp, 10),
	}, ":")
	h := hmac.New(sha256.New, []byte(appSecret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
