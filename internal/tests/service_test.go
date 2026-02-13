package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"ig4llc.com/internal/domain/events"
)

// MockRepository is a mock type for the events.Repository interface
type MockRepository struct {
	mock.Mock
}

// Implement the interface methods using the mock helper
func (m *MockRepository) Create(e *events.Event) error {
	args := m.Called(e)
	return args.Error(0)
}

// Stubs for the rest of the interface - much cleaner!
func (m *MockRepository) GetAll() ([]events.Event, error) {
	args := m.Called(eventID, userID)
	return args.Error(0)
}
func (m *MockRepository) GetByID(id int64) (*events.Event, error) {
	args := m.Called(eventID, userID)
	return args.Error(0)
}
func (m *MockRepository) Update(e *events.Event) error {
	args := m.Called(eventID, userID)
	return args.Error(0)
}
func (m *MockRepository) Delete(id int64) error {
	args := m.Called(eventID, userID)
	return args.Error(0)
}
func (m *MockRepository) Register(eventID int64, userID int) {
	args := m.Called(eventID, userID)
	return args.Error(0)
}
func (m *MockRepository) Unregister(eventID int64, userID int) {
	args := m.Called(eventID, userID)
	return args.Error(0)
}

func TestCreateEvent_Validation(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	svc := events.NewService(mockRepo)

	// Test Case: Empty Name
	err := svc.CreateEvent(&events.Event{Name: ""})

	// Assertions
	assert.ErrorIs(t, err, events.ErrInvalidEvent)
	mockRepo.AssertNotCalled(t, "Create", mock.Anything) // Verify repo wasn't even touched
}

// simulate a successful database save
func TestCreateEvent_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	svc := events.NewService(mockRepo)

	// Create a valid event (Matches logic in service.go)
	validEvent := &events.Event{
		Name:     "Gala Dinner",
		Location: "Gotham Hall",
		DateTime: time.Now().Add(24 * time.Hour), // Not a zero-value
		UserID:   1,
	}

	// Tell the mock: "When Create is called with this exact event, return nil (no error)"
	mockRepo.On("Create", validEvent).Return(nil)

	// Execute
	err := svc.CreateEvent(validEvent)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Verifies that Create() was actually called as expected
}

func TestCreateEvent_DatabaseError(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	svc := events.NewService(mockRepo)

	// A valid event that will pass Service-layer validation
	validEvent := &events.Event{
		Name:     "Charity Auction",
		Location: "Community Center",
		DateTime: time.Now().Add(48 * time.Hour),
		UserID:   2,
	}

	// Define a specific database error
	dbError := errors.New("database connection timeout")

	// Tell the mock: Expect Create() to be called, but return our error
	mockRepo.On("Create", validEvent).Return(dbError)

	// Execute
	err := svc.CreateEvent(validEvent)

	// Assertions
	assert.Error(t, err)                          // Ensure an error was returned
	assert.Equal(t, dbError.Error(), err.Error()) // Verify it's the specific DB error
	mockRepo.AssertExpectations(t)                // Confirm the repo was actually called
}
