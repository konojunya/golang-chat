package main

type Hub struct {
  // 登録されたクライアント
  clients map[*Client]bool

  // クライアントからくるメッセージ
  broadcast chan []byte

  // クライアントからのリクエストを登録する
  register chan *Client

  // クライアントからのリクエストの登録を解除する
  unregister chan *Client
}

func newHub() *Hub {
  // Hubのポインタのアドレスを返す
  return &Hub{
    broadcast: make(chan []byte),
    register: make(chan *Client),
    unregister: make(chan *Client),
    clients: make(map[*Client]bool),
  }
}

func (h *Hub) run() {
  for{
    select {
      // 登録されたらclientに代入
      // Hubのclients配列にクライアント = trueを代入
    case client := <-h.register:
      h.clients[client] = true

      // 登録を解除したらclientに代入
      // クライアントがあったら削除の処理
    case client := <-h.unregister:
      if _, ok := h.clients[client]; ok {
        delete(h.clients,client)
        close(client.send)
      }

      // ブロードキャストをmessageに代入
    case message := <-h.broadcast:
      // クライアントの数分ループ
      for client := range h.clients {
        select {
          // メッセージをそのクライアントに送る
        case client.send <- message:
        default:
          close(client.send)
          delete(h.clients,client)
        }
      }
    }
  }
}