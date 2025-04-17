package handler

import (
	"fmt"

	"go.uber.org/zap"
)

func (h *Handler) SendEmailWarning(email, userID, oldIP, newIP string) {
	zap.L().Info(fmt.Sprintf("Mock Email to %s: IP address changed for user %s. Old IP: %s, New IP: %s", email, userID, oldIP, newIP))
}
