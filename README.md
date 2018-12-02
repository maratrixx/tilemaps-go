# tilemaps-go

基于 Golang 协程开发实现的瓦片地图下载服务

## Requirements

NULL

## Installation

```
    go get github.com/Liyafeng/tilemaps-go
```

## Documentation

TODO

## Example
``` go
// Fetch new store.
store, err := NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
if err != nil {
	panic(err)
}
defer store.Close()

// Get a session.
session, err = store.Get(req, "session-key")
if err != nil {
	log.Error(err.Error())
}

// Add a value.
session.Values["foo"] = "bar"

// Save.
if err = sessions.Save(req, rsp); err != nil {
	t.Fatalf("Error saving session: %v", err)
}

// Delete session.
session.Options.MaxAge = -1
if err = sessions.Save(req, rsp); err != nil {
	t.Fatalf("Error saving session: %v", err)
}

// Change session storage configuration for MaxAge = 10 days.
store.SetMaxAge(10 * 24 * 3600)
```
