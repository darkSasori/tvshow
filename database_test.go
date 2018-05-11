package main

import (
	"testing"
)

func TestNumConnections(t *testing.T) {
	conn := getConnection()

	if conn.numConnections() != 1 {
		t.Errorf("Expected 1 connection, gets %d", conn.numConnections())
	}

	conn2 := getConnection()

	if conn2.numConnections() != 2 {
		t.Errorf("Expected 2 connection, gets %d", conn2.numConnections())
	}

	conn2.disconnect()
	if conn.numConnections() != 1 {
		t.Errorf("Expected 1 connection, gets %d", conn.numConnections())
	}

	conn.disconnect()
	if conn.numConnections() != 0 {
		t.Errorf("Expected 1 connection, gets %d", conn.numConnections())
	}
}
