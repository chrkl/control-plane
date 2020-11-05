package upgrade_kyma

import (
	"time"

	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/broker"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/process"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/runtimeoverrides"

	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/storage"

	"github.com/sirupsen/logrus"
)

type RuntimeOverridesAppender interface {
	Append(input runtimeoverrides.InputAppender, planID, kymaVersion string) error
}

type RuntimeVersionConfiguratorForUpgrade interface {
	ForUpgrade() *internal.RuntimeVersionData
}

type OverridesFromSecretsAndConfigStep struct {
	operationManager       *process.UpgradeKymaOperationManager
	runtimeOverrides       RuntimeOverridesAppender
	runtimeVerConfigurator RuntimeVersionConfiguratorForUpgrade
}

func NewOverridesFromSecretsAndConfigStep(os storage.Operations, runtimeOverrides RuntimeOverridesAppender,
	rvc RuntimeVersionConfiguratorForUpgrade) *OverridesFromSecretsAndConfigStep {
	return &OverridesFromSecretsAndConfigStep{
		operationManager:       process.NewUpgradeKymaOperationManager(os),
		runtimeOverrides:       runtimeOverrides,
		runtimeVerConfigurator: rvc,
	}
}

func (s *OverridesFromSecretsAndConfigStep) Name() string {
	return "Overrides_From_Secrets_And_Config_Step"
}

func (s *OverridesFromSecretsAndConfigStep) Run(operation internal.UpgradeKymaOperation, log logrus.FieldLogger) (internal.UpgradeKymaOperation, time.Duration, error) {
	pp, err := operation.GetProvisioningParameters()
	if err != nil {
		log.Errorf("cannot fetch provisioning parameters from operation: %s", err)
		return s.operationManager.OperationFailed(operation, "invalid operation provisioning parameters")
	}

	planName, exists := broker.PlanNamesMapping[pp.PlanID]
	if !exists {
		log.Errorf("cannot map planID '%s' to planName", pp.PlanID)
		return s.operationManager.OperationFailed(operation, "invalid operation provisioning parameters")
	}

	version := s.getRuntimeVersion(operation)

	if err := s.runtimeOverrides.Append(operation.InputCreator, planName, version.Version); err != nil {
		log.Errorf(err.Error())
		return s.operationManager.RetryOperation(operation, err.Error(), 10*time.Second, 30*time.Minute, log)
	}

	return operation, 0, nil
}

func (s *OverridesFromSecretsAndConfigStep) getRuntimeVersion(operation internal.UpgradeKymaOperation) *internal.RuntimeVersionData {
	// for some previously stored operations the RuntimeVersion property may not be initialized
	if operation.RuntimeVersion.Version != "" {
		return &operation.RuntimeVersion
	}

	// if so, we manually compute the correct version using the same algorithm as when preparing
	// the provisioning operation. The following code can be removed after all operations will use
	// new approach for setting up runtime version in operation struct
	return s.runtimeVerConfigurator.ForUpgrade()
}
