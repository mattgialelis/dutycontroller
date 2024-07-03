package condtions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConditionReason string

const (
	ConditionReasonCreated    ConditionReason = "Created"
	ConditionReasonFailed     ConditionReason = "Error"
	ConditionReasonAssociated ConditionReason = "BusinessServiceAssociated"
)

// GetCondition returns the condition with the provided type.
func GetCondition(conditions []metav1.Condition, conditionType ConditionReason) *metav1.Condition {
	for i := range conditions {
		c := &conditions[i]
		if c.Type == string(conditionType) { // Convert ConditionReason to string
			return c
		}
	}
	return nil
}

// SetCondition sets the specified condition in the slice of conditions.
// If the condition already exists, it updates its status and other fields.
// If not, it creates a new condition and adds it to the slice.
func SetCondition(conditions *[]metav1.Condition, conditionType ConditionReason, status metav1.ConditionStatus, message string) {
	now := metav1.Now() // Get the current time

	existingCondition := GetCondition(*conditions, conditionType) // Get the condition to be modified
	if existingCondition != nil {                                 // Update existing condition
		if existingCondition.Status != status { // Only update if the status changes
			existingCondition.Status = status
			existingCondition.LastTransitionTime = now
		}
		existingCondition.Reason = string(conditionType)
		existingCondition.Message = message
	} else { // Create new condition if it doesn't exist
		condition := metav1.Condition{
			Type:               string(conditionType),
			Status:             status,
			LastTransitionTime: now,
			Reason:             string(conditionType), // Use the condition reason as the reason
			Message:            message,
		}
		*conditions = append(*conditions, condition)
	}
}
