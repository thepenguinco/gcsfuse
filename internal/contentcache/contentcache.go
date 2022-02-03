// Copyright 2021 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package contentcache stores GCS object contents locally.
package contentcache

import (
	"io"

	"github.com/googlecloudplatform/gcsfuse/internal/gcsx"
	"github.com/jacobsa/timeutil"
)

// ContentCache is a directory on local disk to store the object content.
type ContentCache struct {
	tempDir        string
	localFileCache bool
	// TODO do not expose FileMap through public interface
	// Use Add() and Remove() methods to interact with the filemap
	// Filemap maps canononical file prefixes to gcsx.TempFile, wrapper for
	// temp files on disk cache
	FileMap    map[string]gcsx.TempFile
	mtimeClock timeutil.Clock
}

// New creates a ContentCache.
func New(tempDir string, mtimeClock timeutil.Clock) *ContentCache {
	return &ContentCache{
		tempDir:    tempDir,
		FileMap:    make(map[string]gcsx.TempFile),
		mtimeClock: mtimeClock,
	}
}

// Read the metadata and initialize the in memory map
func (c *ContentCache) ReadMetadata() {

}

// Create the content cache metadata and flush to disk
func (c *ContentCache) CreateMetadata() {

}

// NewTempFile returns a handle for a temporary file on the disk. The caller
// must call Destroy on the TempFile before releasing it.
func (c *ContentCache) NewTempFile(rc io.ReadCloser) (gcsx.TempFile, error) {
	return gcsx.NewTempFile(rc, c.tempDir, c.mtimeClock)
}
