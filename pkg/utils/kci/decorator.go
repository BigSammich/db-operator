package kci

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMapBuilder builds kubernetes configmap object
func ConfigMapBuilder(name string, namespace string, data map[string]string) *corev1.ConfigMap {
	configmap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    BaseLabelBuilder(),
		},
		Data: data,
	}

	return configmap
}

// SecretBuilder builds kubernetes secret object
func SecretBuilder(secretName string, namespace string, data map[string][]byte) *corev1.Secret {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
			Labels:    BaseLabelBuilder(),
		},
		Data: data,
	}

	return secret
}

// BuildEnvVarSource builds kubernetes a source for the value of an EnvVar
func BuildEnvVarSource(fieldPath string) *corev1.EnvVarSource {

	return &corev1.EnvVarSource{
		FieldRef: &corev1.ObjectFieldSelector{FieldPath: fieldPath},
	}
}

// BaseLabelBuilder builds source label.
// It will be used as base label for the kubernetes objects which created by db-operator
func BaseLabelBuilder() map[string]string {
	return map[string]string{
		"created-by": "db-operator",
	}
}

// LabelBuilder builds key, value label which can be used for kubernetes object metadata
func LabelBuilder(labels map[string]string) map[string]string {
	for k, v := range BaseLabelBuilder() {
		labels[k] = v
	}

	return labels
}