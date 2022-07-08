package logic

import (
	"context"
	"github.com/pkg/errors"
	"k8s-manager/local"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type IngressLogic struct {
	client *kubernetes.Clientset
}

func NewIngressLogic(clientset *kubernetes.Clientset) *IngressLogic {
	return &IngressLogic{clientset}
}

func (item *IngressLogic) Add(ctx context.Context, name string) error {
	_, err := item.client.NetworkingV1().Ingresses(local.GetNamespace(ctx)).Create(ctx, &v1.Ingress{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.IngressSpec{},
		Status:     v1.IngressStatus{},
	}, metav1.CreateOptions{})
	return err
}

func (item *IngressLogic) List(ctx context.Context) (*v1.IngressList, error) {
	list, err := item.client.NetworkingV1().Ingresses(local.GetNamespace(ctx)).List(ctx, metav1.ListOptions{})
	return list, errors.WithStack(err)
}
