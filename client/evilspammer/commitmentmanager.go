package evilspammer

import (
	"crypto/sha256"
	"github.com/iotaledger/goshimmer/client/evilwallet"
	"github.com/iotaledger/goshimmer/packages/core/commitment"
	"github.com/iotaledger/goshimmer/packages/core/epoch"
	"github.com/pkg/errors"
	"math/rand"
)

type CommitmentManager struct {
	CommitmentType  string
	ParentRefsCount int
	connector       evilwallet.Connector
}

func NewCommitmentManager() *CommitmentManager {
	return &CommitmentManager{
		ParentRefsCount: 2,
	}
}

func (c *CommitmentManager) SetConnector(connector evilwallet.Connector) {
	c.connector = connector
}

func (c *CommitmentManager) SetCommitmentType(commitmentType string) {
	c.CommitmentType = commitmentType
}

// GenerateCommitment generates a commitment based on the commitment type provided in spam details.
func (c *CommitmentManager) GenerateCommitment(clt evilwallet.Client) (*commitment.Commitment, epoch.Index, error) {
	switch c.CommitmentType {
	case "latest":
		resp, err := clt.GetLatestCommitment()
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to get latest commitment")
		}
		comm := commitment.NewEmptyCommitment()
		_, err = comm.FromBytes(resp.Bytes)
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to parse commitment bytes")
		}
		return comm, epoch.Index(resp.LatestConfirmedIndex), err
	case "random":
		resp, err := clt.GetLatestCommitment()
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to get latest commitment")
		}
		comm := commitment.NewEmptyCommitment()
		_, err = comm.FromBytes(resp.Bytes)
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to parse commitment bytes")
		}
		newCommitment := commitment.New(
			comm.Index(),
			comm.PrevID(),
			randomRoot(),
			comm.CumulativeWeight(),
		)
		return newCommitment, epoch.Index(resp.LatestConfirmedIndex), nil
	case "oldest":
		resp, err := clt.GetCommitment("0")
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to get oldest commitment")
		}
		comm := commitment.NewEmptyCommitment()
		b := resp.Bytes
		_, err = comm.FromBytes(b)
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to parse commitment bytes")
		}
		return comm, epoch.Index(resp.LatestConfirmedIndex), nil
	}
	return nil, 0, nil
}

func randomRoot() [32]byte {
	data := make([]byte, 10)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
	return sha256.Sum256(data)
}
