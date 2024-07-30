package transpiler

type (
	String         string
	Bool           bool
	Int            int
	Float          float32
	Strings        []string
	SecondaryFiles []CWLSecondaryFileSchema
)

type CWLFormatKind int32

const (
	FormatStringKind CWLFormatKind = iota
	FormatStringsKind
	FormatExpressionKind
)

type CWLFormat struct {
	Kind       CWLFormatKind
	String     String
	Strings    Strings
	Expression CWLExpression
}

type CWLStdin struct{}

type LoadListingEnum string

const (
	LoadListingNone    LoadListingEnum = "no_listing"
	LoadListingDeep    LoadListingEnum = "deep_listing"
	LoadListingShallow LoadListingEnum = "shallow_listing"
)

type CWLExpressionKind int32

const (
	RawKind CWLExpressionKind = iota
	ExpressionKind
	BoolKind
	IntKind
	FloatKind
)

type CWLExpression struct {
	Kind       CWLExpressionKind
	Raw        string
	Expression string
	Bool       bool
	Int        int
	Float      float64
}

type CWLSecondaryFileSchema struct {
	Pattern  CWLExpression `yaml:"pattern"`
	Required CWLExpression `yaml:"required"`
}

type CWLInputProvider interface {
	GetKind() string
}
