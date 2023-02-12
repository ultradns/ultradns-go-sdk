package integration_test

import (
	"testing"

	"github.com/ultradns/ultradns-go-sdk/internal/testing/integration"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
)

var IPID = ""

func (t *IntegrationTest) TestDirGroupIPResources(ipData *ip.DirGroupIP, ipName string) {
	it := IntegrationTest{}

	t.Test.Run("TestCreatetDirGroupIP",
		func(st *testing.T) {
			it.Test = st
			it.CreatetDirGroupIP(ipData, ipName)
		})
	t.Test.Run("TestListDirGroupIPResource",
		func(st *testing.T) {
			it.Test = st
			it.ListDirGroupIP(ipData, ipName)
		})
	t.Test.Run("TestUpdateDirGroupIP",
		func(st *testing.T) {
			it.Test = st
			it.UpdateDirGroupIP(ipData, ipName)
		})
	t.Test.Run("TestPartialUpdateDirGroupIP",
		func(st *testing.T) {
			it.Test = st
			it.PartialUpdateDirGroupIP(ipData, ipName)
		})
	t.Test.Run("TestReadDirGroupIP",
		func(st *testing.T) {
			it.Test = st
			it.ReadDirGroupIP(ipData, ipName)
		})
	t.Test.Run("TestDeleteDirGroupIP",
		func(st *testing.T) {
			it.Test = st
			it.DeleteDirGroupIP(ipData, ipName)
		})
}

func (t *IntegrationTest) CreatetDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := ipService.Create(ipdata); err != nil {
		t.Test.Fatal(err)

	}
}

func (t *IntegrationTest) ListDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	_, _, err = ipService.List(nil, ipdata)

	if err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) UpdateDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := ipService.Update(ipdata); err != nil {
		t.Test.Fatal(err)
	}

}

func (t *IntegrationTest) PartialUpdateDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := ipService.PartialUpdate(ipdata); err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) ReadDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, _, _, err := ipService.Read(ipdata.Name + ":" + ipData.AccountName); err != nil {
		t.Test.Fatal(err)
	}
}

func (t *IntegrationTest) DeleteDirGroupIP(ipData *ip.DirGroupIP, ipName string) {
	ipdata := getDirGroupIP(ipName)
	ipService, err := ip.Get(integration.TestClient)

	if err != nil {
		t.Test.Fatal(err)
	}

	if _, err := ipService.Delete(ipdata.Name + ":" + ipData.AccountName); err != nil {
		t.Test.Fatal(err)
	}
}

func getDirGroupIP(ipName string) *ip.DirGroupIP {
	ipdata := &ip.DirGroupIP{
		Name:        ipName,
		AccountName: integration.TestAccount,
		IPs: []*ip.IPAddress{
			&ip.IPAddress{
				Start: "192.168.1.1",
				End:   "192.168.1.10",
			},
			&ip.IPAddress{
				Cidr: "192.168.2.0/24",
			},
			&ip.IPAddress{
				Address: "192.168.3.4",
			},
		},
		Description: "Description of GEO",
	}

	return ipdata
}
