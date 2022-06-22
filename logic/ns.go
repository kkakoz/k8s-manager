package logic

import (
	"context"
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NsLogic struct {
	client *kubernetes.Clientset
}

func NewNsLogic(clientset *kubernetes.Clientset) *NsLogic {
	return &NsLogic{clientset}
}

func (item *NsLogic) Add(ctx context.Context, name string) error {
	_, err := item.client.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}, metav1.CreateOptions{})
	return err
}

func (item *NsLogic) List(ctx context.Context) (*v1.NamespaceList, error) {
	list, err := item.client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	return list, errors.WithStack(err)
}
