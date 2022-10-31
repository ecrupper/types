// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package types

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"
)

func TestBuildPackage_WithBuild(t *testing.T) {
	id := int64(1)
	b := &library.Build{ID: &id}

	want := new(BuildPackage)
	want.Build = b

	// run test
	got := new(BuildPackage)

	if !reflect.DeepEqual(got.WithBuild(b), want) {
		t.Errorf("WithBuild is %v, want %v", got, want)
	}
}

func TestBuildPackage_WithPipeline(t *testing.T) {
	id := "build"
	p := &pipeline.Build{ID: id}

	want := new(BuildPackage)
	want.Pipeline = p

	// run test
	got := new(BuildPackage)

	if !reflect.DeepEqual(got.WithPipeline(p), want) {
		t.Errorf("WithPipeline is %v, want %v", got, want)
	}
}

func TestBuildPackage_WithRepo(t *testing.T) {
	id := int64(1)
	r := &library.Repo{ID: &id}

	want := new(BuildPackage)
	want.Repo = r

	// run test
	got := new(BuildPackage)

	if !reflect.DeepEqual(got.WithRepo(r), want) {
		t.Errorf("WithRepo is %v, want %v", got, want)
	}
}

func TestBuildPackage_WithUser(t *testing.T) {
	id := int64(1)
	u := &library.User{ID: &id}

	want := new(BuildPackage)
	want.User = u

	// run test
	got := new(BuildPackage)

	if !reflect.DeepEqual(got.WithUser(u), want) {
		t.Errorf("WithUser is %v, want %v", got, want)
	}
}

func TestBuildPackage_WithToken(t *testing.T) {
	want := new(BuildPackage)
	want.Token = "123abc"

	// run test
	got := new(BuildPackage)

	if !reflect.DeepEqual(got.WithToken("123abc"), want) {
		t.Errorf("WithToken is %v, want %v", got, want)
	}
}
