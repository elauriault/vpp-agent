// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/interfaces"
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/rpc"
	"sync"
)

// Maximum number of messages stored in the buffer. Buffer is always filled from left
// to right (it means that if the buffer is full, a new entry is written to the index 0)
const bufferSize = 100

// NotificationSvc forwards GRPC messages to external servers.
type NotificationSvc struct {
	mx sync.RWMutex

	log logging.Logger

	// VPP notifications available for clients
	nBuffer [bufferSize]*rpc.NotificationsResponse
	nIdx    uint32
}

// Get returns all required VPP notifications (or those available in the buffer) in the same order as they were received
func (svc *NotificationSvc) Get(from *rpc.NotificationRequest, server rpc.NotificationService_GetServer) error {
	svc.mx.RLock()
	defer svc.mx.RUnlock()

	// If required index was not provided, return all notifications in the buffer
	if from.Idx == 0 {
		from.Idx = 1
	}

	// Send messages on the 'right' side of the buffer if needed. If required notification is stored on higher index
	// than the latest (or the gap between required and latest is bigger that buffer size), send all notifications
	// from desired index to the end of the buffer first
	var wasErr error
	if svc.nIdx-from.Idx > bufferSize || from.Idx%bufferSize > svc.nIdx%bufferSize {
		for bufferIdx, entry := range svc.nBuffer {
			// Skip empty entries (should happen only in the beginning while the buffer is empty)
			if entry == nil {
				continue
			}
			if uint32(bufferIdx) >= svc.nIdx%bufferSize && uint32(bufferIdx) >= from.Idx%bufferSize {
				if err := server.Send(entry); err != nil {
					svc.log.Error("Send notification error: %v", err)
					wasErr = err
				}
			}
		}
	}
	// Send all new notifications, looping from the beginning of the buffer.
	for bufferIdx, entry := range svc.nBuffer {
		// Skip empty entries (should happen only in the beginning while the buffer is empty)
		if entry == nil {
			continue
		}
		// Most recent notification reached
		if uint32(bufferIdx) >= svc.nIdx%bufferSize {
			break
		}
		// Send only new entries
		if from.Idx < entry.NextIdx {
			if err := server.Send(entry); err != nil {
				svc.log.Error("Send notification error: %v", err)
				wasErr = err
			}
		}
	}
	return wasErr
}

// Adds new notification to the pool. The order of notifications is preserved
func (svc *NotificationSvc) updateNotifications(ctx context.Context, notification *interfaces.InterfaceNotification) {
	svc.mx.Lock()
	defer svc.mx.Unlock()

	// Notification index starts with 1
	svc.nBuffer[svc.nIdx%bufferSize] = &rpc.NotificationsResponse{
		NextIdx: svc.nIdx + 1,
		NIf:     notification,
	}
	svc.nIdx++
}
