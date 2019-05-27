package adapters

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/apis/common"
	templatev1 "github.com/openshift/api/template/v1"
)

type SystemPostgreSQLImageAdapter struct {
}

func NewSystemPostgreSQLImageAdapter(options []string) Adapter {
	return NewAppenderAdapter(&SystemPostgreSQLImageAdapter{})
}

func (a *SystemPostgreSQLImageAdapter) Parameters() []templatev1.Parameter {
	return []templatev1.Parameter{
		templatev1.Parameter{
			Name:        "SYSTEM_DATABASE_IMAGE",
			Description: "System PostgreSQL image to use",
			Required:    true,
			Value:       "registry.access.redhat.com/rhscl/postgresql-10-rhel7",
		},
	}
}

func (r *SystemPostgreSQLImageAdapter) Objects() ([]common.KubernetesObject, error) {
	systemPostgreSQLOptions, err := r.options()
	if err != nil {
		return nil, err
	}
	systemPostgreSQLComponent := component.NewSystemPostgreSQLImage(systemPostgreSQLOptions)
	return systemPostgreSQLComponent.Objects(), nil
}

func (r *SystemPostgreSQLImageAdapter) options() (*component.SystemPostgreSQLImageOptions, error) {
	aob := component.SystemPostgreSQLImageOptionsBuilder{}
	aob.AppLabel("${APP_LABEL}")
	aob.AmpRelease("${AMP_RELEASE}")
	aob.Image("${SYSTEM_DATABASE_IMAGE}")
	aob.InsecureImportPolicy(component.InsecureImportPolicy)

	return aob.Build()
}
