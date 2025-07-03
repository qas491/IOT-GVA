let ws = null
export function connectAnnounceWS() {
  if (ws) return // 防止重复连接
  ws = new WebSocket(
    (location.protocol === 'https:' ? 'wss://' : 'ws://') + location.host + '/api/ws/announce'
  )
  ws.onmessage = function(event) {
    const msg = JSON.parse(event.data)
    if (msg.type === 'version_update') {
      if (window.$notify) {
        window.$notify({
          title: msg.title,
          message: msg.content + (msg.url ? `<a href="${msg.url}" target="_blank">查看详情</a>` : ''),
          dangerouslyUseHTMLString: true,
          type: 'info',
          duration: 0
        })
      } else {
        alert(msg.title + '\n' + msg.content)
      }
    }
  }
  ws.onclose = function() {
    ws = null
    // 可选：自动重连
    setTimeout(connectAnnounceWS, 5000)
  }
} 