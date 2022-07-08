package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"k8s-manager/local"
	"k8s-manager/request"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"strings"
	"time"
)

type DeploymentLogic struct {
	client *kubernetes.Clientset
}

func NewDeploymentLogic(clientset *kubernetes.Clientset) *DeploymentLogic {
	return &DeploymentLogic{clientset}
}

func (item *DeploymentLogic) Add(ctx context.Context, req *request.DeploymentAddReq) error {
	// 实例化一个数据结构
	labels := strings.Split(req.Labels, ";")
	specLabels := map[string]string{}
	for _, label := range labels {
		split := strings.Split(label, "=")
		if len(split) == 2 {
			specLabels[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
		}

	}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &req.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: specLabels,
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: specLabels,
				},
				Spec: corev1.PodSpec{
					Containers: req.Containers,
				},
			},
		},
	}

	_, err := item.client.AppsV1().Deployments(req.Namespace).Create(ctx, deployment, metav1.CreateOptions{})
	return errors.WithStack(err)
}

func (item *DeploymentLogic) List(ctx context.Context, req *request.ListReq) (*appsv1.DeploymentList, error) {
	list, err := item.client.AppsV1().Deployments(local.GetNamespace(ctx)).List(ctx, metav1.ListOptions{
		LabelSelector: req.Label,
		FieldSelector: req.Field,
		Limit:         req.Limit,
		Continue:      req.Continue,
	})
	return list, errors.WithStack(err)
}

func (item *DeploymentLogic) Delete(ctx context.Context, req *request.DeleteReq) error {
	err := item.client.AppsV1().Deployments(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return errors.WithStack(err)
}

func (item *DeploymentLogic) Update(ctx context.Context, req *request.ApplyReq) error {
	conf := &appsv1.Deployment{}
	err := json.Unmarshal([]byte(req.Content), conf)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = item.client.AppsV1().Deployments(req.Namespace).Update(ctx, conf, metav1.UpdateOptions{})
	return errors.WithStack(err)
}

func (item *DeploymentLogic) Restart(ctx context.Context, req *request.DeploymentRestartReq) error {
	patchData := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"annotations": map[string]interface{}{
						"www.example.com/restartedAt": time.Now().Format(time.Stamp),
					},
				},
			},
		},
	}
	data, err := json.Marshal(patchData)
	if err != nil {
		return errors.WithStack(err)
	}
	res, err := item.client.AppsV1().Deployments(req.Namespace).Patch(ctx, req.Name, types.MergePatchType, data, metav1.PatchOptions{})
	fmt.Println(res)
	return errors.WithStack(err)
}
