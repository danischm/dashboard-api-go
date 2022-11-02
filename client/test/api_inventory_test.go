/*
Meraki Dashboard API

Testing InventoryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_client_InventoryApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test InventoryApiService ClaimIntoOrganizationInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.ClaimIntoOrganizationInventory(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryDevice", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string
        var serial string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService GetOrganizationInventoryDevices", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.GetOrganizationInventoryDevices(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test InventoryApiService ReleaseFromOrganizationInventory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var organizationId string

        resp, httpRes, err := apiClient.InventoryApi.ReleaseFromOrganizationInventory(context.Background(), organizationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}