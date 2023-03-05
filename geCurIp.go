package go_utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/hktalent/htmlquery"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var skpMac = regexp.MustCompile(`docker|lo|utun|gif|stf|awd`)

// 获取当前 mac 地址 hex 格式，可以作为 51pwn.com 的前缀
func GetActiveMac() string {
	ifc, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var a []string
	for _, i := range ifc {
		if 0 < len(skpMac.FindAllString(i.Name, -1)) {
			continue
		}
		macAddr := strings.TrimSpace(hex.EncodeToString(i.HardwareAddr))
		// interface down; loopback interface
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 || macAddr == "" {
			continue
		}

		addrs, _ := i.Addrs()
		bHb := false
		for _, addr := range addrs {
			if bHb {
				break
			}
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if !ip.IsPrivate() {
					continue
				}
				a = append(a, macAddr)
				bHb = true
				fmt.Println(macAddr, ip, addr.String(), addr.Network(), i.Flags.String())
				break
			}
		}
	}
	if 0 < len(a) {
		return strings.Join(a, ",")
	} else {
		m1 := GetIp()
		if nil != m1 {
			szIp := fmt.Sprintf("%v", (*m1)["ip"])
			return Pack32BinaryIP4(szIp)
		}
		return ""
	}
}
func IP4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}

// Pack32BinaryIP4("127.0.0.1") 7f000001
// Pack32BinaryIP4("192.168.0.1") c0a80001
func Pack32BinaryIP4(ip4Address string) string {
	ipv4Decimal := IP4toInt(net.ParseIP(ip4Address))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(ipv4Decimal))

	if err != nil {
		fmt.Println("Unable to write to buffer:", err)
	}

	// present in hexadecimal format
	result := fmt.Sprintf("%x", buf.Bytes())
	return result
}

var PubIp *map[string]interface{}

// 当前ip,自动跳过socks proxy
// X-Limit: current_qps=1; limit_qps=50; current_pv=10197; limit_pv=1000000
func GetIp() *map[string]interface{} {
	if nil != PubIp && 0 < len(*PubIp) {
		return PubIp
	}
	szUrl := "https://apis.map.qq.com/ws/location/v1/ip"
	c := GetClient(szUrl)
	c.UseHttp2 = false
	var m1 map[string]interface{}
	p11 := c.GetClient(nil)
	c.DoGetWithClient4SetHd(p11, szUrl, "POST", strings.NewReader("key="+url.QueryEscape("IVOBZ-QNW6P-SUKDY-LFQSE-LUFCJ-3CFUE")+"&sig=afebe5ad5227ec75a1f3d8b97f888cda"), func(r *http.Response, err1 error, szU string) {
		defer r.Body.Close()
		if data, err := ioutil.ReadAll(r.Body); nil == err {
			log.Println(string(data))
			if nil == Json.Unmarshal(data, &m1) {
				log.Printf("%+v", m1)
			}
		}
	}, func() map[string]string {
		return map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 6.1; rv:45.0) Gecko/20100101 Firefox/45.0", "Accept": "*/*"}
	}, false)
	if m2, ok := m1["result"]; ok {
		m1 = m2.(map[string]interface{})
	} else { // 失败，用其他方法获取
		if doc, err := htmlquery.LoadURL("https://iplocation.com"); nil == err {
			m1 = map[string]interface{}{}
			if node, err := htmlquery.Query(doc, "body > div.top-container > div.bottom-container > div.rubber-container.result > div > table > tbody > tr:nth-child(1) > td > b"); nil == err {
				m1["ip"] = node.Data
			}
			m2 := map[string]interface{}{}
			m1["location"] = m2
			if node, err := htmlquery.Query(doc, "body > div.top-container > div.bottom-container > div.rubber-container.result > div > table > tbody > tr:nth-child(2) > td"); nil == err {
				m2["lat"] = node.Data
			}
			if node, err := htmlquery.Query(doc, "body > div.top-container > div.bottom-container > div.rubber-container.result > div > table > tbody > tr:nth-child(3) > td"); nil == err {
				m2["lng"] = node.Data
			}
			m3 := map[string]interface{}{}
			m1["ad_info"] = m3
			if node, err := htmlquery.Query(doc, "body > div.top-container > div.bottom-container > div.rubber-container.result > div > table > tbody > tr:nth-child(6) > td"); nil == err {
				m3["city"] = node.Data
			}
			if node, err := htmlquery.Query(doc, "body > div.top-container > div.bottom-container > div.rubber-container.result > div > table > tbody > tr:nth-child(4) > td > span"); nil == err {
				m3["country"] = node.Data
			}
		}
	}
	PubIp = &m1
	return PubIp
}
