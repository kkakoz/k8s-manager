package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"io"
	"k8s-manager/k8s"
	"k8s-manager/pkg/mdctx"
	"k8s-manager/pkg/ws"
	"k8s-manager/request"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

type PodLogic struct {
	client *kubernetes.Clientset
}

func NewPodLogic(clientset *kubernetes.Clientset) *PodLogic {
	return &PodLogic{clientset}
}

func (item *PodLogic) Add(ctx context.Context, req *request.PodAddReq) error {
	containers := lo.Map(req.Containers, func(v request.Container, i int) corev1.Container {
		return corev1.Container{
			Name:  v.Name,
			Image: v.Image,
			Ports: []corev1.ContainerPort{{ContainerPort: v.ContainerPort, HostPort: v.HostPort}},
		}
	})
	_, err := item.client.CoreV1().Pods(req.Namespace).Create(ctx, &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   req.Name,
			Labels: map[string]string{},
		},
		Spec:   corev1.PodSpec{Containers: containers},
		Status: corev1.PodStatus{},
	}, metav1.CreateOptions{})
	return errors.WithStack(err)
}

func (item *PodLogic) List(ctx context.Context, req *request.ListReq) (*corev1.PodList, error) {
	list, err := item.client.CoreV1().Pods(mdctx.GetNs(ctx)).List(ctx, metav1.ListOptions{
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

func (item *PodLogic) Update(ctx context.Context, req *request.ApplyReq) error {
	conf := &corev1.Pod{}
	err := json.Unmarshal([]byte(req.Content), conf)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = item.client.CoreV1().Pods(req.Namespace).Update(ctx, conf, metav1.UpdateOptions{})
	return errors.WithStack(err)
}

var defaultTailLen int64 = 100

func (item *PodLogic) GetLog(ctx context.Context, req *request.PodLogReq) (string, error) {

	logs, err := item.client.CoreV1().Pods(req.Namespace).GetLogs(req.Name, &corev1.PodLogOptions{
		Container: req.ContainerName,
		TailLines: lo.Ternary(req.TailLien == 0, &defaultTailLen, &req.TailLien),
	}).Stream(ctx)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer logs.Close()
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, logs)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return buf.String(), err
}

func (item *PodLogic) Terminal(ctx context.Context, c echo.Context, req *request.PodTerminalReq) error {
	sshReq := item.client.CoreV1().RESTClient().Post().Resource("pods").Name(req.Name).Namespace(req.Namespace).
		SubResource("exec").VersionedParams(&corev1.PodExecOptions{
		Container: req.ContainerName,
		Command:   []string{"bash"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)
	wsConn, err := ws.NewConn(ctx, req.Name, c)
	defer wsConn.Close()

	exector, err := remotecommand.NewSPDYExecutor(k8s.GetRestConf(), "POST", sshReq.URL())
	if err != nil {
		return err
	}
	err = exector.Stream(remotecommand.StreamOptions{
		Stdin:             wsConn,
		Stdout:            wsConn,
		Stderr:            wsConn,
		Tty:               true,
		TerminalSizeQueue: wsConn,
	})
	if err != nil {
		return err
	}
	return nil
}
