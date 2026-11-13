package views

// PathFeedback takes bug reports and ideas; they are saved as files on the host.
const PathFeedback = "/feedback"

// Report kinds: a person picks a bug or an idea; the server writes error
// reports itself for a 500.
const (
	FeedbackBug   = "bug"
	FeedbackIdea  = "idea"
	FeedbackError = "error"
)

// MaxScreenshots is how many pictures one report may carry.
const MaxScreenshots = 3

// Classes that show an element for one kind only, keyed on the form's
// checked kind, so switching keeps everything typed.
const (
	bugOnly  = "group-has-[[name=kind][value=idea]:checked]/feedback:hidden"
	ideaOnly = "hidden group-has-[[name=kind][value=idea]:checked]/feedback:inline"
)

const feedbackLabel = "Report a bug or suggest an idea"
