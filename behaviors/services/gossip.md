# Gossip

Connects different nodes over the network with an efficient message broadcast and unicast. This is the primary method nodes in the network use to communicate among themselves.

Inside a node: A pub/sub model where any service subscribes with Gossip to a specific topic (group of messages, all topics are described [here](../../interfaces/protocol/messages.proto)).

Currently, a single instance per virtual chain per node.

#### Interacts with services

* None - Passive. Provides services to others upon request.

&nbsp;
## `Network Requirements`

* If the number of nodes is small:
  * Fully connected, every node is connected to every other node.
  * All communication is direct.
    * Optimization: Forwarding scheme for large messages like `PRE_PREPARE`.
* Network topology is known in advance and controlled by the configuration.
  * Currently assumes all services are found on all nodes.
  * Currently assumes a single instance of a service on a node.
* No special support for retransmission except standard TCP stack.
* Recommended: Nodes should sign socket connections and authenticated them when opened (via TLS).
  * Not a hard requirement because all communication is signed at the application level.
* Reconnect to topology peers when disconnected with `config.GOSSIP_CONNECTION_KEEP_ALIVE_INTERVAL` polling interval.
* Calls `OnMessageReceived` when a gossip message is received.
* Wire format encoding is [Protobuf over TCP](../../encoding/gossip/protobuf-over-tcp.md).

&nbsp;
## `Data Structures`

#### Network address map
* Maps between node id (public key) to an active socket to the Gossip service of this node.
* Used for inter node communication.
* Assumes every node has a single gossip instance.
* Initialized based on the public gossip endpoints from node topology configuration.

#### Topic subscription table
* Map between topic to list of service ids that are subscribed to this topic.
* Used for intra node communication.
* Assumes a single instance of every service on every node.
* Generated dynamically by calls to `TopicSubscribe`.

&nbsp;
## `Init` (flow)

* Initialize the configuration.
* Connect to the gossip endpoints of all relevant peer nodes.

&nbsp;
## `TopicSubscribe` (method)

> Used by services to subscribe to a specific topic.

* Add this service id to the topic subscription table for the requested topic.

&nbsp;
## `TopicUnsubscribe` (method)

> Used by services to unsubscribe from a specific topic.

* Remove this service id from the topic subscription table for the requested topic.

&nbsp;
## `OnMessageReceived` (method)

> Called internally when a gossip message is received from another node (inter node).

#### Check message validity

* Correct block protocol version.
* Correct virtual chain.
* Check that the node is one of the message recipients.
  * According to the recipient list mode and recipient list mode.

#### Deliver the message

* Lookup the list of subscribed services for this topic in the topic subscription table.
* For each service id:
  * Rely on service topology configuration to locate the service endpoint by service id.
  * Call `Target.GossipMessageReceived` with the message content.
  * The target service is responsible for identifying the message type and processing the message accordingly.

&nbsp;
## `SendMessage` (method)

> Sends an inter node message. The message is forwarded to the services subscribed to the topic on the receiving node.

* The gossip message holds the list of destination nodes ids (public keys).
  * Setting `inverse_recipients` causes the message to be sent to all nodes in the virtual chain except the destination node ids.
  * `NULL` value for the list sends a broadcast to all nodes in the virtual chain.
* For each node id:
  * Rely on the network address map to locate the relevant socket.
  * Send the message on the socket.
