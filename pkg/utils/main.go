/*
Copyright © 2023 netr0m <netr0m@pm.me>
*/
package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/netr0m/az-pim-cli/pkg/pim"
)

func PrintEligibleRoles(eligibleRoleAssignments *pim.RoleAssignmentResponse) {
	var eligibleRoles = make(map[string][]string)

	for _, ras := range eligibleRoleAssignments.Value {
		subscriptionName := ras.RoleDefinition.Resource.DisplayName
		roleName := ras.RoleDefinition.DisplayName
		if _, ok := eligibleRoles[subscriptionName]; !ok {
			eligibleRoles[subscriptionName] = []string{}
		}
		eligibleRoles[subscriptionName] = append(eligibleRoles[subscriptionName], roleName)
	}

	for sub, rol := range eligibleRoles {
		fmt.Printf("== %s ==\n", sub)
		for role := range rol {
			fmt.Printf("\t - %s\n", rol[role])
		}
	}
}

func GetRoleAssignment(name interface{}, prefix interface{}, role interface{}, eligibleRoleAssignments *pim.RoleAssignmentResponse) *pim.RoleAssignment {
	if name == nil && prefix == nil {
		log.Fatalf("getSubscriptionId() requires either 'name' or 'prefix' as its input parameter")
	}
	for _, eligibleRoleAssignment := range eligibleRoleAssignments.Value {
		subscriptionName := strings.ToLower(eligibleRoleAssignment.RoleDefinition.Resource.DisplayName)
		role = strings.ToLower(role.(string))

		if prefix, exists := prefix.(string); prefix != "" && exists {
			prefix = strings.ToLower(prefix)
			if strings.HasPrefix(subscriptionName, prefix) && strings.ToLower(eligibleRoleAssignment.RoleDefinition.DisplayName) == role {
				return &eligibleRoleAssignment
			}
		} else if name, exists := name.(string); name != "" && exists {
			name = strings.ToLower(name)
			if subscriptionName == name && strings.ToLower(eligibleRoleAssignment.RoleDefinition.DisplayName) == role {
				return &eligibleRoleAssignment
			}
		}
	}
	log.Fatalln("Unable to find a role assignment matching the parameters.")

	return nil
}
