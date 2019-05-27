package adapters

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/apis/common"
	templatev1 "github.com/openshift/api/template/v1"
)

type RedisAdapter struct {
}

func NewRedisAdapter(options []string) Adapter {
	return NewAppenderAdapter(&RedisAdapter{})
}

func (a *RedisAdapter) Parameters() []templatev1.Parameter {
	return []templatev1.Parameter{
		{
			Name:        "REDIS_IMAGE",
			Description: "Redis image to use",
			Required:    true,
			Value:       "registry.access.redhat.com/rhscl/redis-32-rhel7:3.2",
		},
	}
}

func (r *RedisAdapter) Objects() ([]common.KubernetesObject, error) {
	redisOptions, err := r.options()
	if err != nil {
		return nil, err
	}
	redisComponent := component.NewRedis(redisOptions)
	return redisComponent.Objects(), nil
}

func (r *RedisAdapter) options() (*component.RedisOptions, error) {
	rob := component.RedisOptionsBuilder{}
	rob.AppLabel("${APP_LABEL}")

	return rob.Build()
}
