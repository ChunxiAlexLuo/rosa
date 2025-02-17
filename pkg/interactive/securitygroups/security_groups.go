package securitygroups

import (
	"os"

	"github.com/openshift/rosa/pkg/aws"
	. "github.com/openshift/rosa/pkg/interactive"
	"github.com/openshift/rosa/pkg/rosa"
	"github.com/spf13/cobra"
)

func GetSecurityGroupIds(r *rosa.Runtime, cmd *cobra.Command,
	targetVpcId string, securityGroupFlag string) []string {
	possibleSgs, err := r.AWSClient.GetSecurityGroupIds(targetVpcId)
	if err != nil {
		r.Reporter.Errorf("There was a problem retrieving security groups for VPC '%s': %v", targetVpcId, err)
		os.Exit(1)
	}
	securityGroupIds := []string{}
	if len(possibleSgs) > 0 {
		options := []string{}
		for _, sg := range possibleSgs {
			options = append(options, aws.SetSecurityGroupOption(sg))
		}
		securityGroupIds, err = GetMultipleOptions(Input{
			Question: "Additional Security Group IDs",
			Help:     cmd.Flags().Lookup(securityGroupFlag).Usage,
			Required: false,
			Options:  options,
		})
		if err != nil {
			r.Reporter.Errorf("Expected valid Security Group IDs: %s", err)
			os.Exit(1)
		}
		for i, sg := range securityGroupIds {
			securityGroupIds[i] = aws.ParseOption(sg)
		}
	}
	return securityGroupIds
}
