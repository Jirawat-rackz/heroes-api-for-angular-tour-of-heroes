package heroes

type Hero struct {
	ID   int    `bson:"hero_id" json:"id"`
	Name string `json:"name"`
}
