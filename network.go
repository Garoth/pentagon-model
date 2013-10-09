package pentagonmodel

const (
    COMPONENT_EMAIL = "email"
    SUBCOMPONENT_EMAIL_MAIN = "main"
    COMPONENT_KV = "key-value"
    SUBCOMPONENT_KV_WRITE = "write"
    SUBCOMPONENT_KV_READ = "read"

    KV_ACTION_WRITE = "write"
    KV_ACTION_READ = "read"
)

/* Message client sends to direct it's query to a component */
type ClientHeader struct {
    Component, Subcomponent string
}

/* Header sent between components to determine what content's coming */
type InternalHeader struct {
    MessageType string
}

type GenericSuccessResponse struct {
    Success bool
    Message string
}

/* Mail Component */
type MailComponentMessage struct {
    To, From, Subject, Message string
}

type MailComponentResponse struct {
    GenericSuccessResponse
}

/* Key Value Store Component */
type KeyValueReadMessage struct {
    Category, Key string
}

type KeyValueWriteMessage struct {
    Category, Key, Value string
}

type KeyValueComponentResponse struct {
    GenericSuccessResponse
}
