package pentagonmodel

const (
    COMPONENT_EMAIL = "email"
    SUBCOMPONENT_EMAIL_MAIN = "main"

    COMPONENT_KV = "key-value"
    SUBCOMPONENT_KV_WRITE = "write"
    SUBCOMPONENT_KV_READ = "read"

    COMPONENT_GIT = "git"
    SUBCOMPONENT_GIT_WATCH = "watch"

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
    Error string
}

/* Mail Component */
type MailMessage struct {
    To, From, Subject, Message string
}

type MailResponse struct {
    GenericSuccessResponse
}

/* Key Value Store Component */
type KeyValueReadMessage struct {
    Category, Key string
}

type KeyValueWriteMessage struct {
    Category, Key, Value string
}

type KeyValueResponse struct {
    GenericSuccessResponse
    Key, Value string
}

/* Git Component */
type GitWatchMessage struct {
    URL string
}

type GitWatchResponse struct {
    GenericSuccessResponse
    Hash, Message string
}
