package system

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketHub struct {
	clients map[*websocket.Conn]bool
	lock    sync.Mutex
}

var wsHub = &WebSocketHub{
	clients: make(map[*websocket.Conn]bool),
}

func (hub *WebSocketHub) Broadcast(message []byte) {
	hub.lock.Lock()
	defer hub.lock.Unlock()
	for client := range hub.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			client.Close()
			delete(hub.clients, client)
		}
	}
}

func (hub *WebSocketHub) AddClient(conn *websocket.Conn) {
	hub.lock.Lock()
	defer hub.lock.Unlock()
	hub.clients[conn] = true
}

func (hub *WebSocketHub) RemoveClient(conn *websocket.Conn) {
	hub.lock.Lock()
	defer hub.lock.Unlock()
	delete(hub.clients, conn)
	conn.Close()
}

// WebSocket连接接口
func (EQApi *EquipmentApi) WebSocketAnnounce(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	wsHub.AddClient(conn)
	defer wsHub.RemoveClient(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// 这里只做推送，不处理客户端消息
	}
}

// 推送版本更新公告
func PushVersionUpdate(title, content, version, url string) {
	type AnnounceMsg struct {
		Type    string `json:"type"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Version string `json:"version"`
		Url     string `json:"url,omitempty"`
	}
	msg := AnnounceMsg{
		Type:    "version_update",
		Title:   title,
		Content: content,
		Version: version,
		Url:     url,
	}
	b, _ := json.Marshal(msg)
	wsHub.Broadcast(b)
}

// 推送版本更新公告的API
// @Tags Announce
// @Summary 推送版本更新公告
// @Accept application/json
// @Produce application/json
// @Param data body struct{Title string `json:"title"`; Content string `json:"content"`; Version string `json:"version"`; Url string `json:"url"`} true "公告内容"
// @Success 200 {object} response.Response{msg=string} "推送成功"
// @Router /announce/pushVersionUpdate [post]
func (EQApi *EquipmentApi) PushVersionUpdateAPI(c *gin.Context) {
	type req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Version string `json:"version"`
		Url     string `json:"url"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}
	PushVersionUpdate(r.Title, r.Content, r.Version, r.Url)
	response.OkWithMessage("推送成功", c)
}
