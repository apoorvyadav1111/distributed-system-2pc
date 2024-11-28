package svc

import (
	"errors"

	"github.com/apoorvyadav1111/distributed-systems-2pc/delivery/io"
)

func ReserveAgent() (*Agent, error) {
	txn, err := io.DB.Begin()
	row := txn.QueryRow(`
	SELECT 
		id, is_reserved, order_id 
	FROM agents 
	WHERE 
		is_reserved is False and order_id is Null
	LIMIT 1 
	FOR UPDATE
	`)

	if row == nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err = row.Scan(&agent.ID, &agent.IsReserved, &agent.OrderID)

	if err != nil {
		txn.Rollback()
		return nil, errors.New("No agent available")
	}

	_, err = txn.Exec(`
	UPDATE agents
	SET is_reserved = True
	WHERE id = ?
	`, agent.ID)

	if err != nil {
		txn.Rollback()
		return nil, err
	}

	return &agent, nil

}

func AssignAgent(orderID string) (*Agent, error) {
	txn, err := io.DB.Begin()
	row := txn.QueryRow(`
	SELECT 
		id, is_reserved, order_id 
	FROM agents 
	WHERE 
		is_reserved is True and order_id is Null
	LIMIT 1 
	FOR UPDATE
	`)

	if row == nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err = row.Scan(&agent.ID, &agent.IsReserved, &agent.OrderID)

	if err != nil {
		txn.Rollback()
		return nil, errors.New("No agent available")
	}

	_, err = txn.Exec(`
	UPDATE agents
	SET is_reserved = false, order_id = ?
	WHERE id = ?
	`, orderID, row.ID)

	if err != nil {
		txn.Rollback()
		return nil, err
	}

	return &agent, nil
}
