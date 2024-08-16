package tests

import (
	"fmt"
	"sync"
	"testing"
	"ticket-allocating/dal"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentPurchases(t *testing.T) {
	initialTicket := &dal.Ticket{
		Name:       "Test Ticket",
		Desc:       "Test Description",
		Allocation: 5,
	}

	err := dal.CreateTicket(initialTicket).Error
	assert.NoError(t, err, "Failed to create initial ticket")

	// Create a number of concurrent purchase requests
	numRequests := 6
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Use a mutex to avoid data races when updating the failedRequests variable
	var mu sync.Mutex
	failedRequests := 0

	for i := 0; i < numRequests; i++ {
		go func(requestID int) {
			defer wg.Done()

			purchase := &dal.Purchase{
				Quantity: 1,
				UserId:   fmt.Sprintf("user-%d", requestID),
			}

			err := dal.CreatePurchaseWithTransaction(fmt.Sprintf("%d", initialTicket.ID), purchase)
			if err != nil {
				// Record failed requests
				// Locking the mutex ensures that only one goroutine can access the variable at a time
				mu.Lock()
				failedRequests++
				mu.Unlock()
			}
		}(i)
	}

	// Wait for all requests to complete
	wg.Wait()

	// Check the final state of the ticket
	ticket := &dal.Ticket{}
	err = dal.FindTicket(ticket, initialTicket.ID).Error
	assert.NoError(t, err, "Failed to find ticket")

	// Check the allocation of the ticket as expected
	expectedAllocation := initialTicket.Allocation
	assert.Equal(t, expectedAllocation, ticket.Allocation, "Allocation does not match expected value")

	// Check the number of failed requests as expected
	expectedFailedRequests := numRequests - initialTicket.Allocation
	assert.True(t, failedRequests >= expectedFailedRequests, "Number of failed requests is less than expected")
}
