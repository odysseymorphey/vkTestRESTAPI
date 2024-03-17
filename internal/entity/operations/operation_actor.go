package operations

type ActorOperation struct {
	name      string
	sex       string
	birthdate string
}

func NewActorOperation(name, sex, birthdate string) *ActorOperation {
	return &ActorOperation{
		name:      name,
		sex:       sex,
		birthdate: birthdate,
	}
}
