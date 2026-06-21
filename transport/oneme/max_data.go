package oneme

import (
	"encoding/json"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

type MaxPacket struct {
	Ver     int             `json:"ver"`
	Cmd     int             `json:"cmd"`
	Opcode  int             `json:"opcode"`
	Seq     int             `json:"seq"`
	Payload json.RawMessage `json:"payload"`
}

type InternalCallerParams struct {
	ID           CallerID  `json:"id"`
	IsConcurrent bool      `json:"isConcurrent"`
	Endpoint     string    `json:"endpoint"`
	Turn         IceServer `json:"turn"`
	Stun         IceServer `json:"stun"`
}

type CallerID struct {
	Internal int64  `json:"internal"`
	External string `json:"external"`
}

type IceServer struct {
	URLs       []string `json:"urls"`
	Username   string   `json:"username"`
	Credential string   `json:"credential"`
}

type WebRTCConfig struct {
	Token             string `json:"tkn"`
	WebsocketEndpoint string `json:"wse"`
	TurnUsername      string `json:"trnu"`
}

type UserInfo struct {
	ID        int64
	FirstName string
	LastName  string
	FullName  string
	Phone     int64
}

type MaxClient struct {
	conn          *websocket.Conn
	mu            sync.Mutex
	seq           atomic.Int64
	pending       sync.Map
	deviceID      string
	loggedIn      bool
	keepaliveStop chan struct{}
	onEvent       func(MaxPacket)
}

type CallHandler struct {
	pc                *webrtc.PeerConnection
	dc                *webrtc.DataChannel
	conn              *websocket.Conn
	localID           int64
	remoteID          int64
	seq               int64
	mu                sync.Mutex
	onConnected       func()
	tag               string
	msgHandler        func(string)
	pendingCandidates []map[string]interface{}
	hasRemoteDesc     bool
	redirected        bool
	acceptSent        bool
	callAccepted      bool

	dcInbound func([]byte)
}
