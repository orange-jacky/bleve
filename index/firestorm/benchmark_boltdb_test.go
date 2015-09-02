//  Copyright (c) 2015 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package firestorm

import (
	"os"
	"testing"

	"github.com/blevesearch/bleve/index/store"
	"github.com/blevesearch/bleve/index/store/boltdb"
)

func CreateBoltDB() (store.KVStore, error) {
	s := boltdb.New("test", "bleve")
	return s, nil
}

func DestroyBoltDB() error {
	return os.RemoveAll("test")
}

func BenchmarkBoltDBIndexing1Workers(b *testing.B) {
	CommonBenchmarkIndex(b, CreateBoltDB, DestroyBoltDB, 1)
}

func BenchmarkBoltDBIndexing2Workers(b *testing.B) {
	CommonBenchmarkIndex(b, CreateBoltDB, DestroyBoltDB, 2)
}

func BenchmarkBoltDBIndexing4Workers(b *testing.B) {
	CommonBenchmarkIndex(b, CreateBoltDB, DestroyBoltDB, 4)
}

// batches

func BenchmarkBoltDBIndexing1Workers10Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 1, 10)
}

func BenchmarkBoltDBIndexing2Workers10Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 2, 10)
}

func BenchmarkBoltDBIndexing4Workers10Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 4, 10)
}

func BenchmarkBoltDBIndexing1Workers100Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 1, 100)
}

func BenchmarkBoltDBIndexing2Workers100Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 2, 100)
}

func BenchmarkBoltDBIndexing4Workers100Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 4, 100)
}

func BenchmarkBoltBIndexing1Workers1000Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 1, 1000)
}

func BenchmarkBoltBIndexing2Workers1000Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 2, 1000)
}

func BenchmarkBoltBIndexing4Workers1000Batch(b *testing.B) {
	CommonBenchmarkIndexBatch(b, CreateBoltDB, DestroyBoltDB, 4, 1000)
}