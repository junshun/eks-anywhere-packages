package authenticator

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

type dockersecret struct {
	config *rest.Config
}

var _ Authenticator = (*dockersecret)(nil)

func NewDockerSecret(config *rest.Config) *dockersecret {
	return &dockersecret{config: config}
}

func (s *dockersecret) GetAuthFile() (string, error) {
	token, err := s.getSecretToken()
	if err != nil {
		return "", fmt.Errorf("Failed to get authfile %s\n", err)
	}
	authfile, err := createAuthFile(token)
	if err != nil {
		return "", fmt.Errorf("Failed to get authfile %s\n", err)
	}

	return authfile, nil
}

func (s *dockersecret) getSecretToken() ([]uint8, error) {
	clientset, err := kubernetes.NewForConfig(s.config)
	if err != nil {
		return nil, fmt.Errorf("Getting new clientset")
	}
	//secret, err := clientset.CoreV1().Secrets("eksa-packages").Get(context.Background(), "docker-secret", metav1.GetOptions{})
	dockersecret, err := clientset.CoreV1().Secrets("eksa-packages").Get(context.Background(), "dockersecret", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	dockerconfig := dockersecret.Data[".dockerconfigjson"]

	return dockerconfig, nil
}

func createAuthFile(data []uint8) (string, error) {
	f, err := os.CreateTemp("", "dockerAuth")
	if err != nil {
		return "", fmt.Errorf("Creating tempfile %w", err)
	}
	defer f.Close()
	fmt.Fprint(f, string(data))
	return f.Name(), nil
}
