/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	crypto_rand "crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetKubernetesClient returns a kubernetes api client for us
func GetKubernetesClient() (kubernetes.Interface, error) {
	env := os.Getenv("KUBECONFIG")
	if env == "" {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		return kubernetes.NewForConfig(config)
	} else {
		config, err := clientcmd.LoadFromFile(env)
		if err != nil {
			return nil, err
		}
		overrides := clientcmd.ConfigOverrides{Timeout: "10s"}
		clientConfig, err := clientcmd.NewDefaultClientConfig(*config, &overrides).ClientConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to create API client configuration from kubeconfig: %v", err)
		}

		client, err := kubernetes.NewForConfig(clientConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create API client: %v", err)
		}
		return client, nil
	}
}

// FileExists checks if the file exists
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}

	return true
}

// RandomBytes generates some random bytes
func RandomBytes(length int) (string, error) {
	b := make([]byte, length)
	_, err := crypto_rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
