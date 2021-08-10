package remotemetrics

import (
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"

	"github.com/iotaledger/hive.go/events"
)

// CollectionLogEvents defines the events for the remotelogmetrics package.
type CollectionLogEvents struct {
	// TangleTimeSyncChanged defines the local sync status change event based on tangle time.
	TangleTimeSyncChanged *events.Event
}

// SyncStatusChangedEventCaller is called when a node changes its sync status.
func SyncStatusChangedEventCaller(handler interface{}, params ...interface{}) {
	handler.(func(SyncStatusChangedEvent))(params[0].(SyncStatusChangedEvent))
}

// SyncStatusChangedEvent is triggered by a node when its sync status changes. It is also structure that is sent to remote logger.
type SyncStatusChangedEvent struct {
	// Type defines the type of the message.
	Type string `json:"type" bson:"type"`
	// NodeID defines the ID of the node.
	NodeID string `json:"nodeid" bson:"nodeid"`
	// MetricsLevel defines the amount of metrics that are sent by the node.
	MetricsLevel uint8 `json:"metricsLevel" bson:"metricsLevel"`
	// Time defines the time when the sync status changed.
	Time time.Time `json:"datetime" bson:"datetime"`
	// CurrentStatus contains current sync status
	CurrentStatus bool `json:"currentStatus" bson:"currentStatus"`
	// PreviousStatus contains previous sync status
	PreviousStatus bool `json:"previousStatus" bson:"previousStatus"`
	// LastConfirmedMessageTime contains time of the last confirmed message
	LastConfirmedMessageTime time.Time `json:"lastConfirmedMessageTime" bson:"lastConfirmedMessageTime"`
}

// MessageFinalizedMetrics defines the transaction metrics record to sent be to remote logger.
type MessageFinalizedMetrics struct {
	Type               string    `json:"type" bson:"type"`
	NodeID             string    `json:"nodeID" bson:"nodeID"`
	MetricsLevel       uint8     `json:"metricsLevel" bson:"metricsLevel"`
	MessageID          string    `json:"messageID" bson:"messageID"`
	TransactionID      string    `json:"transactionID,omitempty" bson:"transactionID"`
	IssuedTimestamp    time.Time `json:"issuedTimestamp" bson:"issuedTimestamp"`
	SolidTimestamp     time.Time `json:"solidTimestamp,omitempty" bson:"solidTimestamp"`
	ScheduledTimestamp time.Time `json:"scheduledTimestamp" bson:"scheduledTimestamp"`
	BookedTimestamp    time.Time `json:"bookedTimestamp" bson:"bookedTimestamp"`
	ConfirmedTimestamp time.Time `json:"confirmedTimestamp" bson:"confirmedTimestamp"`
	DeltaSolid         int64     `json:"deltaSolid,omitempty" bson:"deltaSolid"`
	DeltaScheduled     int64     `json:"deltaArrival" bson:"deltaArrival"`
	DeltaBooked        int64     `json:"deltaBooked" bson:"deltaBooked"`
	DeltaConfirmed     int64     `json:"deltaConfirmed" bson:"deltaConfirmed"`
	StrongParentCount  int       `json:"strongParentCount" bson:"strongParentCount"`
	WeakParentsCount   int       `json:"weakParentsCount,omitempty" bson:"weakParentsCount"`
	LikeParentCount    int       `json:"likeParentCount,omitempty" bson:"likeParentCount"`
}

// BranchConfirmationMetrics defines the branch confirmation metrics record to sent be to remote logger.
type BranchConfirmationMetrics struct {
	Type               string                     `json:"type" bson:"type"`
	NodeID             string                     `json:"nodeID" bson:"nodeID"`
	MetricsLevel       uint8                      `json:"metricsLevel" bson:"metricsLevel"`
	MessageID          string                     `json:"messageID" bson:"messageID"`
	BranchID           string                     `json:"transactionID" bson:"transactionID"`
	CreatedTimestamp   time.Time                  `json:"createdTimestamp" bson:"createdTimestamp"`
	ConfirmedTimestamp time.Time                  `json:"confirmedTimestamp" bson:"confirmedTimestamp"`
	DeltaConfirmed     int64                      `json:"deltaConfirmed" bson:"deltaConfirmed"`
	InclusionState     ledgerstate.InclusionState `json:"inclusionState" bson:"inclusionState"`
}

// BranchCountUpdate defines the branch confirmation metrics record to sent be to remote logger.
type BranchCountUpdate struct {
	Type                 string `json:"type" bson:"type"`
	NodeID               string `json:"nodeID" bson:"nodeID"`
	MetricsLevel         uint8  `json:"metricsLevel" bson:"metricsLevel"`
	TotalBranchCount     uint64 `json:"totalBranchCount" bson:"totalBranchCount"`
	FinalizedBranchCount uint64 `json:"finalizedBranchCount" bson:"finalizedBranchCount"`
}

// DRNGMetrics defines the DRNG metrics record to sent be to remote logger.
type DRNGMetrics struct {
	Type              string    `json:"type" bson:"type"`
	NodeID            string    `json:"nodeID" bson:"nodeID"`
	MetricsLevel      uint8     `json:"metricsLevel" bson:"metricsLevel"`
	InstanceID        uint32    `json:"instanceID" bson:"instanceID"`
	Round             uint64    `json:"round" bson:"round"`
	IssuedTimestamp   time.Time `json:"issuedTimestamp" bson:"issuedTimestamp"`
	ReceivedTimestamp time.Time `json:"receivedTimestamp" bson:"receivedTimestamp"`
	DeltaReceived     int64     `json:"deltaReceived"  bson:"deltaReceived"`
}
