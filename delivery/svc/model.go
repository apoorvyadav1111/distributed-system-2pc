package svc

type Agent struct {
	ID         string `json:"id"`
	IsReserved bool   `json:"is_reserved"`
	OrderID    string `json:"order_id"`
}
type ReserveAgentRequest struct {
	OrderID string `json:"order_id"`
}

type ReserveAgentResponse struct {
	AgentID string `json:"agent_id"`
}

type AssignAgentRequest struct {
	OrderID string `json:"order_id"`
}

type AssignAgentResponse struct {
	AgentID string `json:"agent_id"`
	OrderID string `json:"order_id"`
}
