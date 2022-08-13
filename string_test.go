// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestHexColor(t *testing.T) {
	a := assert.New(t, false)

	a.True(HexColor("#123"))
	a.True(HexColor("#fff"))
	a.True(HexColor("#f0f0f0"))
	a.True(HexColor("#fafafa"))
	a.True(HexColor("#224422"))

	a.False(HexColor("#2244"))
	a.False(HexColor("#hhh"))
	a.False(HexColor("#asdf"))
	a.False(HexColor("#ffff"))
}

func TestIP(t *testing.T) {
	a := assert.New(t, false)

	a.True(IP("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))
	a.True(IP("fe80:0:0:0:204:61ff:fe9d:f156"))
	a.True(IP("0.0.0.0"))
	a.True(IP("255.255.255.255"))
	a.True(IP("255.0.3.255"))

	a.False(IP("255.0:3.255"))
	a.False(IP("275.0.3.255"))
}

func TestIP6(t *testing.T) {
	a := assert.New(t, false)

	a.True(IP6("fe80:0000:0000:0000:0204:61ff:fe9d:f156"))      // full form of IPv6
	a.True(IP6("fe80:0:0:0:204:61ff:fe9d:f156"))                // drop leading zeroes
	a.True(IP6("fe80::204:61ff:fe9d:f156"))                     // collapse multiple zeroes to :: in the IPv6 address
	a.True(IP6("fe80:0000:0000:0000:0204:61ff:254.157.241.86")) // IPv4 dotted quad at the end
	a.True(IP6("fe80:0:0:0:0204:61ff:254.157.241.86"))          // drop leading zeroes, IPv4 dotted quad at the end
	a.True(IP6("fe80::204:61ff:254.157.241.86"))                // dotted quad at the end, multiple zeroes collapsed
	a.True(IP6("::1"))                                          // localhost
	a.True(IP6("fe80::"))                                       // link-local prefix
	a.True(IP6("2001::"))                                       // global unicast prefix
}

func TestIP4(t *testing.T) {
	a := assert.New(t, false)

	a.True(IP4("0.0.0.0"))
	a.True(IP4("255.255.255.255"))
	a.True(IP4("255.0.3.255"))
	a.True(IP4("127.10.0.1"))
	a.True(IP4("27.1.0.1"))

	a.False(IP4("1127.01.0.1"))
}

func TestEmail(t *testing.T) {
	a := assert.New(t, false)

	a.True(Email("email@email.com"))
	a.True(Email("em2il@email.com.cn"))
	a.True(Email("12345@qq.com"))
	a.True(Email("email.test@email.com"))
	a.True(Email("email.test@email123.com"))
	a.True(Email("em2il@email"))

	// 2个@
	a.False(Email("em@2l@email.com"))
	// 没有@
	a.False(Email("email2email.com.cn"))
}

func TestURL(t *testing.T) {
	a := assert.New(t, false)

	a.True(URL("http://www.example.com"))
	a.True(URL([]byte("http://example.com")))
	a.True(URL("http://www.example.com/path/?a=b"))
	a.True(URL("https://www.example.com:88/path1/path2"))
	a.True(URL("ftp://pwd:user@www.example.com/index.go?a=b"))
	a.True(URL([]byte("ftp://pwd:user@www.example.com/index.go?a=b")))
	a.True(URL("pwd:user@www.example.com/path/"))
	a.True(URL("https://[fe80:0:0:0:204:61ff:fe9d:f156]/path/"))
	a.True(URL("https://127.0.0.1/path//index.go?arg1=val1&arg2=val/2"))
	a.True(URL("https://[::1]/path/index.go?arg1=val1"))
}
