package go_utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/hktalent/PipelineHttp"
	"github.com/pion/webrtc/v3"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	E2ePath = "/e2eRelay"
	// Allows compressing offer/answer to bypass terminal input limits.
	compress = true
)

var pipE = PipelineHttp.NewPipelineHttp()

// 发送通讯信号
func SignalCandidate(addr string, c *webrtc.ICECandidate) error {
	payload := []byte(c.ToJSON().Candidate)
	pipE.DoGetWithClient4SetHd(
		nil,
		fmt.Sprintf("https://%s%s", addr, E2ePath),
		"POST",
		bytes.NewReader(payload),
		func(resp *http.Response, err error, szU string) {

		}, func() map[string]string {
			return map[string]string{"Content-Type": "application/json; charset=utf-8"}
		}, true)

	return nil
}

// key 标识不同用户，对等的p2p
func GetPeerConnection(key string, certificates *[]webrtc.Certificate) *webrtc.PeerConnection {
	config := webrtc.Configuration{
		PeerIdentity: key,
		Certificates: *certificates,
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println(err)
		return nil
	}
	return peerConnection
}

// Encode encodes the input in base64
// It can optionally zip the input before encoding
func Encode(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	if compress {
		b = zip(b)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Decode decodes the input from base64
// It can optionally unzip the input after decoding
func Decode(in string, obj interface{}) {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		panic(err)
	}

	if compress {
		b = unzip(b)
	}

	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}

func zip(in []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write(in)
	if err != nil {
		panic(err)
	}
	err = gz.Flush()
	if err != nil {
		panic(err)
	}
	err = gz.Close()
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func unzip(in []byte) []byte {
	var b bytes.Buffer
	_, err := b.Write(in)
	if err != nil {
		panic(err)
	}
	r, err := gzip.NewReader(&b)
	if err != nil {
		panic(err)
	}
	res, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return res
}
