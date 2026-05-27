// Copyright (c) 2018-2022 Burak Sezer
// All rights reserved.
//
// This code is licensed under the MIT License.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files(the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and / or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions :
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package consistent provides a consistent hashing function with bounded loads. This implementation also adds
// partitioning logic on top of the original algorithm. For more information about the underlying algorithm,
// please take a look at https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
//
// Example Use:
//
//	cfg := consistent.Config{
//		PartitionCount:    71,
//		ReplicationFactor: 20,
//		Load:              1.25,
//		Hasher:            hasher{},
//	}
//
// Now you can create a new Consistent instance. This function can take a list of the members.
//
//	c := consistent.New(members, cfg)
//
// In the following sample, you add a new Member to the consistent hash ring. myMember is just a Go struct that
// implements the Member interface. You should know that modifying the consistent hash ring distributes partitions among
// members using the algorithm defined on Google Research Blog.
//
//	c.Add(myMember)
//
// Remove a member from the consistent hash ring:
//
//	c.Remove(member-name)
//
// LocateKey hashes the key and calculates partition ID with this modulo operation: MOD(hash result, partition count)
// The owner of the partition is already calculated by New/Add/Remove. LocateKey just returns the member that is responsible
// for the key.
//
//	key := []byte("my-key")
//	member := c.LocateKey(key)
package consistent

import (
	"errors"
	"sync"
)

const (
	DefaultPartitionCount    int     = 271
	DefaultReplicationFactor int     = 20
	DefaultLoad              float64 = 1.25
)

// ErrInsufficientMemberCount represents an error which means there are not enough members to complete the task.
var ErrInsufficientMemberCount = errors.New("insufficient member count")

// Hasher is responsible for generating unsigned, 64-bit hash of provided byte slice.
// Hasher should minimize collisions (generating same hash for different byte slice)
// and while performance is also important fast functions are preferable (i.e.
// you can use FarmHash family).
type Hasher interface {
	Sum64([]byte) uint64
}

// Member interface represents a member in consistent hash ring.
type Member interface {
	String() string
}

// Config represents a structure to control consistent package.
type Config struct {
	// Hasher is responsible for generating unsigned, 64-bit hash of provided byte slice.
	Hasher Hasher

	// Keys are distributed among partitions. Prime numbers are good to
	// distribute keys uniformly. Select a big PartitionCount if you have
	// too many keys.
	PartitionCount int

	// Members are replicated on consistent hash ring. This number means that a member
	// how many times replicated on the ring.
	ReplicationFactor int

	// Load is used to calculate average load. See the code, the paper and Google's blog post to learn about it.
	Load float64
}

// Consistent holds the information about the members of the consistent hash circle.
type Consistent struct {
	mu sync.RWMutex

	config         Config
	hasher         Hasher
	sortedSet      []uint64
	partitionCount uint64
	loads          map[string]float64
	members        map[string]*Member
	partitions     map[int]*Member
	ring           map[uint64]*Member
}

// New creates and returns a new Consistent object.
func New(members []Member, config Config) *Consistent { _ = "STUB: not implemented"; return nil }

// GetMembers returns a thread-safe copy of members. If there are no members, it returns an empty slice of Member.
func (c *Consistent) GetMembers() []Member { _ = "STUB: not implemented"; return nil }

// Create a thread-safe copy of member list.

// AverageLoad exposes the current average load.
func (c *Consistent) AverageLoad() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Consistent) averageLoad() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Consistent) distributeWithLoad(partID, idx int, partitions map[int]*Member, loads map[string]float64) {
	_ = "STUB: not implemented"
	return
}

// User needs to decrease partition count, increase member count or increase load factor.

func (c *Consistent) distributePartitions() { _ = "STUB: not implemented"; return }

func (c *Consistent) add(member Member) { _ = "STUB: not implemented"; return }

// sort hashes ascendingly

// Storing member at this map is useful to find backup members of a partition.

// Add adds a new member to the consistent hash circle.
func (c *Consistent) Add(member Member) { _ = "STUB: not implemented"; return }

// We already have this member. Quit immediately.

func (c *Consistent) delSlice(val uint64) { _ = "STUB: not implemented"; return }

// Remove removes a member from the consistent hash circle.
func (c *Consistent) Remove(name string) { _ = "STUB: not implemented"; return }

// There is no member with that name. Quit immediately.

// consistent hash ring is empty now. Reset the partition table.

// LoadDistribution exposes load distribution of members.
func (c *Consistent) LoadDistribution() map[string]float64 { _ = "STUB: not implemented"; return nil }

// Create a thread-safe copy

// FindPartitionID returns partition id for given key.
func (c *Consistent) FindPartitionID(key []byte) int { _ = "STUB: not implemented"; return 0 }

// GetPartitionOwner returns the owner of the given partition.
func (c *Consistent) GetPartitionOwner(partID int) Member {
	_ = "STUB: not implemented"
	return *new(Member)
}

// getPartitionOwner returns the owner of the given partition. It's not thread-safe.
func (c *Consistent) getPartitionOwner(partID int) Member {
	_ = "STUB: not implemented"
	return *new(Member)
}

// Create a thread-safe copy of member and return it.

// LocateKey finds a home for given key
func (c *Consistent) LocateKey(key []byte) Member { _ = "STUB: not implemented"; return *new(Member) }

func (c *Consistent) getClosestN(partID, count int) ([]Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hash and sort all the names.

// Find the key owner

// Find the closest(replica owners) members.

// GetClosestN returns the closest N member to a key in the hash ring.
// This may be useful to find members for replication.
func (c *Consistent) GetClosestN(key []byte, count int) ([]Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClosestNForPartition returns the closest N member for given partition.
// This may be useful to find members for replication.
func (c *Consistent) GetClosestNForPartition(partID, count int) ([]Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
