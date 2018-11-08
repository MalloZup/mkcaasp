package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"utilities"
)

const (
	command  = "terraform %s -var auth_url=$OS_AUTH_URL -var domain_name=$OS_USER_DOMAIN_NAME -var region_name=$OS_REGION_NAME -var project_name=$OS_PROJECT_NAME -var user_name=$OS_USERNAME -var password=$OS_PASSWORD -var-file=openstack.tfvars -auto-approve"
	howtouse = `
			 Make sure you have terraform installed and in $PATH
			 git clone https://github.com/kubic-project/automation.git
			 
			 cd automation
			 
			 run terraform init in the directories you want to use, for example in caasp-openstack-terraform and/or ses-openstack-terraform

			 put openstack.json in the directories you want to use, for example in caasp-openstack-terraform and/or ses-openstack-terraform
			 
			 openstack.json should look like this:
			 {
				"OSAuthURL":"https://smtg:5000/v3",
				"OSRegionName":"Region",
				"OSProjectName":"caasp",
				"OSUserDomainName":"users",
				"OSIdentityAPIVersion":"3",
				"OSInterface":"public",
				"OSUsername":"user",
				"OSPassword":"pass",
				"OSProjectID":"00000000000000000000000000"
			 }

			 run the utility: mkcaasp -caasp -ses -action destroy -auth openstack.json
			 `
)

var (
	openstack     = flag.String("auth", "openstack.json", "name of the json file containing openstack variables")
	action        = flag.String("action", "apply", "terraform action to run, example: apply, destroy")
	caasp         = flag.Bool("createcaasp", false, "enables/disables caasp terraform openstack setup")
	ses           = flag.Bool("createses", false, "enables/disables ses terraform openstack setup")
	howto         = flag.Bool("usage", false, "prints usage information")
	caasptfoutput = flag.Bool("caasptfoutput", false, "loads in memory caasp terraform ouput json")
	sestfoutput   = flag.Bool("sestfoutput", false, "loads in memory ses terraform ouput json")
)

func main() {
	flag.Parse()
	if *howto {
		fmt.Fprintf(os.Stdout, "%v\n", howtouse)
		os.Exit(0)
	}
	if *caasp {
		utilities.CmdRun("caasp-openstack-terraform", *openstack, fmt.Sprintf(command, *action))
	}
	if *ses {
		utilities.CmdRun("ses-openstack-terraform", *openstack, fmt.Sprintf(command, *action))
	}
	if *caasptfoutput {
		out, _ := utilities.CmdRun("caasp-openstack-terraform", *openstack, "terraform output -json")
		a := utilities.TFOutput{}
		// WIP
		json.Unmarshal([]byte(out), &a)
		fmt.Println(a.IPAdminExternal.Value, a.IPAdminInternal.Value, a.IPMastersExternal.Value, a.IPWorkersExternal.Value)
	}
	if *sestfoutput {
		out, _ := utilities.CmdRun("ses-openstack-terraform", *openstack, "terraform output -json")
		a := utilities.TFOutput{}
		// WIP
		json.Unmarshal([]byte(out), &a)
		fmt.Println(a.IPAdminExternal.Value, a.IPAdminInternal.Value, a.IPMastersExternal.Value, a.IPWorkersExternal.Value)
	}
}
