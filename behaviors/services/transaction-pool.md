# Transaction Pool

Holds pending transactions until added to a block, and keeps committed transactions for a time window to avoid duplication. Proposes transactions for new blocks and validates proposals by other nodes.

Currently, each node has a single instance of a Transaction Pool per virtual chain per node.

#### Interacts with services

* `VirtualMachine` - Uses it to perform pre order checks on transactions (like subscription is valid).
* `PublicApi` - Notifies it about committed transactions.
* `Gossip` - Uses it to communicate with other nodes.

&nbsp;
## `Data Structures`

#### Pending transaction pool
* Holds transactions until they are added to a block and helps to prevent transaction duplication.
* Needs to support efficient query by `txhash`.
* Needs to be sorted by time to allow preparing block proposals according to policy.
* Associates every transaction with the node id (public key) of the gateway that added it to the network.
  * Must only hold a single copy of a transaction (`txhash`) regardless of its associated gateway node (the first that added it).
* No need to be persistent, pending transactions may be lost without impact.
* Max size of `config.TRANSACTION_POOL_PENDING_POOL_SIZE_IN_BYTES`.
* Interval to clear expired transactions of `config.TRANSACTION_POOL_PENDING_POOL_CLEAR_EXPIRED_INTERVAL`.
  * Transaction is expired if its timestamp is before current time minus `config.TRANSACTION_EXPIRATION_WINDOW` (eg. 30 min).
  * Notify public API about expired transactions it needs to respond to:
    * Marked as the gateway for this transaction in the pending pool if the node's public API originated it.
    * If indeed local, update the registered public api service by calling its `HandleTransactionError`.
      * Provide block height and timestamp according to the last committed block.

#### Committed transaction pool
* Holds the receipts of committed transactions to return their result and avoid transaction duplication.
  * Also holds the block height and block timestamp for every transaction.
* Needs to support efficient query by `txhash`.
* No need to be persistent, can re-sync from block storage.
* No limit on max size (depends on expiration window).
* Transactions must be stored in the pool until they are expired. 
  * A transaction is considered expired if its timestamp is before the `last_committed_block` timestamp minus `config.TRANSACTION_EXPIRATION_WINDOW`. 
    * For simplicity, note that for a committed transaction, the timestamp of the block it was included in, plus `config.TRANSACTION_POOL_FUTURE_TIMESTAMP_GRACE_TIMEOUT` may be used as an upper bound for its timestamp. Meaning, in the above calculation, replace the transaction timestamp with this sum.

#### Synchronization state
* `last_committed_block` - The last valid committed block that the transaction pool is synchronized to (persistent if the pools are).

&nbsp;
## `Init` (flow)

* Initialize the configuration.
* Load persistent data.
* If no persistent data, init `last_committed_block` to empty (symbolize the empty genesis block) and the pools to empty.
* Subscribe to gossip messages in topic `TRANSACTION_RELAY` by calling `Gossip.TopicSubscribe`.

&nbsp;
## `AddNewTransaction` (method)

> Add a new transaction from a client to the network (propagate to all pending transaction pools). Runs on the gateway node.

#### Check transaction validity
* Correct protocol version. If incorrect, return `REJECTED_UNSUPPORTED_VERSION`.
* Sender virtual chain matches the transaction pool's virtual chain. If mismatch, return `REJECTED_VIRTUAL_CHAIN_MISMATCH`.
* Check that the node is in sync:
  * Node is considered out of sync if the current time is later than the last committed block timestamp + `config.TRANSACTION_POOL_NODE_SYNC_REJECT_TIME` (eg. 2 min).
  * If out of sync, return `REJECTED_NODE_OUT_OF_SYNC`.
* Check Transaction timestamp:
  * Only accept transactions that haven't expired.
    * Transaction is expired if its timestamp is earlier than the current time minus `config.TRANSACTION_EXPIRATION_WINDOW` (eg. 30 min).
    * If expired, return `REJECTED_TIMESTAMP_WINDOW_EXCEEDED`.
  * Only accepts transactions that aren't in the future (according to the node's clock).
    * Transaction timestamp is in future if it is later than current time + `config.TRANSACTION_POOL_FUTURE_TIMESTAMP_GRACE_TIMEOUT` clock jitter grace window (e.g. 3 min).
    * If in future, return `REJECTED_TIMESTAMP_AHEAD_OF_NODE_TIME`.
* Transaction (`txhash`) doesn't already exist in the pending pool or committed pool (duplicate).
  * If already exists, return `DUPLICATE_TRANSACTION_ALREADY_PENDING` or `DUPLICATE_TRANSACTION_ALREADY_COMMITTED`.
* Verify pre order checks (like signature and subscription) by calling `VirtualMachine.TransactionSetPreOrder`.
  * Provide speculative blockHeight and timestamp (last committed blockHeight + 1 and timestamp of current time).  
  * If fails, return the returned status.
* On any failure, return the relevant error status and an empty receipt.
  * Always (even on errors) include reference block height and block timestamp in the response. 
  * For an already committed transaction, return the receipt.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if the pool is not full.
* Associate the transaction in the pool with the local node id (public key) as the gateway.

#### Batch before broadcast
* Batch until a `config.TRANSACTION_POOL_PROPAGATION_BATCH_SIZE` number of transactions is reached or until `config.TRANSACTION_POOL_PROPAGATION_BATCHING_TIMEOUT` time passes.
* Sign the batch with the local node private key.
* Broadcast the batch to all other transaction pools by calling `Gossip.SendMessage`.

&nbsp;
## `OnForwardedTransactions` (method)

> Add a transaction batch forwarded from another node to the pending transaction pool. Part of transaction propagation.

#### Check batch validity
* Check the batch signature, discard on error.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if the pool is not full.
* Associate the transaction in the pool with the batch node id (public key) as the gateway.

&nbsp;
## `GetTransactionsForOrdering` (method)

> Proposes a set of N transactions for block building based on the block building policy (currently first come first served). The block builder calls it during consensus.

#### Check synchronization status
* If a requested block height is in the future but `last_committed_block` is close to it `config.BLOCK_TRACKER_GRACE_DISTANCE` block the call until the requested height is committed. Or fail on `config.BLOCK_TRACKER_GRACE_TIMEOUT`.
* If requested block height is in the future, but `last_committed_block` is far, fail.
* If requested block height is in the past, panic.

#### Create proposal
* Propose a new block timestamp using the node's current time.  
  * If the timestamp is less or equal to the previous block's timestamp, use the previous block's timestamp + 1 nano.
* Prepare a transactions list proposal (current policy is first come first serve).
* If there are no pending transactions:
  * Wait up to `config.TRANSACTION_POOL_TIME_BETWEEN_EMPTY_BLOCKS` (eg. 5 sec).
  * If any transactions arrive before the wait is over, stop waiting and return them immediately.
  * If no transactions after the wait, continue with an empty block (number of transactions is 0).
  * If ended up waiting, update the proposed block time to reflect the time after the wait.

#### Check transactions validity
* For each transaction, check:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches the contract's virtual chain and matches the transaction pool's virtual chain.
  * The transaction was not expired.
    * Transaction is expired if its timestamp is earlier than the proposed block timestamp minus `config.TRANSACTION_EXPIRATION_WINDOW`.
  * The transaction isn't in the future (according to the proposed block timestamp).
    * Transaction timestamp is in future if it is later than the proposed block timestamp plus `config.TRANSACTION_POOL_FUTURE_TIMESTAMP_GRACE_TIMEOUT` clock jitter grace window.
  * Transaction wasn't already committed (exists in the committed pool).
  * Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.
    * Reference block provided to virtual machine is the current block, and its timestamp received as an argument.
* Transactions that failed the checks should be excluded from the result and removed from the pending pool.
  * A best effort should be made to return the requested number of transactions, meaning that if transactions were dropped, further transactions are to be requested from the pending pool.
  * If we are marked as the gateway for this transaction in the pending pool, it was originated by the node's public api.
    * If indeed local, update the registered public API service by calling its `HandleTransactionError`.
      * Provide block height and timestamp according to the last committed block.

&nbsp;
## `ValidateTransactionsForOrdering` (method)

> Verifies that an ordered list of transactions complies with the ordering policy, called by the block builder during consensus when validating a new block proposal.

#### Check synchronization status
* If requested block height is in the future, but `last_committed_block` is close to it `config.BLOCK_TRACKER_GRACE_DISTANCE`, block the call until the requested height is committed. Or fail on `config.BLOCK_TRACKER_GRACE_TIMEOUT`.
* If requested block height is in the future, but `last_committed_block` is far, fail.
* If requested block height is in the past, panic.

#### Check transactions validity
* For each transaction, check:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches contract virtual chain and matches the transaction pool's virtual chain.
  * The transaction was not expired.
    * Transaction is expired if its timestamp is earlier than the proposed block timestamp minus `config.TRANSACTION_EXPIRATION_WINDOW`.
  * The transaction isn't in the future (according to the proposed block timestamp).
    * Transaction timestamp is in future if it is later than the proposed block timestamp plus `config.TRANSACTION_POOL_FUTURE_TIMESTAMP_GRACE_TIMEOUT` clock jitter grace window.
  * Transaction wasn't already committed (exists in the committed pool).
  * Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.
    * Reference block provided to virtual machine is the current block and its timestamp received as argument.
* If one of the transaction checks fails, return an error (for all transactions).

#### Check proposal policy
* Future: Verify that the proposal matches the selection policy (current policy is first come first serve).

&nbsp;
## `CommitTransactionsReceipts` (method)

> Commits a new approved block (the transaction receipts from it is what we care about). This is the main way to update the transaction pool as new blocks are generated and approved in the system.

#### Check block height
* We assume here that the caller of this method inside the node is trusted and has already verified the block.
* We want to commit blocks in sequence, so make sure the given [block height](../../terminology.md) is the next of `last_committed_block`.
* If not, discard the commit and return the next desired block height (which is the next of `last_committed_block`).
* Optimization: If the transaction pool missed a day, it only needs to receive the last configurable expiration window (e.g. 30 min). The beginning of the day can be skipped. This will probably require a new method from block storage that translates past time stamp to block height.

#### Ignore blocks with an expired timestamp
* If the block timestamp is beyond the expiration time (and so are its transactions), it can be ignored.
  * Transaction is expired if its timestamp is earlier than the current time minus `config.TRANSACTION_EXPIRATION_WINDOW` (eg. 30 min).
  * If the pools contain any expired transactions, they have periodical timers that clear them.

#### Commit receipts
* For each transaction receipt:
  * Add the receipt, block height and block timestamp to the committed pool.
  * Notify public API about transactions it needs to respond to:
    * If we are marked as the gateway for this transaction in the pending pool, it was originated by the node's public API.
    * If indeed local, update the registered public API service by calling its `HandleTransactionsBlock`.
  * Remove the corresponding transactions (based on their `txhash`) from the pending pool.
* Update `last_committed_block` to match the given block.

&nbsp;
## `GetCommittedTransactionReceipt` (method)

> Returns the transaction receipt for a past transaction based on its id. Used when a client asks to query transaction status for an older transaction.

* If `txhash` is present in the pending transaction pool, return status `PENDING` along with the last committed block height and timestamp.
* If `txhash` is present in the committed transaction pool, return status `COMMITTED` and the receipt along with the block height and timestamp of the receipt.
* else return status `NO_RECORD_FOUND` along with the last committed block height and timestamp.

&nbsp;
## Gossip Messages Handlers

> Handles gossip messages from other nodes. Relevant messages include transaction relay messages.

#### `HandleForwardedTransactions`
* Handles `TRANSACTION_RELAY_FORWARDED_TRANSACTIONS` messages, see `send-transaction` flow.



<!-- TODO: oded, add the diagrams again

![alt text][transaction_preorder_validation_flow] <br/><br/>

[transaction_preorder_validation_flow]: transaction_preorder_validation_flow.png "PreOrder Validation Flow"

-->
