module github.com/kiel-live/kiel-live/packages/backend

go 1.16

replace github.com/kiel-live/kiel-live/packages/pub-sub-proto => ../pub-sub-proto

require (
	github.com/gorilla/websocket v1.4.2
	github.com/kiel-live/kiel-live/packages/pub-sub-proto v2.0.0+incompatible
	github.com/pborman/uuid v1.2.1
)
