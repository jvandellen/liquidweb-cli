/*
Copyright © LiquidWeb

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
package apiTypes

import (
	"fmt"
	"strings"
)

type NetworkIpPoolListEntry struct {
	Id     int64 `json:"id" mapstructure:"id"`
	ZoneId int64 `json:"zone_id" mapstructure:"zone_id"`
}

type NetworkIpPoolDetails struct {
	Accnt       int64                            `json:"accnt" mapstructure:"accnt"`
	Id          int64                            `json:"id" mapstructure:"id"`
	UniqId      string                           `json:"uniq_id" mapstructure:"uniq_id"`
	ZoneId      int64                            `json:"zone_id" mapstructure:"zone_id"`
	Assignments []NetworkIpPoolDetailsAssignment `json:"assignments" mapstructure:"assignments"`
}

type NetworkIpPoolDetailsAssignment struct {
	BeginRange string `json:"begin_range" mapstructure:"begin_range"`
	Broadcast  string `json:"broadcast" mapstructure:"broadcast"`
	EndRange   string `json:"end_range" mapstructure:"end_range"`
	Gateway    string `json:"gateway" mapstructure:"gateway"`
	Id         int64  `json:"id" mapstructure:"id"`
	Netmask    string `json:"netmask" mapstructure:"netmask"`
	Network    string `json:"network" mapstructure:"network"`
	ZoneId     int64  `json:"zone_id" mapstructure:"zone_id"`
}

func (x NetworkIpPoolDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("IP Pool id [%d] uniq_id [%s]\n", x.Id, x.UniqId))
	slice = append(slice, fmt.Sprintf("\tZoneId: %d\n", x.ZoneId))
	slice = append(slice, fmt.Sprintf("\tAccount: %d\n", x.Accnt))
	slice = append(slice, fmt.Sprintf("\tAssignments:\n"))
	for _, assignment := range x.Assignments {
		slice = append(slice, fmt.Sprintf("\t\tassignment:\n"))
		slice = append(slice, fmt.Sprintf("\t\t\tBeginRange: %s\n", assignment.BeginRange))
		slice = append(slice, fmt.Sprintf("\t\t\tEndRange: %s\n", assignment.EndRange))
		if assignment.Broadcast != "" {
			slice = append(slice, fmt.Sprintf("\t\t\tBroadcast: %s\n", assignment.Broadcast))
		}
		slice = append(slice, fmt.Sprintf("\t\t\tGateway: %s\n", assignment.Gateway))
		slice = append(slice, fmt.Sprintf("\t\t\tNetmask: %s\n", assignment.Netmask))
		slice = append(slice, fmt.Sprintf("\t\t\tNetwork: %s\n", assignment.Network))
		slice = append(slice, fmt.Sprintf("\t\t\tId: %d\n", assignment.Id))
		slice = append(slice, fmt.Sprintf("\t\t\tZoneId: %d\n", assignment.ZoneId))
	}

	return strings.Join(slice[:], "")
}

type NetworkIpPoolDelete struct {
	Deleted bool `json:"deleted" mapstructure:"deleted"`
}

type NetworkIpAdd struct {
	Adding string `json:"adding" mapstructure:"adding"`
}

type NetworkIpRemove struct {
	Removing string `json:"removing" mapstructure:"removing"`
}

type NetworkAssignmentListEntry struct {
	Broadcast string `json:"broadcast" mapstructure:"broadcast"`
	Ip        string `json:"ip" mapstructure:"ip"`
	Gateway   string `json:"gateway" mapstructure:"gateway"`
	Id        int64  `json:"id" mapstructure:"id"`
	Netmask   string `json:"netmask" mapstructure:"netmask"`
	Network   string `json:"network" mapstructure:"network"`
}

func (x NetworkAssignmentListEntry) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("\tIP: %s\n", x.Ip))
	slice = append(slice, fmt.Sprintf("\t\tId: %d\n", x.Id))
	slice = append(slice, fmt.Sprintf("\t\tGateway: %s\n", x.Gateway))
	slice = append(slice, fmt.Sprintf("\t\tBroadcast: %s\n", x.Broadcast))
	slice = append(slice, fmt.Sprintf("\t\tNetmask: %s\n", x.Netmask))
	slice = append(slice, fmt.Sprintf("\t\tNetwork: %s\n", x.Netmask))

	return strings.Join(slice[:], "")
}

type NetworkLoadBalancerDetails struct {
	Name               string                              `json:"name" mapstructure:"name"`
	Nodes              []NetworkLoadBalancerDetailsNode    `json:"nodes" mapstructure:"nodes"`
	RegionId           int64                               `json:"region_id" mapstructure:"region_id"`
	Services           []NetworkLoadBalancerDetailsService `json:"services" mapstructure:"services"`
	SessionPersistence bool                                `json:"session_persistence" mapstructure:"session_persistence"`
	SslIncludes        bool                                `json:"ssl_includes" mapstructure:"ssl_includes"`
	SslTermination     bool                                `json:"ssl_termination" mapstructure:"ssl_termination"`
	Strategy           string                              `json:"strategy" mapstructure:"strategy"`
	UniqId             string                              `json:"uniq_id" mapstructure:"uniq_id"`
	Vip                string                              `json:"vip" mapstructure:"vip"`
}

type NetworkLoadBalancerDetailsNode struct {
	Domain string `json:"domain" mapstructure:"domain"`
	Ip     string `json:"ip" mapstructure:"ip"`
	UniqId string `json:"uniq_id" mapstructure:"mapstructure"`
}

type NetworkLoadBalancerDetailsService struct {
	DestPort     int64                                          `json:"dest_port" mapstructure:"dest_port"`
	Protocol     string                                         `json:"protocol" mapstructure:"protocol"`
	SrcPort      int64                                          `json:"src_port" mapstructure:"src_port"`
	HealthChecks []NetworkLoadBalancerDetailsServiceHealthCheck `json:"health_check" mapstructure:"health_check"`
}

type NetworkLoadBalancerDetailsServiceHealthCheck struct {
	FailureThreshold  int64   `json:"failure_threshold" mapstructure:"failure_threshold"`
	HttpBodyMatch     string  `json:"http_body_match" mapstructure:"http_body_match"`
	HttpPath          string  `json:"http_path" mapstructure:"http_path"`
	HttpResponseCodes []int64 `json:"http_response_codes" mapstructure:"http_response_codes"`
	HttpUseTls        bool    `json:"http_use_tls" mapstructure:"http_use_tls"`
	Interval          int64   `json:"interval" mapstructure:"interval"`
	Protocol          string  `json:"protocol" mapstructure:"protocol"`
	Timeout           int64   `json:"timeout" mapstructure:"timeout"`
}

func (x NetworkLoadBalancerDetails) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Name: %s\n", x.Name))
	slice = append(slice, fmt.Sprintf("\tUniqId: %s\n", x.UniqId))
	slice = append(slice, fmt.Sprintf("\tRegionId: %d\n", x.RegionId))
	slice = append(slice, fmt.Sprintf("\tVip: %s\n", x.Vip))
	slice = append(slice, fmt.Sprintf("\tStrategy: %s\n", x.Strategy))
	slice = append(slice, fmt.Sprintf("\tSession Persistence: %t\n", x.SessionPersistence))
	slice = append(slice, fmt.Sprintf("\tSSL Termination: %t\n", x.SslTermination))
	slice = append(slice, fmt.Sprintf("\tSSL Includes: %t\n", x.SslIncludes))
	slice = append(slice, "\tNodes:\n")
	for _, node := range x.Nodes {
		slice = append(slice, "\t\tNode:\n")
		slice = append(slice, fmt.Sprintf("\t\t\tDomain: %s\n", node.Domain))
		slice = append(slice, fmt.Sprintf("\t\t\tIP: %s\n", node.Ip))
		slice = append(slice, fmt.Sprintf("\t\t\tUniqId: %s\n", node.UniqId))
	}
	slice = append(slice, "\tServices:\n")
	for _, service := range x.Services {
		slice = append(slice, "\t\tService:\n")
		slice = append(slice, fmt.Sprintf("\t\t\tProtocol: %s\n", service.Protocol))
		slice = append(slice, fmt.Sprintf("\t\t\tSource Port: %d\n", service.SrcPort))
		slice = append(slice, fmt.Sprintf("\t\t\tDestination Port: %d\n", service.DestPort))
		slice = append(slice, fmt.Sprintf("\t\t\tHealth Checks: %+v\n", service.HealthChecks))
	}

	return strings.Join(slice[:], "")
}

type NetworkLoadBalancerStrategies struct {
	Strategies []NetworkLoadBalancerStrategy `json:"strategies" mapstructure:"strategies"`
}

type NetworkLoadBalancerStrategy struct {
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	Strategy    string `json:"strategy" mapstructure:"strategy"`
}

func (x NetworkLoadBalancerStrategies) String() string {
	var slice []string

	slice = append(slice, "Strategies:\n")
	for _, strategy := range x.Strategies {
		slice = append(slice, fmt.Sprintf("\tName: %s\n", strategy.Name))
		slice = append(slice, fmt.Sprintf("\t\tDescription: %s\n", strategy.Description))
		slice = append(slice, fmt.Sprintf("\t\tStrategy: %s\n", strategy.Strategy))
	}

	return strings.Join(slice[:], "")
}

type NetworkLoadBalancerPossibleNodes struct {
	Items []NetworkLoadBalancerPossibleNodesNode `json:"items" mapstructure:"items"`
}

type NetworkLoadBalancerPossibleNodesNode struct {
	Domain   string `json:"domain" mapstructure:"domain"`
	Ip       string `json:"domain" mapstructure:"domain"`
	RegionId int64  `json:"region_id" mapstructure:"region_id"`
	UniqId   string `json:"uniq_id" mapstructure:"uniq_id"`
}

func (x NetworkLoadBalancerPossibleNodes) String() string {
	var slice []string

	slice = append(slice, "Possible Nodes:\n")
	for _, possibleNode := range x.Items {
		slice = append(slice, fmt.Sprintf("\tDomain: %s\n", possibleNode.Domain))
		slice = append(slice, fmt.Sprintf("\t\tUniqId: %s\n", possibleNode.UniqId))
		slice = append(slice, fmt.Sprintf("\t\tIP: %s\n", possibleNode.Ip))
		slice = append(slice, fmt.Sprintf("\t\tRegionId: %d\n", possibleNode.RegionId))
	}

	return strings.Join(slice[:], "")
}
