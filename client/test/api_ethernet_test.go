/*
Meraki Dashboard API

Testing EthernetApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/meraki/dashboard-api-go/client"
)

func Test_client_EthernetApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EthernetApiService AssignNetworkWirelessEthernetPortsProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.EthernetApi.AssignNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService CreateNetworkWirelessEthernetPortsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.EthernetApi.CreateNetworkWirelessEthernetPortsProfile(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService DeleteNetworkWirelessEthernetPortsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var profileId string

		httpRes, err := apiClient.EthernetApi.DeleteNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService GetNetworkWirelessEthernetPortsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var profileId string

		resp, httpRes, err := apiClient.EthernetApi.GetNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService GetNetworkWirelessEthernetPortsProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.EthernetApi.GetNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService GetOrganizationWirelessDevicesEthernetStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.EthernetApi.GetOrganizationWirelessDevicesEthernetStatuses(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService SetNetworkWirelessEthernetPortsProfilesDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.EthernetApi.SetNetworkWirelessEthernetPortsProfilesDefault(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthernetApiService UpdateNetworkWirelessEthernetPortsProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var profileId string

		resp, httpRes, err := apiClient.EthernetApi.UpdateNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
