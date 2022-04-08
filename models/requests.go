package models

type AuthenticateRequest struct {
	Login       string
	Password    string
	GoogleToken string
	TelegramID  string
}

type AuthenticateResponse struct {
	Account Account
	Token   string
}

type CreateEventRequest struct {
	Date        string
	Title       string
	Duration    string
	Link        string
	Org         string
	Target      string
	Where       string
	Description string
	Amount      string
	Place       string
}

type CreateEventResponse struct {
	Event Event
}

type GetEventRequest struct {
	EventId string
}

type GetEventResponse struct {
	Event Event
}

type DeleteEventRequest struct {
	EventId string
}

type DeleteEventResponse struct {
	Event Event
}

type UpdateEventRequest struct {
	EventId     string
	Date        string
	Title       string
	Duration    string
	Link        string
	Org         string
	Target      string
	Where       string
	Description string
	Amount      string
	Place       string
}

type UpdateEventResponse struct {
	Event Event
}

type ListEventsRequest struct {
	Date string
}

type ListEventsResponse struct {
	Events []Event
}
