# Consensus Context

Provides the system context for the consensus algorithm and deals with the actual content of blocks. Responsible primarily for the creation of new blocks (populated with transactions from `TransactionPool`) and for content validation of proposed blocks (verifying transaction results).

Currently, a single instance per virtual chain per node.

#### Interacts with services

* `TransactionPool` - Uses it to populate transactions in proposed blocks and validate ordering of others' proposed blocks.
* `VirtualMachine` - Uses it to execute transactions and generate receipts and state diffs.
* `StateStorage` - Queries the merkle root of the state for writing in proposed blocks and validating others' proposed ones.
* Passive towards `ConsensusAlgo`. Provides services to it upon request.

&nbsp;
## `Init` (flow)

* Initialize the configuration.

&nbsp;
## `RequestOrderingCommittee` (method)
> Returns a sorted list of nodes (public keys) that participate in the committee for the ordering of a given block height. Called by consensus algo.
> Note: currently single phase consensus is used where the committee is based on `RequestValidationCommittee`.


&nbsp;
## `RequestValidationCommittee` (method)

> Returns a sorted list of nodes (public keys) that participate in the committee for the execution validation of a given block height. Called by consensus algo.

* The current ordered list of nodes (per [block height](../../terminology.md)) is retrieved by a call to `getOrderedCommittee` in Committee contract.
* If the size of the requested committee is larger than total nodes, select all nodes as the committee.

&nbsp;
## `RequestNewTransactionsBlock` (method)

> Performed by consensus leader only, upon request from consensus algo, to perform the ordering phase of the consensus during a live round.

#### Choose pending transactions
* Get pending transactions by calling `TransactionPool.GetTransactionsForOrdering`.

#### Add trigger transaction
* If the configuration flag `TRIGGERS_ENABLED` is set to `TRUE`:
  * Generate a trigger transaction and add it as the last transaction.

#### Build Transactions block
* Current protocol version (`0x1`).
* Current virtual chain.
* Block height is given.
* Block proposer is given.
* Hash pointer to the previous (latest) Transactions block is given.
* Block timestamp.
* The merkle root hash of the transactions in the block.
* Placeholder: metadata - holds reputation / algorithm data.
* Hash of the block metadata.

<!--
#### Prepare for Results block
* Cache the Transactions block for execution (Results block).
* Optimization: Warm up by running the logic in `RequestNewResultsBlock` right now.
-->

&nbsp;
## `RequestNewResultsBlock` (method)

> Performed by the leader only, upon request from consensus algo to perform the execution phase of the consensus during a live round.

#### Execute transactions
* The Transactions block for this block height should be cached from a previous call to `RequestNewTransactionsBlock`.
* Get the block reference timestamp from the `TransactionsBlock` header.
  
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet` creating receipts and state diff.
  * Using the block reference timestamp.
  * Using the ExecutionMode set to EXECUTION_PROPOSE.


#### Build Results block
* Current protocol version (`0x1`).
* Current virtual chain.
* Block height.
* Block proposer.
* Hash pointer to the previous (latest) Results block.
* Block timestamp.
* The merkle root hash of the transaction receipts in the block.
* The hash of the state diff in the block.
* Hash pointer to the Transactions block of the same height.
* Merkle root of the state prior to the block execution, retrieved by calling `StateStorage.GetStateHash`.
  * Called with block height equals to the last committed block height (current block height - 1)

&nbsp;
## `ValidateTransactionsBlock` (method)

> Validates another node's proposed block. Performed upon a request from consensus algo when receiving a proposal during a live consensus round.

#### Check Transactions block header
* Check protocol version.
* Check virtual chain.
* Check that the header's block height matches the provided one.  
* Check that the header's previous block hash pointer matches the provided one.
* Check timestamp is within `config.CONSENSUS_CONTEXT_SYSTEM_TIMESTAMP_ALLOWED_JITTER` of current system time, and later than the previous block.
* Check the transactions merkle root matches the transactions.
* Check metadata hash.
* Check that the block proposer in the block header matches the provided one.

#### Validate trigger transaction
* If the configuration flag `TRIGGERS_ENABLED` is set to `TRUE`:
  * Validates that a trigger transaction is present as the last transaction.
  * Validate the trigger transaction content:
    * protocol_version = protocol version
    * virtual_chain_id = virtual_chain_id
    * timestamp = block.timestamp
    * signer = empty
    * contract_name = `triggers`
    * method_name = `blockPost`
    * input_argument_array = empty
    * signature = empty
* If the configuration flag `TRIGGERS_ENABLED` is set to `FALSE`, verify that no trigger transaction is present.

#### Validate transaction choice
* Call `TransactionPool.ValidateTransactionsForOrdering` to validate pre order checks, expiration and no duplication.
  * Using the provided header timestamp as a reference timestamp.
  * Remove trigger transaction (if present := if `TRIGGERS_ENABLED` is set to `TRUE`). 

&nbsp;
## `ValidateResultsBlock` (method)

> Validates another node's proposed block. Performed upon request from consensus algo when receiving a proposal during a live consensus round.

#### Check Results block header
* Check protocol version.
* Check virtual chain.
* Check that the header's block height matches the provided one. 
* Check hash pointer indeed matches the given previous block hash.
* Check timestamp equals the `TransactionsBlock` timestamp.
* Check the receipts merkle root matches the receipts.
* Check the hash of the state diff in the block.
* Check hash pointer to the Transactions block of the same height.
* Check merkle root of the state prior to the block execution, retrieved by calling `StateStorage.GetStateHash`.
  * Called with block height equals to the last committed block height (current block height - 1)
* Check that the block proposer in the block header matches the provided one.

#### Validate transaction execution
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet` creating receipts and state diff.
  * Using the provided header timestamp as a reference timestamp.
  * Using the ExecutionMode set to EXECUTION_VERIFY.
* Compare the receipts merkle root hash to the one in the block.
* Compare the state diff hash to the one in the block (supports only deterministic execution).


<!--
TODO: oded, add the diagrams again

![alt text][consensus_core_interfaces] <br/><br/>

[consensus_core_interfaces]: consensus_core_interfaces.png "Consensus - Core Interfaces"
-->
