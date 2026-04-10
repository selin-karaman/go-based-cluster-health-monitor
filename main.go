package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	fmt.Println("[KubeCheck] Starting Cluster Health Monitor...")

	var config *rest.Config
	var err error

	config, err = rest.InClusterConfig()
	if err != nil {
		fmt.Println("Running outside cluster, looking for kubeconfig...")

		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "path to kubeconfig")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "path to kubeconfig")
		}
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			fmt.Printf("Fatal Error: Could not load any config: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Running inside Kubernetes cluster! Using ServiceAccount.")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error building clientset: %v\n", err)
		os.Exit(1)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing pods: %v\n", err)
		fmt.Println("Check if your RBAC (Role/ServiceAccount) permissions are correct.")
		os.Exit(1)
	}

	fmt.Printf("\n CLUSTER HEALTH REPORT\n")
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', 0)
	fmt.Fprintln(w, "NAMESPACE\tNAME\tSTATUS\tHEALTH")
	fmt.Fprintln(w, "---------\t----\t------\t------")

	unhealthyCount := 0
	for _, pod := range pods.Items {
		status := string(pod.Status.Phase)
		health := "Healthy"

		if status != "Running" && status != "Succeeded" {
			health = "Unhealthy"
			unhealthyCount++
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", pod.Namespace, pod.Name, status, health)
	}
	w.Flush()

	if unhealthyCount == 0 {
		fmt.Println("\n Cluster Status: PERFECT")
	} else {
		fmt.Printf("\n Cluster Status: ATTENTION REQUIRED (%d issues found)\n", unhealthyCount)
	}
}
