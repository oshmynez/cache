It is necessary to write a separate package that will implement an in-memory cache using the following methods:

- `Set(key string, value interface{})` - record value `value` in cache by key `key`
- `Get(key string)`
- `Delete(key)`