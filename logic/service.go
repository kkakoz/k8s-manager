package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyv1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

type ServiceLogic struct {
	client *kubernetes.Clientset
}

func NewServiceLogic(clientset *kubernetes.Clientset) *ServiceLogic {
	return &ServiceLogic{clientset}
}

func (item *ServiceLogic) Add(ctx context.Context, req *request.ServiceAddReq) error {
	selectors := strings.Split(req.Selector, ";")
	specSelector := map[string]string{}
	for _, selector := range selectors {
		split := strings.Split(selector, "=")
		if len(split) == 2 {
			specSelector[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		}
	}
	_, err := item.client.CoreV1().Services("").Create(ctx, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:     "http",
				Port:     req.Port,
				NodePort: req.NodePort,
			},
			},
			Selector: specSelector,
			Type:     lo.Ternary(req.Type == 1, corev1.ServiceTypeClusterIP, corev1.ServiceTypeNodePort),
		},
		Status: corev1.ServiceStatus{},
	}, metav1.CreateOptions{})
	return errors.WithStack(err)
}

func (item *ServiceLogic) List(ctx context.Context, req *request.ListReq) (*corev1.ServiceList, error) {
	list, err := item.client.CoreV1().Services(mdctx.GetNs(ctx)).List(ctx, metav1.ListOptions{
		LabelSelector: req.Label,
		FieldSelector: req.Field,
		Limit:         req.Limit,
		Continue:      req.Continue,
	})
	return list, errors.WithStack(err)
}

func (item *ServiceLogic) Apply(ctx context.Context, req *request.ApplyReq) error {
	conf := &applyv1.ServiceApplyConfiguration{}
	err := json.Unmarshal([]byte(req.Content), conf)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = item.client.CoreV1().Services(req.Namespace).Apply(ctx, conf, metav1.ApplyOptions{})
	return errors.WithStack(err)
}
