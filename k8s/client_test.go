package k8s

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNewK8sClientSet(t *testing.T) {
	client := lo.Must(NewK8sClientSet())

	list := lo.Must(client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{Limit: 3}))

	fmt.Println(list.Continue)
	lo.ForEach(list.Items, func(t v1.Pod, i int) {
		fmt.Println(t.Name)
	})

	list = lo.Must(client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{Limit: 3, Continue: list.Continue}))

	fmt.Println(list.Continue)
	lo.ForEach(list.Items, func(t v1.Pod, i int) {
		fmt.Println(t.Name)
	})
}
