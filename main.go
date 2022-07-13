package main

import (
	"az-nsg/utils"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"log"
	"os"
	"strings"
)

type IP struct {
	Ip string `json:"ip"`
}

func RemoveRedundantIPs(ipList []string, currentIP string) []string {
	// Unique number that is assigned to your network.
	// First 2 bytes of an IP address.
	currentIPNetworkPart := strings.Join(strings.Split(currentIP, ".")[:2], ".")

	// If Network Part of current IP exists in an ipList element, remove it.
	for i := 0; i < len(ipList); i++ {
		if strings.Contains(ipList[i], currentIPNetworkPart) {
			//remove i-th ip address
			ipList = append(ipList[:i], ipList[i+1:]...)
		}
	}

	return ipList
}

func main() {

	ips := []string{"49.207.210.12", "49.112.123.32", "49.207.222.12", "49.321.33.21"}
	fmt.Println(RemoveRedundantIPs(ips, "49.207.12.212"))

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

	SourceAddressPrefixes := *rule.SourceAddressPrefixes
	fmt.Println(SourceAddressPrefixes)

	ip := utils.GetPublicIP()
	fmt.Println(ip)

	//fmt.Println(*rule.Body)
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
