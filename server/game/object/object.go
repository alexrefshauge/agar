package object

type GameObject interface {

	// GetId returns the id of the GameObject
	GetId() int

	// Update updates the gameobject, and returns true if any changes were made to the object, false otherwise.
	// dt (delta time) is the time in seconds, since the last update were made.
	// BUG(alex): delta time is curruently calculated using the time since the beginning of the last world update. This should be corrected to the time since individual game objects were updated.
	Update(dt float64) bool
}

type BaseGameObject struct {
	Id       int `json:"id"`
	Position Vec `json:"pos"`
}

func (o *BaseGameObject) GetId() int {
	return o.Id
}

type Player struct {
	BaseGameObject
	Name string `json:"name"`

	Velocity  Vec     `json:"vel"` // velocity as 2D vector
	Direction float32 `json:"dir"` // direction in radians
	Size      int     `json:"size"`
}

type Blob struct {
	BaseGameObject
	Size int `json:"size"`
}
