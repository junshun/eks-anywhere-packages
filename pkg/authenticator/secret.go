package authenticator

import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

type secret struct {
	config *rest.Config
}

var _ Authenticator = (*secret)(nil)

func NewSecret(config *rest.Config) *secret {
	return &secret{config: config}
}

func (s *secret) GetAuthFile() (string, error) {
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

func (s *secret) getSecretToken() (string, error){
	clientset, err := kubernetes.NewForConfig(s.config)
	if err != nil {
		return "", fmt.Errorf("Getting new clientset")
	}
	secret, err := clientset.CoreV1().Secrets("eksa-packages").Get(context.Background(), "docker-secret", metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	//TODO Unhard code name of data
	dockerconfig := secret.Data["access_key"]

	return string(dockerconfig), nil
}

func createAuthFile(token string) (string, error){
	// TODO Remove hardcoded account ID/region
	dockerStruct := &DockerAuth{
		Auths: map[string]DockerAuthRegistry{
			fmt.Sprintf("563489572116.dkr.ecr.us-west-2.amazonaws.com"): DockerAuthRegistry{token},
		},
	}
	jsonbytes, err := json.Marshal(*dockerStruct)
	if err != nil {
		return "", fmt.Errorf("Marshalling docker auth file to json %w", err)
	}
	f, err := os.CreateTemp("", "dockerAuth")
	if err != nil {
		return "", fmt.Errorf("Creating tempfile %w", err)
	}
	defer f.Close()
	fmt.Fprint(f, string(jsonbytes))
	return f.Name(), nil
}