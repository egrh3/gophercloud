package servergroups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List returns a Pager that allows you to iterate over a collection of ServerGroups.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return ServerGroupPage{pagination.SinglePageBase(r)}
	})
}

// CreateOptsBuilder describes struct types that can be accepted by the Create call. Notably, the
// CreateOpts struct in this package does.
type CreateOptsBuilder interface {
	ToServerGroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies a Server Group allocation request
type CreateOpts struct {
	// Name is the name of the server group
	Name string `json:"name" required:"true"`
	// Policies are the server group policies
	Policies []string `json:"policies" required:"true"`
}

// ToServerGroupCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToServerGroupCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "server_group")
}

// Create requests the creation of a new Server Group
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var r CreateResult
	b, err := opts.ToServerGroupCreateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return r
}

// Get returns data about a previously created ServerGroup.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var r GetResult
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return r
}

// Delete requests the deletion of a previously allocated ServerGroup.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var r DeleteResult
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return r
}
