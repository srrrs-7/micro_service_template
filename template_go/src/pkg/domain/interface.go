package domain

type Store interface {
	GetTable() ([]*Entity, error)
	CreateTable() ([]*Entity, error)
	DeleteTable() ([]*Entity, error)
	UpdateTable() ([]*Entity, error)
}

type Server interface {
	GetDomain() ([]*Entity, error)
	CreateDomain() ([]*Entity, error)
	DeleteDomain() ([]*Entity, error)
	UpdateDomain() ([]*Entity, error)
}

type QueueHandler interface {
	SendMessage() ([]*Entity, error)
	ReceiveMessage() ([]*Entity, error)
	DeleteMessage() ([]*Entity, error)
	CreateQueue() ([]*Entity, error)
	DeleteQueue() ([]*Entity, error)
	GetQueueAttributes() ([]*Entity, error)
	SetQueueAttributes() ([]*Entity, error)
}

type SftpHandler interface {
	Upload() ([]*Entity, error)
	Download() ([]*Entity, error)
}
