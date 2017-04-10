package filesystem

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// AppendModeType enumerates the values for append mode type.
type AppendModeType string

const (
	// Autocreate specifies the autocreate state for append mode type.
	Autocreate AppendModeType = "autocreate"
)

// FileType enumerates the values for file type.
type FileType string

const (
	// DIRECTORY specifies the directory state for file type.
	DIRECTORY FileType = "DIRECTORY"
	// FILE specifies the file state for file type.
	FILE FileType = "FILE"
)

// AclStatus is data Lake Store file or directory Access Control List
// information.
type AclStatus struct {
	Entries   *[]string `json:"entries,omitempty"`
	Group     *string   `json:"group,omitempty"`
	Owner     *string   `json:"owner,omitempty"`
	StickyBit *bool     `json:"stickyBit,omitempty"`
}

// AclStatusResult is data Lake Store file or directory Access Control List
// information.
type AclStatusResult struct {
	autorest.Response `json:"-"`
	AclStatus         *AclStatus `json:"AclStatus,omitempty"`
}

// AdlsAccessControlException is a WebHDFS exception thrown indicating that
// access is denied due to insufficient permissions. Thrown when a 403 error
// response code is returned (forbidden).
type AdlsAccessControlException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsBadOffsetException is a WebHDFS exception thrown indicating the append
// or read is from a bad offset. Thrown when a 400 error response code is
// returned for append and open operations (Bad request).
type AdlsBadOffsetException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsError is data Lake Store filesystem error containing a specific WebHDFS
// exception.
type AdlsError struct {
	RemoteException *AdlsRemoteException `json:"RemoteException,omitempty"`
}

// AdlsFileAlreadyExistsException is a WebHDFS exception thrown indicating the
// file or folder already exists. Thrown when a 403 error response code is
// returned (forbidden).
type AdlsFileAlreadyExistsException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsFileNotFoundException is a WebHDFS exception thrown indicating the file
// or folder could not be found. Thrown when a 404 error response code is
// returned (not found).
type AdlsFileNotFoundException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsIllegalArgumentException is a WebHDFS exception thrown indicating that
// one more arguments is incorrect. Thrown when a 400 error response code is
// returned (bad request).
type AdlsIllegalArgumentException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsIOException is a WebHDFS exception thrown indicating there was an IO
// (read or write) error. Thrown when a 403 error response code is returned
// (forbidden).
type AdlsIOException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsRemoteException is data Lake Store filesystem exception based on the
// WebHDFS definition for RemoteExceptions.
type AdlsRemoteException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsRuntimeException is a WebHDFS exception thrown when an unexpected error
// occurs during an operation. Thrown when a 500 error response code is
// returned (Internal server error).
type AdlsRuntimeException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsSecurityException is a WebHDFS exception thrown indicating that access
// is denied. Thrown when a 401 error response code is returned
// (Unauthorized).
type AdlsSecurityException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// AdlsUnsupportedOperationException is a WebHDFS exception thrown indicating
// that the requested operation is not supported. Thrown when a 400 error
// response code is returned (bad request).
type AdlsUnsupportedOperationException struct {
	JavaClassName *string `json:"javaClassName,omitempty"`
	Message       *string `json:"message,omitempty"`
}

// ContentSummary is data Lake Store content summary information
type ContentSummary struct {
	DirectoryCount *int64 `json:"directoryCount,omitempty"`
	FileCount      *int64 `json:"fileCount,omitempty"`
	Length         *int64 `json:"length,omitempty"`
	SpaceConsumed  *int64 `json:"spaceConsumed,omitempty"`
}

// ContentSummaryResult is data Lake Store filesystem content summary
// information response.
type ContentSummaryResult struct {
	autorest.Response `json:"-"`
	ContentSummary    *ContentSummary `json:"ContentSummary,omitempty"`
}

// FileOperationResult is the result of the request or operation.
type FileOperationResult struct {
	autorest.Response `json:"-"`
	Boolean           *bool `json:"boolean,omitempty"`
}

// FileStatuses is data Lake Store file status list information.
type FileStatuses struct {
	FileStatus *[]FileStatusProperties `json:"FileStatus,omitempty"`
}

// FileStatusesResult is data Lake Store filesystem file status list
// information response.
type FileStatusesResult struct {
	autorest.Response `json:"-"`
	FileStatuses      *FileStatuses `json:"FileStatuses,omitempty"`
}

// FileStatusProperties is data Lake Store file or directory information.
type FileStatusProperties struct {
	AccessTime       *int64   `json:"accessTime,omitempty"`
	BlockSize        *int64   `json:"blockSize,omitempty"`
	ChildrenNum      *int64   `json:"childrenNum,omitempty"`
	Group            *string  `json:"group,omitempty"`
	Length           *int64   `json:"length,omitempty"`
	ModificationTime *int64   `json:"modificationTime,omitempty"`
	Owner            *string  `json:"owner,omitempty"`
	PathSuffix       *string  `json:"pathSuffix,omitempty"`
	Permission       *string  `json:"permission,omitempty"`
	Type             FileType `json:"type,omitempty"`
}

// FileStatusResult is data Lake Store filesystem file status information
// response.
type FileStatusResult struct {
	autorest.Response `json:"-"`
	FileStatus        *FileStatusProperties `json:"FileStatus,omitempty"`
}

// ReadCloser is
type ReadCloser struct {
	autorest.Response `json:"-"`
	Value             *io.ReadCloser `json:"value,omitempty"`
}
