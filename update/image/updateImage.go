package update

import (
	"context"
	"flag"
	"fmt"
	"kubeclient/kubelog"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
)

type updateImage struct {
	clientset *kubernetes.Clientset
}

func NewUpdateImage() *updateImage {
	var kubeconfig *string
	updateImage := new(updateImage)
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join("config/kubeconfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	updateImage.clientset = clientset
	return updateImage
}

// 指定版本更新镜像
func (update *updateImage) UpdateDeploymentsImage(namespace string, podName string, containerName, image string) error {
	podClient := update.clientset.AppsV1().Deployments(namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("获取deployment的最新版本失败: %v", getErr))
		}
		containers := result.Spec.Template.Spec.Containers
		for i, container := range containers {
			if container.Name == containerName {
				info := fmt.Sprintf("发起更新%s/%s/%s| image=%s => image=%s", namespace, podName, containerName, result.Spec.Template.Spec.Containers[i].Image, image)
				kubelog.Kinfo(info)
				result.Spec.Template.Spec.Containers[i].Image = image
				_, updateErr := podClient.Update(context.TODO(), result, metav1.UpdateOptions{})
				return updateErr
			}
		}
		return fmt.Errorf("没有更新任何镜像,namespace=%s,pod=%s中,没有找到名称为%s的容器", namespace, podName, containerName)
	})
	return retryErr
}

func (update *updateImage) UpdateStatefulSetImage(namespace string, podName string, containerName, image string) error {
	podClient := update.clientset.AppsV1().StatefulSets(namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("获取StatefulSet的最新版本失败: %v", getErr))
		}
		containers := result.Spec.Template.Spec.Containers
		for i, container := range containers {
			if container.Name == containerName {
				info := fmt.Sprintf("发起更新%s/%s/%s| image=%s => image=%s", namespace, podName, containerName, result.Spec.Template.Spec.Containers[i].Image, image)
				kubelog.Kinfo(info)
				result.Spec.Template.Spec.Containers[i].Image = image
				_, updateErr := podClient.Update(context.TODO(), result, metav1.UpdateOptions{})

				return updateErr
			}
		}
		return fmt.Errorf("没有更新任何镜像,namespace=%s,pod=%s中,没有找到名称为%s的容器", namespace, podName, containerName)
	})
	return retryErr
}

func (update *updateImage) UpdateDaemonSetSetImage(namespace string, podName string, containerName, image string) error {
	podClient := update.clientset.AppsV1().DaemonSets(namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("获取DaemonSet的最新版本失败: %v", getErr))
		}
		containers := result.Spec.Template.Spec.Containers
		for i, container := range containers {
			if container.Name == containerName {
				info := fmt.Sprintf("发起更新%s/%s/%s| image=%s => image=%s", namespace, podName, containerName, result.Spec.Template.Spec.Containers[i].Image, image)
				kubelog.Kinfo(info)
				result.Spec.Template.Spec.Containers[i].Image = image
				_, updateErr := podClient.Update(context.TODO(), result, metav1.UpdateOptions{})
				return updateErr
			}
		}
		return fmt.Errorf("没有更新任何镜像,namespace=%s,pod=%s中,没有找到名称为%s的容器", namespace, podName, containerName)
	})
	return retryErr
}

func (update *updateImage) UpdateReplicaSetSetImage(namespace string, podName string, containerName, image string) error {
	podClient := update.clientset.AppsV1().ReplicaSets(namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("获取ReplicaSets的最新版本失败: %v", getErr))
		}
		containers := result.Spec.Template.Spec.Containers
		for i, container := range containers {
			if container.Name == containerName {
				info := fmt.Sprintf("发起更新%s/%s/%s| image=%s => image=%s", namespace, podName, containerName, result.Spec.Template.Spec.Containers[i].Image, image)
				kubelog.Kinfo(info)
				result.Spec.Template.Spec.Containers[i].Image = image
				_, updateErr := podClient.Update(context.TODO(), result, metav1.UpdateOptions{})
				return updateErr
			}
		}
		return fmt.Errorf("没有更新任何镜像,namespace=%s,pod=%s中,没有找到名称为%s的容器", namespace, podName, containerName)
	})
	return retryErr
}
