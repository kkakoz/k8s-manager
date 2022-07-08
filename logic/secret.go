package logic

import (
	"context"
	"github.com/pkg/errors"
	"k8s-manager/local"
	"k8s-manager/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type SecretLogic struct {
	client *kubernetes.Clientset
}

func NewSecretLogic(clientset *kubernetes.Clientset) *SecretLogic {
	return &SecretLogic{clientset}
}

func (item *SecretLogic) Add(ctx context.Context, req *request.SecretAddReq) error {
	_, err := item.client.CoreV1().Secrets(req.Namespace).Create(ctx, &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{},
		StringData: req.Data,
		Type:       "",
	}, metav1.CreateOptions{})
	return err
}

func (item *SecretLogic) List(ctx context.Context) (*corev1.SecretList, error) {
	list, err := item.client.CoreV1().Secrets(local.GetNamespace(ctx)).List(ctx, metav1.ListOptions{})
	return list, errors.WithStack(err)
}
