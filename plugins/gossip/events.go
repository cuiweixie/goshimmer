package gossip

import (
    "github.com/iotaledger/goshimmer/packages/events"
    "github.com/iotaledger/goshimmer/packages/identity"
    "github.com/iotaledger/goshimmer/packages/network"
    "github.com/iotaledger/goshimmer/packages/transaction"
)

var Events = pluginEvents{
    // neighbor events
    AddNeighbor:    events.NewEvent(peerCaller),
    UpdateNeighbor: events.NewEvent(peerCaller),
    RemoveNeighbor: events.NewEvent(peerCaller),

    // low level network events
    IncomingConnection: events.NewEvent(connectionCaller),

    // high level protocol events
    DropNeighbor:              events.NewEvent(peerCaller),
    SendTransaction:           events.NewEvent(transactionCaller),
    SendTransactionRequest:    events.NewEvent(transactionCaller), // TODO
    ReceiveTransaction:        events.NewEvent(transactionCaller),
    ReceiveTransactionRequest: events.NewEvent(transactionCaller), // TODO
    ProtocolError:             events.NewEvent(transactionCaller), // TODO

    // generic events
    Error: events.NewEvent(errorCaller),
}

type pluginEvents struct {
    // neighbor events
    AddNeighbor    *events.Event
    UpdateNeighbor *events.Event
    RemoveNeighbor *events.Event

    // low level network events
    IncomingConnection *events.Event

    // high level protocol events
    DropNeighbor              *events.Event
    SendTransaction           *events.Event
    SendTransactionRequest    *events.Event
    ReceiveTransaction        *events.Event
    ReceiveTransactionRequest *events.Event
    ProtocolError             *events.Event

    // generic events
    Error *events.Event
}

type protocolEvents struct {
    ReceiveVersion         *events.Event
    ReceiveIdentification  *events.Event
    AcceptConnection       *events.Event
    RejectConnection       *events.Event
    DropConnection         *events.Event
    ReceiveTransactionData *events.Event
    ReceiveRequestData     *events.Event
    Error                  *events.Event
}

func intCaller(handler interface{}, params ...interface{}) { handler.(func(int))(params[0].(int)) }

func identityCaller(handler interface{}, params ...interface{}) { handler.(func(*identity.Identity))(params[0].(*identity.Identity)) }

func connectionCaller(handler interface{}, params ...interface{}) { handler.(func(*network.ManagedConnection))(params[0].(*network.ManagedConnection)) }

func peerCaller(handler interface{}, params ...interface{}) { handler.(func(*Peer))(params[0].(*Peer)) }

func errorCaller(handler interface{}, params ...interface{}) { handler.(func(error))(params[0].(error)) }

func transactionCaller(handler interface{}, params ...interface{}) { handler.(func(*transaction.Transaction))(params[0].(*transaction.Transaction)) }
