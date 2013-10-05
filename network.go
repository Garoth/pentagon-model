package pentagonmodel

const (
    COMPONENT_EMAIL = "email"
    SUBCOMPONENT_EMAIL_MAIN = "main"
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

type MailComponentMessage struct {
    To, From, Subject, Message string
}

type MailComponentResponse struct {
    GenericSuccessResponse
}
