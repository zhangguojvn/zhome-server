package WeiChat

type WeiChatFunction interface {
	Init() error
	Stop() error
}
