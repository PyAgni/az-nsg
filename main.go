package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"log"
	"os"
)

func main() {
	fmt.Println("abcd")
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Printf("Error while creating an new authentication, %v ", err)

	}
	//fmt.Println(authorizer)

	var subscriptionId string = os.Getenv("SUBSCRIPTION_ID")

	nsgClient := network.NewSecurityGroupsClient(subscriptionId)
	nsgClient.Authorizer = authorizer

	group, error := nsgClient.Get(context.Background(), "demorg",
		"staging-deployer-nsg", "")
	if error != nil {
		log.Printf("Abce %v", error)
	}
	fmt.Println(*group.ID)

	rulesClient := network.NewSecurityRulesClient(subscriptionId)
	rulesClient.Authorizer = authorizer

	rule, _ := rulesClient.Get(context.Background(), "demorg",
		"staging-deployer-nsg", "SSH")

	var IPs []string = []string{"abvc"}
	fmt.Println(IPs)

	fmt.Println(*rule.SourceAddressPrefixes)
	//fmt.Println(rule)
	//rgClient := resources.NewGroupsClient(subscriptionId)
	//rgClient.Authorizer = authorizer
	//ctx := context.Background()
	//for list, err := rgClient.ListComplete(ctx, "", nil); list.NotDone(); err = list.NextWithContext(ctx) {
	//	if err != nil {
	//		log.Fatalf("got error: %s", err)
	//	}
	//	rgName := *list.Value().Name
	//	fmt.Println(rgName)
	//}
}
