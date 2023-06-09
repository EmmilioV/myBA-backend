package common

type Event struct {
	Action string      `json:"action"`
	UserId string      `json:"user_id"`
	Old    interface{} `json:"old"`
	New    interface{} `json:"new"`
}
