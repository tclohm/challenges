package main

import (
	_"fmt"
	"testing"
)

func TestNewSocialGraph(test *testing.T) {
	var graph *SocialGraph

	graph = NewSocialGraph()

	if graph == nil {
		test.Errorf("error in creating a graph")
	}
}