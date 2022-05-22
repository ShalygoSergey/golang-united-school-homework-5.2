package cache

import "time"

type Value struct {
	value    string
	deadline *time.Time
}

func NewValue(value string, deadline *time.Time) Value {
	return Value{
		value:    value,
		deadline: deadline,
	}
}

type Cache struct {
	m map[string]Value
}

func NewCache() Cache {
	c := Cache{}
	c.m = make(map[string]Value)
	return c
}

func (c Cache) Get(key string) (string, bool) {

	value := ""
	exist := false

	v, ok := c.m[key]

	if ok {
		if v.deadline != nil && v.deadline.Before(time.Now()) {
			delete(c.m, key)
		} else {
			value = v.value
			exist = true
		}
	}

	return value, exist
}

func (c Cache) Put(key, value string) {
	c.m[key] = NewValue(value, nil)
}

func (c Cache) Keys() []string {

	var keys []string

	now := time.Now()
	for k, v := range c.m {
		if v.deadline != nil && v.deadline.Before(now) {
			delete(c.m, k)
		} else {
			keys = append(keys, k)
		}
	}

	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = NewValue(value, &deadline)
}
