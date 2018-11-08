package utilities

// OpenStackAPI holds openstack API variables
type OpenStackAPI struct {
	OSAuthURL            string
	OSRegionName         string
	OSProjectName        string
	OSUserDomainName     string
	OSIdentityAPIVersion string
	OSInterface          string
	OSUsername           string
	OSPassword           string
	OSProjectID          string
}

// EnvOS holds as slice with openstack API variables
type EnvOS []string

// TFOutput is holding terraform output json variables
type TFOutput struct {
	IPAdminExternal   *Admin   `json:"ip_admin_external"`
	IPAdminInternal   *Admin   `json:"ip_admin_internal"`
	IPMastersExternal Machines `json:"ip_masters"`
	IPWorkersExternal Machines `json:"ip_workers"`
}

type Admin struct {
	Value string
}

type Machines struct {
	Value []string
}
