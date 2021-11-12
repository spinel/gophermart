package memory

type MemoryDB struct {
	data map[string]int
}

func New() *MemoryDB {
	db := make(map[string]int)
	memory := &MemoryDB{
		data: db,
	}

	return memory
}

func (m *MemoryDB) Add(sessionToken string, userID int) {
	m.data[sessionToken] = userID
}

func (m *MemoryDB) Get(sessionToken string) int {
	return m.data[sessionToken]
}
