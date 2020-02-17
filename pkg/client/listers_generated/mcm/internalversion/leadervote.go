// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	mcm "github.com/open-cluster-management/multicloud-operators-foundation/pkg/apis/mcm"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LeaderVoteLister helps list LeaderVotes.
type LeaderVoteLister interface {
	// List lists all LeaderVotes in the indexer.
	List(selector labels.Selector) (ret []*mcm.LeaderVote, err error)
	// Get retrieves the LeaderVote from the index for a given name.
	Get(name string) (*mcm.LeaderVote, error)
	LeaderVoteListerExpansion
}

// leaderVoteLister implements the LeaderVoteLister interface.
type leaderVoteLister struct {
	indexer cache.Indexer
}

// NewLeaderVoteLister returns a new LeaderVoteLister.
func NewLeaderVoteLister(indexer cache.Indexer) LeaderVoteLister {
	return &leaderVoteLister{indexer: indexer}
}

// List lists all LeaderVotes in the indexer.
func (s *leaderVoteLister) List(selector labels.Selector) (ret []*mcm.LeaderVote, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*mcm.LeaderVote))
	})
	return ret, err
}

// Get retrieves the LeaderVote from the index for a given name.
func (s *leaderVoteLister) Get(name string) (*mcm.LeaderVote, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(mcm.Resource("leadervote"), name)
	}
	return obj.(*mcm.LeaderVote), nil
}
