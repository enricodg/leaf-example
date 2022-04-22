package forwarder

import "context"

func (u *ucForwardMessage) ForwardMessage(ctx context.Context, message string) error {
	return u.Outbound.Messaging.Hello.PublishMessage(ctx, message)
}
