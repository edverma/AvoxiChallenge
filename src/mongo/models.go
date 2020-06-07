package mongo

type Whitelist struct {
	IP []byte `"json:ip" "bson:ip"`
}