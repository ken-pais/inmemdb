# inmemdb
/*
In Memory Database/Cache
Features:
	- CRUD Operations 
		SET
		GET
		DEL
	Eviction Policies: Implements a FIFO eviction policy. You can extend it to add more policies like LRU, LFU, etc.
	TTL Support: Cache entries can expire based on a TTL (Time-to-Live).
	Clear Cache: Clears all data from the cache, resetting both the Data map and the Queue.
	Exposed APIs for testing purposes.
Design Patterns:
	Singleton Pattern: Ensures a single instance of the cache throughout the application lifecycle.
	Factory Pattern: Used for creating eviction policy struct by policy type
	Strategy Pattern: Used for implementing eviction policy struct by policy type
Data structures:
	Queue:
		- To keep track of the order of items inserted into DB
		+-----------+     +-----------+     +-----------+
		|    key1   | --> |    key2   | --> |    key3   |
		+-----------+     +-----------+     +-----------+
	Hashmap:
		+-----------+     +-----------+     +-----------+     +-----------+
		|   key1    | --> |   key2    | --> |   key3    | --> |   key4    |
		|   value1  |     |   value2  |     |   value3  |     |   value4  |
		+-----------+     +-----------+     +-----------+     +-----------+
*/