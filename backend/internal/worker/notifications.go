package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/adapter/sms"
)

// NotificationType defines the type of notification to send.
type NotificationType string

const (
	NotificationNewLeadAssignment NotificationType = "new_lead_assignment"
	NotificationVisitScheduled    NotificationType = "visit_scheduled"
	NotificationCommissionApproved NotificationType = "commission_approved"
	NotificationLeadEscalated     NotificationType = "lead_escalated"
)

// NotificationPayload is the JSON payload for notification tasks.
type NotificationPayload struct {
	Type    NotificationType       `json:"type"`
	Phone   string                 `json:"phone"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// NotificationHandler processes notification tasks.
type NotificationHandler struct {
	smsProvider sms.SMSProvider
}

// NewNotificationHandler creates a new NotificationHandler.
func NewNotificationHandler(smsProvider sms.SMSProvider) *NotificationHandler {
	return &NotificationHandler{smsProvider: smsProvider}
}

// HandleSendNotification sends a WhatsApp/SMS notification for the given task payload.
func (h *NotificationHandler) HandleSendNotification(ctx context.Context, task *asynq.Task) error {
	var payload NotificationPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("notification: failed to unmarshal payload: %w", err)
	}

	log.Info().
		Str("type", string(payload.Type)).
		Str("phone", payload.Phone).
		Msg("worker: sending notification")

	// Send SMS notification.
	if h.smsProvider != nil && payload.Phone != "" {
		message := payload.Message
		if message == "" {
			message = buildNotificationMessage(payload.Type, payload.Data)
		}

		if err := h.smsProvider.SendSMS(payload.Phone, message); err != nil {
			log.Error().Err(err).
				Str("phone", payload.Phone).
				Str("type", string(payload.Type)).
				Msg("notification: failed to send SMS")
			return fmt.Errorf("notification: SMS send failed: %w", err)
		}

		log.Info().
			Str("phone", payload.Phone).
			Str("type", string(payload.Type)).
			Msg("notification: SMS sent successfully")
	}

	// TODO: Integrate WhatsApp Business API for richer notifications.
	// The whatsapp adapter is already available at internal/adapter/whatsapp/.

	return nil
}

// buildNotificationMessage creates a default message based on notification type.
func buildNotificationMessage(notifType NotificationType, data map[string]interface{}) string {
	switch notifType {
	case NotificationNewLeadAssignment:
		name, _ := data["lead_name"].(string)
		project, _ := data["project_name"].(string)
		return fmt.Sprintf("New lead assigned: %s interested in %s. Please contact within 5 minutes.", name, project)

	case NotificationVisitScheduled:
		date, _ := data["visit_date"].(string)
		project, _ := data["project_name"].(string)
		return fmt.Sprintf("Site visit scheduled for %s at %s.", date, project)

	case NotificationCommissionApproved:
		amount, _ := data["amount"].(float64)
		return fmt.Sprintf("Your commission of INR %.2f has been approved.", amount)

	case NotificationLeadEscalated:
		name, _ := data["lead_name"].(string)
		return fmt.Sprintf("Escalation: Lead %s has been reassigned to you. Please contact immediately.", name)

	default:
		return "You have a new notification from PropTech."
	}
}

// NewSendNotificationTask creates a new Asynq task for sending a notification.
func NewSendNotificationTask(payload NotificationPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal notification payload: %w", err)
	}
	return asynq.NewTask(TaskSendNotification, data), nil
}
