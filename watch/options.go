/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package watch

import (
	pbhook "bscp.io/pkg/protocol/core/hook"

	"github.com/TencentBlueKing/bscp-go/types"
)

// WatchOptions options for watch bscp config items
type WatchOptions struct {
	// FeedAddrs bscp feed server addresses
	FeedAddrs []string
	// DialTimeoutMS dial timeout milliseconds
	DialTimeoutMS int64
	// Fingerprint watch fingerprint
	Fingerprint string
	// BizID watch biz id
	BizID uint32
}

// TODO: how to wrapper file, kv, table types in one watch result
// Callback watch callback
type Callback func(releaseID uint32, files []*types.ConfigItemFile, pre *pbhook.HookSpec, post *pbhook.HookSpec) error

// SubscribeOptions options for subscribe
type SubscribeOptions struct {
	// App BSCP app name
	App string
	// Key watch config item key
	Key string
	// Labels instance labels
	Labels map[string]string
	// UID instance unique uid
	UID string
}

// SubscribeOption setter for subscribe options
type SubscribeOption func(*SubscribeOptions)

// WithApp set watch app name
func WithApp(app string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.App = app
	}
}

// WithKey set watch config item key
func WithKey(key string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Key = key
	}
}

// WithLabels set watch labels
func WithLabels(labels map[string]string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.Labels = labels
	}
}

// WithUID set watch uid
func WithUID(uid string) SubscribeOption {
	return func(o *SubscribeOptions) {
		o.UID = uid
	}
}
