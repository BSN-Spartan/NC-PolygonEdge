package propose

import (
	"context"
	"errors"
	ibftOp "github.com/0xPolygon/polygon-edge/consensus/ibft/proto"

	"github.com/0xPolygon/polygon-edge/command"
	"github.com/0xPolygon/polygon-edge/command/helper"
	"github.com/0xPolygon/polygon-edge/types"
)

const (
	voteFlag    = "vote"
	addressFlag = "addr"
	modeFlag    = "mode"
)

const (
	authVote = "auth"
	dropVote = "drop"

	modeConsensus = "consensus"
	modeWhite     = "white"
)

var (
	errInvalidVoteType      = errors.New("invalid vote type")
	errInvalidAddressFormat = errors.New("invalid address format")
)

var (
	params = &proposeParams{}
)

type proposeParams struct {
	addressRaw string

	vote    string
	address types.Address

	mode string
}

func (p *proposeParams) getRequiredFlags() []string {
	return []string{
		voteFlag,
		addressFlag,
		modeFlag,
	}
}

func (p *proposeParams) validateFlags() error {
	if !isValidVoteType(p.vote) {
		return errInvalidVoteType
	}

	return nil
}

func (p *proposeParams) initRawParams() error {
	p.address = types.Address{}
	if err := p.address.UnmarshalText([]byte(p.addressRaw)); err != nil {
		return errInvalidAddressFormat
	}

	return nil
}

func isValidVoteType(vote string) bool {
	return vote == authVote || vote == dropVote
}

func (p *proposeParams) proposeCandidate(grpcAddress string) error {
	ibftClient, err := helper.GetIBFTOperatorClientConnection(grpcAddress)
	if err != nil {
		return err
	}

	var candidateMode ibftOp.CandidateType
	if p.mode == modeWhite {
		candidateMode = ibftOp.CandidateType_whiteAccount
	} else {
		candidateMode = ibftOp.CandidateType_consensus
	}

	if _, err := ibftClient.Propose(
		context.Background(),
		&ibftOp.Candidate{
			Address: p.address.String(),
			Auth:    p.vote == authVote,
			Mode:    candidateMode,
		},
	); err != nil {
		return err
	}

	return nil
}

func (p *proposeParams) getResult() command.CommandResult {
	return &IBFTProposeResult{
		Address: p.address.String(),
		Vote:    p.vote,
		Mode:    p.mode,
	}
}
