package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func simulateWalletRequest(conn *websocket.Conn, clientID int) error {

	payload, err := json.Marshal(WalletPayload{ClientID: clientID})
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := WsMessage{
		Type:    MessageTypeWallet,
		Payload: payload,
	}

	err = conn.WriteJSON(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var wsResponse WsMessage
	err = json.Unmarshal(response, &wsResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if wsResponse.Type != message.Type {
		return fmt.Errorf("unexpected message type: got %v, want %v", wsResponse.Type, message.Type)
	}

	log.Printf("Client %d: Wallet request successful", clientID)
	return nil
}
func simulateEndPlayRequest(conn *websocket.Conn, clientID int) error {

	payload, err := json.Marshal(EndPlayPayload{ClientID: clientID})
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := WsMessage{
		Type:    MessageTypeEndPlay,
		Payload: payload,
	}

	err = conn.WriteJSON(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var wsResponse WsMessage
	err = json.Unmarshal(response, &wsResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if wsResponse.Type != message.Type {
		return fmt.Errorf("unexpected message type: got %v, want %v", wsResponse.Type, message.Type)
	}

	log.Printf("Client %d: End play request successful", clientID)
	return nil
}

func simulatePlayRequest(conn *websocket.Conn, clientID int) error {

	payload, err := json.Marshal(PlayPayload{ClientID: clientID, BetAmount: 10.0, BetType: Odd})
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := WsMessage{
		Type:    MessageTypePlay,
		Payload: payload,
	}

	err = conn.WriteJSON(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	_, response, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var wsResponse WsMessage
	err = json.Unmarshal(response, &wsResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if wsResponse.Type != message.Type {
		return fmt.Errorf("unexpected message type: got %v, want %v", wsResponse.Type, message.Type)
	}

	log.Printf("Client %d: Play request successful", clientID)
	return nil
}
