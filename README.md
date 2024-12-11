In-Memory Database/Cache

Features:
CRUD Operations:
SET
GET
DEL
Eviction Policies:
Implements a FIFO (First In, First Out) eviction policy.
You can extend it to add more policies like LRU (Least Recently Used), LFU (Least Frequently Used), etc.
TTL Support:
Cache entries can expire based on a TTL (Time-to-Live).
Clear Cache:
Clears all data from the cache, resetting both the data map and the queue.
Exposed APIs:
Exposed APIs for testing purposes.
Design Patterns:
Singleton Pattern:
Ensures a single instance of the cache throughout the application's lifecycle.

Factory Pattern:
Used for creating eviction policy structs based on policy type.

Strategy Pattern:
Used for implementing eviction policy logic based on policy type.

Data Structures:
Queue:

Keeps track of the order of items inserted into the database.
plaintext
Copy code
+-----------+     +-----------+     +-----------+
|    key1   | --> |    key2   | --> |    key3   |
+-----------+     +-----------+     +-----------+
Hashmap:

Stores the actual data with key-value pairs.
plaintext
Copy code
+-----------+     +-----------+     +-----------+     +-----------+
|   key1    | --> |   key2    | --> |   key3    | --> |   key4    |
|   value1  |     |   value2  |     |   value3  |     |   value4  |
+-----------+     +-----------+     +-----------+     +-----------+
