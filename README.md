# GoChat
Golang chat rest server
# Endpoints
## POST /chat/
Create new chat

**Request:**
```
{
    "title": "justtitle"
}
```
**Response:**
```
{
    "id": "1"
}
```
## GET /chat/
Return chat list

**Response:**
```
[
    {
        "id": "1",
        "title": "justtitle1",
        "created_at": "2013-02-01T12:24:54.168Z"
    },
    {
        "id": "2",
        "title": "justtitle2",
        "created_at": "2013-02-01T12:25:12.543Z"
    }
]
```
## GET /chat/{id}
Return chat info

**Response:**
```
{
    "id": "1",
    "title": "justtitle1",
    "created_at": "2013-02-01T12:24:54.168Z"
}
```
## PUT /chat/{id}
Update chat

**Request:**
```
{
    "title": "anothertitle"
}
```
**Response:**
```
{
    "status": "ok"
}
```
## DELETE /chat/{id}
Delete chat

**Response:**
```
{
    "status": "ok"
}
```
## POST /message
Broadcast message to all websocket clients

**Request:**
```
{
    "username": "username",
    "chat_id": "1",
    "message": "Test message"
}
```
**Response:**
```
{
    "status": "ok"
}
```
**Websocket response:**
```
{
    "username": "username",
    "chat_id": "1",
    "message": "Test message"
}
```
# Run
To run local project:
`go run main.go`
