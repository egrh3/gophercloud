package floatingips

import (
	"testing"

	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	count := 0
	err := List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractFloatingIPs(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedFloatingIPsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)

	actual, err := Create(client.ServiceClient(), CreateOpts{
		Pool: "nova",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedFloatingIP, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := Get(client.ServiceClient(), "2").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecondFloatingIP, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)

	err := Delete(client.ServiceClient(), "1").ExtractErr()
	th.AssertNoErr(t, err)
}

func TestAssociate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssociateSuccessfully(t)

	associateOpts := AssociateOpts{
		FloatingIP: "10.10.10.2",
	}

	err := AssociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", associateOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestAssociateFixed(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAssociateFixedSuccessfully(t)

	associateOpts := AssociateOpts{
		FloatingIP: "10.10.10.2",
		FixedIP:    "166.78.185.201",
	}

	err := AssociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", associateOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestDisassociateInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDisassociateSuccessfully(t)

	disassociateOpts := DisassociateOpts{
		FloatingIP: "10.10.10.2",
	}

	err := DisassociateInstance(client.ServiceClient(), "4d8c3732-a248-40ed-bebc-539a6ffd25c0", disassociateOpts).ExtractErr()
	th.AssertNoErr(t, err)
}
