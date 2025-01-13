package k8s

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sHelper struct {
	clientset *kubernetes.Clientset
}

// Kubernetes Client'ını oluşturmak için fonksiyon
func (k *K8sHelper) GetK8sClient() (*kubernetes.Clientset, error) {
	// Kubernetes cluster içinde çalışırken in-cluster config'i yükle
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get in-cluster config: %v", err)
	}

	// Kubernetes clientset oluştur
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes clientset: %v", err)
	}

	return clientset, nil
}

// Node bilgilerini almak için fonksiyon
func (k *K8sHelper) GetNodeInfo(nodeName string) (*corev1.Node, error) {
	clientset, err := k.GetK8sClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get Kubernetes clientset: %v", err)
	}

	// Node bilgilerini Kubernetes API'sinden al
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get node information for %s: %v", nodeName, err)
	}

	return node, nil
}
