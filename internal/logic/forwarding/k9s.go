package forwarding

import (
	"context"
	"fmt"
	"github.com/rollicks-c/kgate/internal/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
)

func createClient(target config.Target) (*kubernetes.Clientset, *rest.Config, error) {

	// load kube config
	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: target.K8sConfigFile}
	configOverrides := &clientcmd.ConfigOverrides{
		CurrentContext: target.K8sContext,
	}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, nil, err
	}

	// create client
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	return clientSet, clientConfig, nil
}

func getPodForService(clientSet *kubernetes.Clientset, serviceName, namespace string) (string, error) {

	// gather all pods in NS
	pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list pods: %v", err)
	}

	// find pod for service
	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, serviceName) {
			return pod.Name, nil
		}
	}

	return "", fmt.Errorf("no pod found for service %s", serviceName)
}
