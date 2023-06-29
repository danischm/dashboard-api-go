/*
Meraki Dashboard API

Testing QosRulesApiService

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

func Test_client_QosRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QosRulesApiService CreateNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.QosRulesApi.CreateNetworkSwitchQosRule(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService DeleteNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		httpRes, err := apiClient.QosRulesApi.DeleteNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService GetNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		resp, httpRes, err := apiClient.QosRulesApi.GetNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService GetNetworkSwitchQosRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.QosRulesApi.GetNetworkSwitchQosRules(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService GetNetworkSwitchQosRulesOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.QosRulesApi.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService UpdateNetworkSwitchQosRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var qosRuleId string

		resp, httpRes, err := apiClient.QosRulesApi.UpdateNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QosRulesApiService UpdateNetworkSwitchQosRulesOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.QosRulesApi.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}