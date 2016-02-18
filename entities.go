package main

type entities struct {
	pool []autonomous
}

func NewEntities() *entities {
	n := new(entities)
	n.pool = make([]autonomous, 0)
	return n
}

/* Find the entity position inside the entities slice
 */
func (e entities) findEntityPosition(entity autonomous) int {
	for i := range e.pool {
		if entity == e.pool[i] {
			return i;
		}
	}
	return -1
}

func (e *entities) add(entity autonomous) {
	e.pool = append(e.pool, entity)
}

func (e *entities) remove(entity autonomous) {
	position := e.findEntityPosition(entity)
	e.pool[position] = nil
	e.pool = append(e.pool[:position], e.pool[position+1:]...)
}

func (e *entities) each(callback func(autonomous)) {
	for i := range e.pool {
		callback(e.pool[i])
	}
}
