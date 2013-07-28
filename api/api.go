package api

//a "Load" rpc instruction
type Load struct {
	Key string
}

//a "Store" rpc instruction
type Store struct {
	Key   string
	Value string
}

type NullResult int

type ValueResult struct {
	Value string
}
