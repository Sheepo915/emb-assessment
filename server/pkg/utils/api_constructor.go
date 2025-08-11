package utils

import "net/url"

type QueryKVPair struct {
	Key   string
	Value string
}

func APIConstruct(baseUrl string, kv ...*QueryKVPair) (string, error) {
	l, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	if len(kv) > 0 {
		q := l.Query()
		for _, v := range kv {
			q.Add(v.Key, v.Value)
		}
		l.RawQuery = q.Encode()
	}

	return l.String(), nil
}
