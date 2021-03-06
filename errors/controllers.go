/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errors

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// HasRequeueAfterError represents that an actuator managed object should
// be requeued for further processing after the given RequeueAfter time has
// passed.
//
// DEPRECATED: This error is deprecated and should not be used for new code.
// See https://github.com/kubernetes-sigs/cluster-api/issues/3370 for more information.
//
// Users should switch their methods and functions to return a (ctrl.Result, error) pair,
// instead of relying on this error. Controller runtime exposes a Result.IsZero() (from 0.5.9, and 0.6.2)
// which can be used from callers to see if reconciliation should be stopped or continue.
type HasRequeueAfterError interface {
	// GetRequeueAfter gets the duration to wait until the managed object is
	// requeued for further processing.
	GetRequeueAfter() time.Duration
}

// RequeueAfterError represents that an actuator managed object should be
// requeued for further processing after the given RequeueAfter time has
// passed.
//
// DEPRECATED: This error is deprecated and should not be used for new code.
// See https://github.com/kubernetes-sigs/cluster-api/issues/3370 for more information.
//
// Users should switch their methods and functions to return a (ctrl.Result, error) pair,
// instead of relying on this error. Controller runtime exposes a Result.IsZero() (from 0.5.9, and 0.6.2)
// which can be used from callers to see if reconciliation should be stopped or continue.
type RequeueAfterError struct {
	RequeueAfter time.Duration
}

// Error implements the error interface
func (e *RequeueAfterError) Error() string {
	return fmt.Sprintf("requeue in %v", e.RequeueAfter)
}

// GetRequeueAfter gets the duration to wait until the managed object is
// requeued for further processing.
func (e *RequeueAfterError) GetRequeueAfter() time.Duration {
	return e.RequeueAfter
}

// IsRequeueAfter returns true if the error satisfies the interface HasRequeueAfterError.
//
// DEPRECATED: This error is deprecated and should not be used for new code.
// See https://github.com/kubernetes-sigs/cluster-api/issues/3370 for more information.
//
// Users should switch their methods and functions to return a (ctrl.Result, error) pair,
// instead of relying on this error. Controller runtime exposes a Result.IsZero() (from 0.5.9, and 0.6.2)
// which can be used from callers to see if reconciliation should be stopped or continue.
func IsRequeueAfter(err error) bool {
	_, ok := errors.Cause(err).(HasRequeueAfterError)
	return ok
}
