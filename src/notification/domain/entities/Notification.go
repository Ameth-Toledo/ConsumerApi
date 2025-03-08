package entities

type Notification struct {
	ID         int
	Asignature string
	Message    string
}

func NewNotification(id int, asignature string) *Notification {
	return &Notification{
		ID:         id,
		Asignature: asignature,
		Message:    "La asignatura se ha registrado: " + asignature,
	}
}
