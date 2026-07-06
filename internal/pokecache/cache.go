package pokecache
import ("sync"
"time")

type Cache struct{
	entry map[string]cacheEntry
	mu sync.Mutex
	duration time.Duration

}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache{
    c:= &Cache{
		entry: make(map[string]cacheEntry),
		duration:interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val  []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entry[key]
	if !ok{
		return nil,false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(){
	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	for range ticker.C{
		c.mu.Lock()

		for key,entry := range c.entry{
			if time.Since(entry.createdAt)>c.duration{
				delete(c.entry,key)
			}
		}
		c.mu.Unlock()
	}
	


}