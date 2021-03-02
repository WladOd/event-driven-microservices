package model


type Service interface {
    Start(bootstrapServers, stateDir string, properties map[string]interface{})
	Stop()
}