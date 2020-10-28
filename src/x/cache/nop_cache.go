// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cache

import (
	"context"
	"time"
)

// NewNop returns a new nop cache.
func NewNop() Cache {
	return &nopCache{}
}

type nopCache struct{}

func (n *nopCache) Put(_ string, _ interface{})                           {}
func (n *nopCache) PutWithTTL(_ string, _ interface{}, ttl time.Duration) {}

func (n *nopCache) Get(ctx context.Context, key string, loader LoaderFunc) (interface{}, error) {
	return loader(ctx, key)
}

func (n *nopCache) GetWithTTL(ctx context.Context, key string, loader LoaderWithTTLFunc) (interface{}, error) {
	val, _, err := loader(ctx, key)
	return val, err
}

var _ Cache = &nopCache{}