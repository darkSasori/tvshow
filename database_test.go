package main

import (
    "testing"
)

func TestNumConnections(t *testing.T) {
    conn := GetConnection()

    if conn.NumConnections() != 1 {
        t.Fatalf("Expected 1 connection, gets %d", conn.NumConnections())
    }

    conn2 := GetConnection()

    if conn2.NumConnections() != 2 {
        t.Fatalf("Expected 2 connection, gets %d", conn2.NumConnections())
    }

    conn2.Disconnect()
    if conn.NumConnections() != 1 {
        t.Fatalf("Expected 1 connection, gets %d", conn.NumConnections())
    }

    conn.Disconnect()
    if conn.NumConnections() != 0 {
        t.Fatalf("Expected 1 connection, gets %d", conn.NumConnections())
    }
}
