package rest

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/utils"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/authz"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`
		* run RPC service with `zblcli rest-server --chain-id zblchain`

	TODO: prepare environment automatically
	TODO: provide tests for error cases
*/

func /*Test*/ComplianceDemo_KeepTrackCompliance(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(test_constants.AccountName)

	// Assign Vendor role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.Vendor)

	// Get all compliance infos
	inputComplianceInfos, _ := utils.GetComplianceInfos()

	// Get all certified models
	inputCertifiedModels, _ := utils.GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := utils.GetAllRevokedModels()

	// Publish model info
	modelInfo := utils.NewMsgAddModelInfo(jackKeyInfo.Address)
	_, _ = utils.PublishModelInfo(modelInfo)

	// Check if model either certified or revoked before Compliance record was created
	modelIsCertified, _ := utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, compliance.ZbCertificationType)
	require.False(t, modelIsCertified.Value)

	modelIsRevoked, _ := utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, compliance.ZbCertificationType)
	require.False(t, modelIsRevoked.Value)

	// Assign TestHouse role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.TestHouse)

	// Publish testing result
	testingResult := utils.NewMsgAddTestingResult(modelInfo.VID, modelInfo.PID, jackKeyInfo.Address)
	_, _ = utils.PublishTestingResult(testingResult)

	// Assign ZBCertificationCenter role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.ZBCertificationCenter)

	// Certify model
	certifyModelMsg := compliance.NewMsgCertifyModel(modelInfo.VID, modelInfo.PID, time.Now().UTC(),
		compliance.CertificationType(test_constants.CertificationType), test_constants.EmptyString, jackKeyInfo.Address)
	_, _ = utils.PublishCertifiedModel(certifyModelMsg)

	// Check model is certified
	modelIsCertified, _ = utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.True(t, modelIsCertified.Value)

	modelIsRevoked, _ = utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.False(t, modelIsRevoked.Value)

	// Get all certified models
	certifiedModels, _ := utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total)+1, utils.ParseUint(certifiedModels.Total))

	// Revoke model certification
	revocationTime := certifyModelMsg.CertificationDate.AddDate(0, 0, 1)
	revokeModelMsg := compliance.NewMsgRevokeModel(modelInfo.VID, modelInfo.PID, revocationTime,
		compliance.CertificationType(test_constants.CertificationType), test_constants.RevocationReason, jackKeyInfo.Address)
	_, _ = utils.PublishRevokedModel(revokeModelMsg)

	// Check model is revoked
	modelIsCertified, _ = utils.GetCertifiedModel(modelInfo.VID, modelInfo.PID, revokeModelMsg.CertificationType)
	require.False(t, modelIsCertified.Value)

	modelIsRevoked, _ = utils.GetRevokedModel(modelInfo.VID, modelInfo.PID, revokeModelMsg.CertificationType)
	require.True(t, modelIsRevoked.Value)

	// Get all revoked models
	revokedModels, _ := utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total)+1, utils.ParseUint(revokedModels.Total))

	// Get all certified models
	certifiedModels, _ = utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total), utils.ParseUint(certifiedModels.Total))

	// Get all compliance infos
	complianceInfos, _ := utils.GetComplianceInfos()
	require.Equal(t, utils.ParseUint(inputComplianceInfos.Total)+1, utils.ParseUint(complianceInfos.Total))

	// Get compliance info
	complianceInfo, _ := utils.GetComplianceInfo(modelInfo.VID, modelInfo.PID, certifyModelMsg.CertificationType)
	require.Equal(t, complianceInfo.State, compliance.RevokedState)
	require.Equal(t, 1, len(complianceInfo.History))
	require.Equal(t, complianceInfo.History[0].State, compliance.CertifiedState)
}

func /*Test*/ComplianceDemo_KeepTrackRevocation(t *testing.T) {
	// Get key info for Jack
	jackKeyInfo, _ := utils.GetKeyInfo(test_constants.AccountName)

	// Get all certified models
	inputCertifiedModels, _ := utils.GetAllCertifiedModels()

	// Get all revoked models
	inputRevokedModels, _ := utils.GetAllRevokedModels()

	// Assign ZBCertificationCenter role to Jack
	utils.AssignRole(jackKeyInfo.Address, jackKeyInfo, authz.ZBCertificationCenter)

	vid, pid := test_constants.VID, test_constants.PID

	// Revoke model
	revocationTime := time.Now().UTC()
	revokeModelMsg := compliance.NewMsgRevokeModel(vid, pid, revocationTime,
		compliance.CertificationType(test_constants.CertificationType), test_constants.RevocationReason, jackKeyInfo.Address)
	_, _ = utils.PublishRevokedModel(revokeModelMsg)

	// Check model is revoked
	modelIsRevoked, _ := utils.GetRevokedModel(revokeModelMsg.VID, revokeModelMsg.PID, revokeModelMsg.CertificationType)
	require.True(t, modelIsRevoked.Value)

	modelIsCertified, _ := utils.GetCertifiedModel(revokeModelMsg.VID, revokeModelMsg.PID, revokeModelMsg.CertificationType)
	require.False(t, modelIsCertified.Value)

	// Get all revoked models
	revokedModels, _ := utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total)+1, utils.ParseUint(revokedModels.Total))

	// Certify model
	certificationTime := revocationTime.AddDate(0, 0, 1)
	certifyModelMsg := compliance.NewMsgCertifyModel(vid, pid, certificationTime,
		compliance.CertificationType(test_constants.CertificationType), test_constants.EmptyString, jackKeyInfo.Address)
	_, _ = utils.PublishCertifiedModel(certifyModelMsg)

	// Check model is certified
	modelIsRevoked, _ = utils.GetRevokedModel(certifyModelMsg.VID, certifyModelMsg.PID, certifyModelMsg.CertificationType)
	require.False(t, modelIsRevoked.Value)

	modelIsCertified, _ = utils.GetCertifiedModel(certifyModelMsg.VID, certifyModelMsg.PID, certifyModelMsg.CertificationType)
	require.True(t, modelIsCertified.Value)

	// Get all certified models
	certifiedModels, _ := utils.GetAllCertifiedModels()
	require.Equal(t, utils.ParseUint(inputCertifiedModels.Total)+1, utils.ParseUint(certifiedModels.Total))

	// Get all revoked models
	revokedModels, _ = utils.GetAllRevokedModels()
	require.Equal(t, utils.ParseUint(inputRevokedModels.Total), utils.ParseUint(revokedModels.Total))
}
