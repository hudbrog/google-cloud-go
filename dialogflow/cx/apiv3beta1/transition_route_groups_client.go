// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newTransitionRouteGroupsClientHook clientHook

// TransitionRouteGroupsCallOptions contains the retry settings for each method of TransitionRouteGroupsClient.
type TransitionRouteGroupsCallOptions struct {
	ListTransitionRouteGroups  []gax.CallOption
	GetTransitionRouteGroup    []gax.CallOption
	CreateTransitionRouteGroup []gax.CallOption
	UpdateTransitionRouteGroup []gax.CallOption
	DeleteTransitionRouteGroup []gax.CallOption
}

func defaultTransitionRouteGroupsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("dialogflow.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTransitionRouteGroupsCallOptions() *TransitionRouteGroupsCallOptions {
	return &TransitionRouteGroupsCallOptions{
		ListTransitionRouteGroups: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// TransitionRouteGroupsClient is a client for interacting with Dialogflow API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type TransitionRouteGroupsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	transitionRouteGroupsClient cxpb.TransitionRouteGroupsClient

	// The call options for this service.
	CallOptions *TransitionRouteGroupsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTransitionRouteGroupsClient creates a new transition route groups client.
//
// Service for managing TransitionRouteGroups.
func NewTransitionRouteGroupsClient(ctx context.Context, opts ...option.ClientOption) (*TransitionRouteGroupsClient, error) {
	clientOpts := defaultTransitionRouteGroupsClientOptions()

	if newTransitionRouteGroupsClientHook != nil {
		hookOpts, err := newTransitionRouteGroupsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &TransitionRouteGroupsClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultTransitionRouteGroupsCallOptions(),

		transitionRouteGroupsClient: cxpb.NewTransitionRouteGroupsClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TransitionRouteGroupsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TransitionRouteGroupsClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TransitionRouteGroupsClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListTransitionRouteGroups returns the list of all transition route groups in the specified flow.
func (c *TransitionRouteGroupsClient) ListTransitionRouteGroups(ctx context.Context, req *cxpb.ListTransitionRouteGroupsRequest, opts ...gax.CallOption) *TransitionRouteGroupIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListTransitionRouteGroups[0:len(c.CallOptions.ListTransitionRouteGroups):len(c.CallOptions.ListTransitionRouteGroups)], opts...)
	it := &TransitionRouteGroupIterator{}
	req = proto.Clone(req).(*cxpb.ListTransitionRouteGroupsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.TransitionRouteGroup, string, error) {
		var resp *cxpb.ListTransitionRouteGroupsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.transitionRouteGroupsClient.ListTransitionRouteGroups(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTransitionRouteGroups(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// GetTransitionRouteGroup retrieves the specified TransitionRouteGroup.
func (c *TransitionRouteGroupsClient) GetTransitionRouteGroup(ctx context.Context, req *cxpb.GetTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetTransitionRouteGroup[0:len(c.CallOptions.GetTransitionRouteGroup):len(c.CallOptions.GetTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.GetTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateTransitionRouteGroup creates an TransitionRouteGroup in the specified flow.
func (c *TransitionRouteGroupsClient) CreateTransitionRouteGroup(ctx context.Context, req *cxpb.CreateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateTransitionRouteGroup[0:len(c.CallOptions.CreateTransitionRouteGroup):len(c.CallOptions.CreateTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.CreateTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateTransitionRouteGroup updates the specified TransitionRouteGroup.
func (c *TransitionRouteGroupsClient) UpdateTransitionRouteGroup(ctx context.Context, req *cxpb.UpdateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "transition_route_group.name", url.QueryEscape(req.GetTransitionRouteGroup().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateTransitionRouteGroup[0:len(c.CallOptions.UpdateTransitionRouteGroup):len(c.CallOptions.UpdateTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.UpdateTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteTransitionRouteGroup deletes the specified TransitionRouteGroup.
func (c *TransitionRouteGroupsClient) DeleteTransitionRouteGroup(ctx context.Context, req *cxpb.DeleteTransitionRouteGroupRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteTransitionRouteGroup[0:len(c.CallOptions.DeleteTransitionRouteGroup):len(c.CallOptions.DeleteTransitionRouteGroup)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.transitionRouteGroupsClient.DeleteTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// TransitionRouteGroupIterator manages a stream of *cxpb.TransitionRouteGroup.
type TransitionRouteGroupIterator struct {
	items    []*cxpb.TransitionRouteGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.TransitionRouteGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TransitionRouteGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TransitionRouteGroupIterator) Next() (*cxpb.TransitionRouteGroup, error) {
	var item *cxpb.TransitionRouteGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TransitionRouteGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *TransitionRouteGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
