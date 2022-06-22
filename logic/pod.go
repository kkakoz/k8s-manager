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

type PodLogic struct {
	client *kubernetes.Clientset
}

func NewPodLogic(clientset *kubernetes.Clientset) *PodLogic {
	return &PodLogic{clientset}
}

func (item *PodLogic) Add(ctx context.Context, req *request.PodAddReq) error {
	_, err := item.client.CoreV1().Pods(req.Namespace).Create(ctx, &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   req.Name,
			Labels: map[string]string{},
		},
		Spec:   corev1.PodSpec{Containers: req.Containers},
		Status: corev1.PodStatus{},
	}, metav1.CreateOptions{})
	return errors.WithStack(err)
}

func (item *PodLogic) List(ctx context.Context, req *request.ListReq) (*corev1.PodList, error) {
	list, err := item.client.CoreV1().Pods(req.Namespace).List(ctx, metav1.ListOptions{
		LabelSelector: req.Label,
		FieldSelector: req.Field,
		Limit:         req.Limit,
		Continue:      req.Continue,
	})
	return list, errors.WithStack(err)
}

func (item *PodLogic) Delete(ctx context.Context, req *request.DeleteReq) error {
	err := item.client.CoreV1().Pods(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return errors.WithStack(err)
}

func (item *PodLogic) Apply(ctx context.Context, req *request.ApplyReq) error {
	conf := &applyv1.PodApplyConfiguration{}
	err := json.Unmarshal([]byte(req.Content), conf)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = item.client.CoreV1().Pods(req.Namespace).Apply(ctx, conf, metav1.ApplyOptions{})
	return errors.WithStack(err)
}
