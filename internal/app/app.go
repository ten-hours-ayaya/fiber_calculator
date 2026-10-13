package app

import (
	"fmt"
	"log"

	"github.com/ten-hours-ayaya/fiber-calculator/internal/model"
)

type App struct {
	dividers []*model.Divider
	targets  []*model.Target
}

func NewApp(dividers []*model.Divider, targets []*model.Target) *App {
	return &App{dividers: dividers, targets: targets}
}

func (a *App) Calculate(input float32, streets []*model.Street) (*model.Result, error) {
	streetResults := make([]*model.StreetResult, len(streets))
	var lastStreetResult *model.StreetResult
	for i, street := range streets {
		streetInput := input
		if lastStreetResult != nil {
			streetInput = lastStreetResult.Input - lastStreetResult.Divider.Right
		}
		streetResult, err := a.calculateStreetResult(streetInput, *street)
		streetResults[i] = streetResult
		lastStreetResult = streetResult
		if err != nil {
			return &model.Result{StreetResults: streetResults}, err
		}
	}

	return &model.Result{StreetResults: streetResults}, nil
}

func (a *App) calculateStreetResult(input float32, street model.Street) (*model.StreetResult, error) {
	result := &model.StreetResult{StreetName: street.Name, Input: input}
	var lastDivider *model.Divider
	var lastNodes []*model.Node
	for _, divider := range a.dividers {
		lastDivider = divider
		firstNodeInput := input - divider.Left
		nodes, err := a.calculateStreetResultNodes(firstNodeInput, street)
		lastNodes = nodes
		if err == nil {
			result.Divider = divider
			result.Nodes = nodes
			return result, nil
		}
	}
	result.Divider = lastDivider
	result.Nodes = lastNodes
	return result, fmt.Errorf("Cannot calculate street: %v", street.Name)
}

func (a *App) calculateStreetResultNodes(firstNodeInput float32, street model.Street) ([]*model.Node, error) {
	nodesCount := len(street.Ends)
	nodes := make([]*model.Node, nodesCount)
	var prevNode *model.Node
	for i, end := range street.Ends {
		target := a.getTarget(end)
		nodeInput := firstNodeInput
		if prevNode != nil {
			nodeInput = prevNode.Input - prevNode.Divider.Right
		}
		if i < nodesCount-1 {
			node, err := a.calculateStreetResultNode(nodeInput, target)
			if err != nil {
				return nodes, fmt.Errorf("Error calculation nodes for street %v", street.Name)
			}
			prevNode = node
			nodes[i] = node
		} else if nodeInput >= target {
			node := &model.Node{Input: nodeInput, Target: target, Divider: nil}
			nodes[i] = node
		} else {
			return nil, fmt.Errorf("Error calculating last end for street: %v", street.Name)
		}
	}
	return nodes, nil
}

func (a *App) calculateStreetResultNode(input float32, target float32) (*model.Node, error) {
	if input <= target {
		return nil, fmt.Errorf("Input (%v) less then target (%v)", input, target)
	}
	for _, divider := range a.dividers {
		value := input - divider.Left
		if value >= target {
			return &model.Node{Input: input, Target: target, Divider: divider}, nil
		}
	}
	return nil, fmt.Errorf("Divider not found for input = %v, target = %v", input, target)
}

func (a *App) getTarget(end int) float32 {
	for _, target := range a.targets {
		if end == target.For {
			return target.Exit
		}
	}
	log.Fatalf("Cannot find target for end: %d", end)
	return 0
}
