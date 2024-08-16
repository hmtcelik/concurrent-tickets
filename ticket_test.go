package main

import (
	"testing"
	"ticket-allocating/dal"

	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	ticket := &dal.Ticket{
		Name:       "Test Ticket",
		Desc:       "Test Description",
		Allocation: 100,
	}

	err := dal.CreateTicket(ticket).Error
	assert.NoError(t, err, "Failed to create ticket")
}

func TestGetTicket(t *testing.T) {
	newTicket := &dal.Ticket{
		Name:       "Test Ticket",
		Desc:       "Test Description",
		Allocation: 100,
	}

	err := dal.CreateTicket(newTicket).Error
	assert.NoError(t, err, "Failed to create ticket")

	ticket := new(dal.Ticket)
	err = dal.FindTicket(ticket, ticket.ID).Error
	assert.NoError(t, err, "Failed to find ticket")

	assert.Equal(t, newTicket.ID, ticket.ID, "Ticket ID mismatch")
	assert.Equal(t, newTicket.Name, ticket.Name, "Ticket Name mismatch")
	assert.Equal(t, newTicket.Desc, ticket.Desc, "Ticket Desc mismatch")
	assert.Equal(t, newTicket.Allocation, ticket.Allocation, "Ticket Allocation mismatch")
}
