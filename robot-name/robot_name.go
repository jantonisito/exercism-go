package robotname

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ID generator
type idGen struct {
	mutex sync.Mutex
	seen  map[string]bool
}

// Generate generates a random unique ID
// Note: we use a mutex to ensure that the map is thread-safe
func (g *idGen) Generate() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	var id string
	for {
		id = ""
		for i := 0; i < 2; i++ {
			id += fmt.Sprintf("%c", rand.Intn(26)+65)
		}
		for i := 0; i < 3; i++ {
			id += fmt.Sprintf("%d", rand.Intn(10))
		}
		if _, ok := g.seen[id]; !ok {
			g.seen[id] = true
			break
		}
	}
	return id
}

// isIDUnique returns true if the ID is unique
// Note: we use a mutex to ensure that the map is thread-safe
func (g *idGen) IsIDUnique(id string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.seen[id]
}

// Singleton instance of the ID generator
var idGenerator *idGen
var syncOnce sync.Once

// getID returns a unique ID generator singleton instance
func getIdGen() *idGen {
	syncOnce.Do(func() {
		idGenerator = &idGen{
			seen: make(map[string]bool),
		}
	})
	return idGenerator
}

func init() {
	rand.Seed(time.Now().UnixNano())
	idGenerator = getIdGen()
}

// Define the Robot type here.
type Robot struct {
	name string
	seen map[string]bool
}

func (r *Robot) Name() (string, error) {
	if r.seen == nil {
		r.seen = make(map[string]bool)
	}
	if r.name == "" {
		r.name = idGenerator.Generate()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
