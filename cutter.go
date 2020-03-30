package cookiecutter

import (
	"net/http"
	"strings"
)

type CookieCutter struct {
	Req *http.Request
}

func (c *CookieCutter) Cookie(name string) (*http.Cookie, error) {
	return c.Req.Cookie(name)
}

func From(r *http.Request) *CookieCutter {
	cc := CookieCutter{Req: r}
	return &cc
}

type SliceCookie struct {
	Cookie *http.Cookie
	Values []string
}

type MapCookie struct {
	Cookie *http.Cookie
	Values map[string]string
}

func (cc *CookieCutter) MapCookie(name string, valueSeparator string, keypairSeparator string) (MapCookie, error) {
	c, err := cc.Cookie(name)
	if err != nil {
		return MapCookie{}, err
	}

	pairs := strings.Split(c.Value, valueSeparator)
	mCookie := MapCookie{
		Cookie: c,
	}
	for _, pair := range pairs {
		kp := strings.Split(pair, keypairSeparator)
		mCookie.Values[kp[0]] = kp[1]
	}

	return mCookie, nil
}

func (cc *CookieCutter) SliceCookie(name string, separator string) (SliceCookie, error) {
	c, err := cc.Cookie(name)
	if err != nil {
		return SliceCookie{}, err
	}

	values := strings.Split(c.Value, separator)
	sCookie := SliceCookie{
		Cookie: c,
	}
	for _, value := range values {
		sCookie.Values = append(sCookie.Values, value)
	}

	return sCookie, nil
}
