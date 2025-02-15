# Optimistic & Pessimistic Locking

## Optimistic Locking

### Definition

- A *strategy* in which we take note a version identifier e.g., date, timestamp, checksums or hashes of a record when reading it.
- When writing the record back we *check if the record's current version identifier* to make sure it is **atomic** i.e., has not been updated when we read the version and write the record back to the disk.
- If there is a mismatch i.e., some one has modified the record in the meantime, the save will be rejected and the user can restart it.

### Usage

- This strategy is mostly applicable to **high-volume systems** (multple users and processes accessing the same data) and **three-tier architecture** (clients use connection pooling instead of persisting connection) as optimistic locking *ensure most operations will not clash* and *clients do not hold database locks*.

## Pessimistic Locking

### Definition

- PL is when we *lock the record for our exclusive use* until we have finished with it.
- This has much better integrity, but we need to ensure we design our apps to avoid deadlocks

### Usage
- To use PL you need either *a direct connection to the DB* or an *externally transaction ID* that can be used independently of the connection (in case the connection is lost and the lock associated with it is orphaned). Both methods ensure that the lock is verified to be owned by a certain user during that user's session without blocking other users.

