package domain

import "time"

// Product represents a product in the system.
type Product struct {
	ID    string
	Name  string
	Price float64
}

// Order represents an order in the system.
type Order struct {
	ID        string
	UserID    string
	ProductID string
	Quantity  int
}

// Message represents a chat message in a meeting
type Message struct {
	ID        string
	Content   string
	SenderID  string
	Timestamp time.Time
}

// Meeting represents a video meeting session
type Meeting struct {
	ID        string
	CreatedAt time.Time
	Users     []string   // Array of User IDs in the meeting
	Messages  []*Message // Chat messages in the meeting
	Status    MeetStatus
}

type MeetStatus string

const (
	MeetStatusPending   MeetStatus = "pending"
	MeetStatusActive    MeetStatus = "active"
	MeetStatusCompleted MeetStatus = "completed"
)

// NewProduct creates a new Product instance.
func NewProduct(id, name string, price float64) *Product {
	return &Product{ID: id, Name: name, Price: price}
}

// NewOrder creates a new Order instance.
func NewOrder(id, userID, productID string, quantity int) *Order {
	return &Order{ID: id, UserID: userID, ProductID: productID, Quantity: quantity}
}

// NewMeeting creates a new Meeting instance
func NewMeeting(id string, creatorID string) *Meeting {
	return &Meeting{
		ID:        id,
		CreatedAt: time.Now(),
		Users:     []string{creatorID},
		Messages:  make([]*Message, 0),
		Status:    MeetStatusPending,
	}
}

// NewMessage creates a new Message instance
func NewMessage(id string, content string, senderID string) *Message {
	return &Message{
		ID:        id,
		Content:   content,
		SenderID:  senderID,
		Timestamp: time.Now(),
	}
}
