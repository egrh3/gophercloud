package users

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List lists the existing users.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, rootURL(client), func(r pagination.PageResult) pagination.Page {
		return UserPage{pagination.SinglePageBase(r)}
	})
}

// CommonOpts are the parameters that are shared between CreateOpts and
// UpdateOpts
type CommonOpts struct {
	// Either a name or username is required. When provided, the value must be
	// unique or a 409 conflict error will be returned. If you provide a name but
	// omit a username, the latter will be set to the former; and vice versa.
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	// The ID of the tenant to which you want to assign this user.
	TenantID string `json:"tenant_id,omitempty"`
	// Indicates whether this user is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
	// The email address of this user.
	Email string `json:"email,omitempty"`
}

// CreateOpts represents the options needed when creating new users.
type CreateOpts CommonOpts

// CreateOptsBuilder describes struct types that can be accepted by the Create call.
type CreateOptsBuilder interface {
	ToUserCreateMap() (map[string]interface{}, error)
}

// ToUserCreateMap assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToUserCreateMap() (map[string]interface{}, error) {
	if opts.Name == "" && opts.Username == "" {
		err := gophercloud.ErrMissingInput{}
		err.Argument = "users.CreateOpts.Name/users.CreateOpts.Username"
		err.Info = "Either a Name or Username must be provided"
		return nil, err
	}
	return gophercloud.BuildRequestBody(opts, "user")
}

// Create is the operation responsible for creating new users.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToUserCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(rootURL(client), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201},
	})

	return res
}

// Get requests details on a single user, either by ID.
func Get(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(ResourceURL(client, id), &result.Body, nil)
	return result
}

// UpdateOptsBuilder allows extensions to add additional attributes to the Update request.
type UpdateOptsBuilder interface {
	ToUserUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the base attributes that may be updated on an existing server.
type UpdateOpts CommonOpts

// ToUserUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToUserUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "user")
}

// Update is the operation responsible for updating exist users by their UUID.
func Update(client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) UpdateResult {
	var r UpdateResult
	b, err := opts.ToUserUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	_, r.Err = client.Put(ResourceURL(client, id), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return r
}

// Delete is the operation responsible for permanently deleting an API user.
func Delete(client *gophercloud.ServiceClient, id string) DeleteResult {
	var r DeleteResult
	_, r.Err = client.Delete(ResourceURL(client, id), nil)
	return r
}

// ListRoles lists the existing roles that can be assigned to users.
func ListRoles(client *gophercloud.ServiceClient, tenantID, userID string) pagination.Pager {
	return pagination.NewPager(client, listRolesURL(client, tenantID, userID), func(r pagination.PageResult) pagination.Page {
		return RolePage{pagination.SinglePageBase(r)}
	})
}
