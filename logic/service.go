package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"k8s-manager/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyv1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceLogic struct {
	client *kubernetes.Clientset
}

func NewServiceLogic(clientset *kubernetes.Clientset) *ServiceLogic {
	return &ServiceLogic{clientset}
}

func (item *ServiceLogic) Add(ctx context.Context, serviceName string, selector map[string]string, port, nodePort int32) error {
	item.client.CoreV1().Services("").Create(ctx, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:     "http",
				Port:     port,
				NodePort: nodePort,
			},
			},
			Selector: selector,
			Type:     corev1.ServiceTypeNodePort,
		},
		Status: corev1.ServiceStatus{},
	}, metav1.CreateOptions{})
	return nil
}

func (item *ServiceLogic) List(ctx context.Context, req *request.ListReq) (*corev1.ServiceList, error) {
	ns := ctx.Value("ns").(string)
	list, err := item.client.CoreV1().Services(ns).List(ctx, metav1.ListOptions{
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
