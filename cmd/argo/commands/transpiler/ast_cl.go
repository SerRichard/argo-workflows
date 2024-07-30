package transpiler

// interface due to large number of child types
type CWLRequirements interface {
	isCWLRequirement()
	getClass() string
}

type DockerRequirement struct {
	Class                 string  `yaml:"class"`
	DockerPull            *string `yaml:"dockerPull"`
	DockerLoad            *string `yaml:"dockerLoad"`
	DockerFile            *string `yaml:"dockerFile"`
	DockerImport          *string `yaml:"dockerImport"`
	DockerImageId         *string `yaml:"dockerImageId"`
	DockerOutputDirectory *string `yaml:"dockerOutputDirectory"`
}

type SoftwarePackage struct {
	Package string  `yaml:"package"`
	Version Strings `yaml:"version"`
	Specs   Strings `yaml:"specs"`
}

type SoftwareRequirement struct {
	Class    string            `yaml:"class"` // constant SoftwareRequirement
	Packages []SoftwarePackage `yaml:"packages"`
}

type LoadListingRequirement struct {
	Class       string           `yaml:"class"` // constant LoadListingRequirement
	LoadListing *LoadListingEnum `yaml:"loadListing"`
}

/* type Dirent struct {
	entry     CWLExpression  `yaml:"entry"`
	entryName *CWLExpression `yaml:"entryName"`
	writeable *bool          `yaml:"writeable"`
} */

type InitialWorkDirRequirementListing interface {
	isInitialWorkDirRequirementListing()
}

type InitialWorkDirRequirement struct {
	Class   string                           `yaml:"class"` // constant InitialWorkDirRequirement
	Listing InitialWorkDirRequirementListing `yaml:"listing"`
}

type InlineJavascriptRequirement struct {
	Class         string  `yaml:"class"` // constant InlineJavascriptRequirement
	ExpressionLib Strings `yaml:"expressionLib"`
}

type SchemaDefRequirementType interface {
	isSchemaDefRequirementType()
}

type SchemaDefRequirement struct {
	Class string                     `yaml:"class"` // constant SchemaDefRequirement
	Types []SchemaDefRequirementType `yaml:"types"`
}

type EnvironmentDef struct {
	EnvName  string        `yaml:"envName"`
	EnvValue CWLExpression `yaml:"envValue"`
}

type EnvVarRequirement struct {
	Class  string           `yaml:"class"` // constant EnvVarRequirement
	EnvDef []EnvironmentDef `yaml:"envDef"`
}

type ShellCommandRequirement struct {
	Class string `yaml:"class"` // constant ShellCommandRequirement
}

type WorkReuse struct {
	Class       string        `yaml:"class"` // constant WorkReuse
	EnableReuse CWLExpression `yaml:"enableReuse"`
}

type NetworkAccess struct {
	Class         string // constant NetworkAccess
	NetworkAccess CWLExpression
}

type InplaceUpdateRequirement struct {
	Class         string `yaml:"class"` // constant InplaceUpdateRequirement
	InplaceUpdate Bool   `yaml:"inplaceUpdate"`
}

type ToolTimeLimit struct {
	Class     string        `yaml:"class"` // constant ToolTimeLimit
	TimeLimit CWLExpression `yaml:"timeLimit"`
}

type ResourceRequirement struct {
	Class     string         `yaml:"class"` // constand ResourceRequirement
	CoresMin  *CWLExpression `yaml:"coresMin"`
	CoresMax  *CWLExpression `yaml:"coresMax"`
	RamMin    *CWLExpression `yaml:"ramMin"`
	RamMax    *CWLExpression `yaml:"ramMax"`
	TmpdirMin *CWLExpression `yaml:"tmpdirMin"`
	TmpdirMax *CWLExpression `yaml:"tmpdirMax"`
	OutdirMin *CWLExpression `yaml:"outdirMin"`
	OutdirMax *CWLExpression `yaml:"outdirMax"`
}

func (InlineJavascriptRequirement) isCWLRequirement()  {}
func (d InlineJavascriptRequirement) getClass() string { return d.Class }

func (SchemaDefRequirement) isCWLRequirement()  {}
func (d SchemaDefRequirement) getClass() string { return d.Class }

func (LoadListingRequirement) isCWLRequirement()  {}
func (d LoadListingRequirement) getClass() string { return d.Class }

func (DockerRequirement) isCWLRequirement()  {}
func (d DockerRequirement) getClass() string { return d.Class }

func (SoftwareRequirement) isCWLRequirement()  {}
func (d SoftwareRequirement) getClass() string { return d.Class }

func (InitialWorkDirRequirement) isCWLRequirement()  {}
func (d InitialWorkDirRequirement) getClass() string { return d.Class }

func (EnvVarRequirement) isCWLRequirement()  {}
func (d EnvVarRequirement) getClass() string { return d.Class }

func (ShellCommandRequirement) isCWLRequirement()  {}
func (d ShellCommandRequirement) getClass() string { return d.Class }

func (WorkReuse) isCWLRequirement()  {}
func (d WorkReuse) getClass() string { return d.Class }

func (NetworkAccess) isCWLRequirement()  {}
func (d NetworkAccess) getClass() string { return d.Class }

func (InplaceUpdateRequirement) isCWLRequirement()  {}
func (d InplaceUpdateRequirement) getClass() string { return d.Class }

func (ToolTimeLimit) isCWLRequirement()  {}
func (d ToolTimeLimit) getClass() string { return d.Class }

func (ResourceRequirement) isCWLRequirement()  {}
func (d ResourceRequirement) getClass() string { return d.Class }

func (CommandlineInputRecordSchema) isSchemaDefRequirementType() {}
func (CommandlineInputEnumSchema) isSchemaDefRequirementType()   {}
func (CommandlineInputArraySchema) isSchemaDefRequirementType()  {}
func (DockerRequirement) isSchemaDefRequirementType()            {}
func (SoftwareRequirement) isSchemaDefRequirementType()          {}
func (InitialWorkDirRequirement) isSchemaDefRequirementType()    {}

type CommandlineInputRecordField struct {
	Name           string              `yaml:"name"`
	Type           CWLTypes            `yaml:"type"` // len(1) represents scalar len > 1 represents array
	Doc            Strings             `yaml:"doc"`
	Label          *string             `yaml:"label"`
	SecondaryFiles SecondaryFiles      `yaml:"secondaryFiles"`
	Streamable     *bool               `yaml:"streamable"`
	Format         CWLFormat           `yaml:"format"`
	LoadContents   *bool               `yaml:"loadContents"`
	LoadListing    LoadListingEnum     `yaml:"loadListing"`
	InputBinding   *CommandlineBinding `yaml:"inputBinding"`
}

type CommandlineInputArraySchema struct {
	Items        CWLTypes            `yaml:"items"`
	Type         string              `yaml:"type"` // MUST be array
	Label        *string             `yaml:"label"`
	Doc          Strings             `yaml:"doc"`
	Name         *string             `yaml:"name"`
	InputBinding *CommandlineBinding `yaml:"inputBinding"`
}

type CommandlineInputEnumSchema struct {
	Symbols      Strings `yaml:"symbols"`
	Type         string  `yaml:"type"` // MUST BE enum, only a placeholder for type verification purposes
	Label        *string `yaml:"label"`
	Doc          Strings `yaml:"doc"`
	Name         *string `yaml:"name"`
	InputBinding *CommandlineBinding
}

type CommandlineInputRecordSchema struct {
	Type   string                         `yaml:"type"` // MUST BE "record"
	Fields *[]CommandlineInputRecordField `yaml:"fields"`
	Label  *string                        `yaml:"label"`
	Doc    *Strings                       `yaml:"doc"`
	Name   *string                        `yaml:"name"`
	// will be used for processing later on hence we disable the linter
	inputBinding *CommandlineBinding `yaml:"inputBinding"` //nolint:unused,structcheck
}

type Type int32

const (
	CWLNullKind Type = iota
	CWLBoolKind
	CWLIntKind
	CWLLongKind
	CWLFloatKind
	CWLDoubleKind
	CWLFileKind
	CWLDirectoryKind
	CWLStdinKind
	CWLStringKind
	CWLRecordKind
	CWLRecordFieldKind
	CWLEnumKind
	CWLArrayKind
)

type CWLType struct {
	Kind   Type
	Record *CommandlineInputRecordSchema
	Enum   *CommandlineInputEnumSchema
	Array  *CommandlineInputArraySchema
	File   *CWLFile
}

type CWLTypes []CWLType

type CommandlineBinding struct {
	LoadContents  *bool         `yaml:"loadContents"`
	Position      *int          `yaml:"position"`
	Prefix        *string       `yaml:"prefix"`
	Separate      *bool         `yaml:"separate"`
	ItemSeperator *string       `yaml:"itemSeperator"`
	ValueFrom     CWLExpression `yaml:"valueFrom"`
	ShellQuote    *bool         `yaml:"bool"`
}

type CommandlineInputParameter struct {
	Type           CWLTypes            `yaml:"type"` // len(1) == scalar while len > 1 == array
	Label          *string             `yaml:"label"`
	SecondaryFiles SecondaryFiles      `yaml:"secondaryFiles"` // len(1) == scalar while len > 1 == array
	Streamable     *bool               `yaml:"streamable"`
	Doc            Strings             `yaml:"doc"`
	Id             *string             `yaml:"id"`
	Format         *CWLFormat          `yaml:"format"`
	LoadContents   *bool               `yaml:"loadContents"`
	LoadListing    *LoadListingEnum    `yaml:"loadListing"`
	Default        interface{}         `yaml:"default"`
	InputBinding   *CommandlineBinding `yaml:"inputBinding"`
}

type OutputBindingGlobKind int32

const (
	GlobStringKind OutputBindingGlobKind = iota
	GlobStringsKind
	GlobExpressionKind
)

type CommandlineOutputBindingGlob struct {
	Kind       OutputBindingGlobKind
	String     *string
	Strings    []string
	Expression CWLExpression
}

type CommandlineOutputBinding struct {
	LoadContents *bool                        `yaml:"loadContents"`
	LoadListing  LoadListingEnum              `yaml:"loadListing"`
	Glob         CommandlineOutputBindingGlob `yaml:"glob"`
	OutputEval   CWLExpression                `yaml:"outputEval"`
}

type CommandlineOutputParameter struct {
	Type           CWLTypes                  `yaml:"type"`
	Label          *string                   `yaml:"label"`
	SecondaryFiles SecondaryFiles            `yaml:"secondaryFiles"`
	Streamable     *bool                     `yaml:"streamable"`
	Doc            Strings                   `yaml:"doc"`
	Id             *string                   `yaml:"id"`
	Format         *CWLFormat                `yaml:"format"`
	OutputBinding  *CommandlineOutputBinding `yaml:"outputBinding"`
}

type CommandlineArgumentKind int32

const (
	ArgumentStringKind CommandlineArgumentKind = iota
	ArgumentExpressionKind
	ArgumentCLIBindingKind
)

type CommandlineArgument struct {
	Kind               CommandlineArgumentKind
	String             String
	Expression         CWLExpression
	CommandlineBinding CommandlineBinding
}

type Inputs []CommandlineInputParameter
type Outputs []CommandlineOutputParameter
type Requirements []CWLRequirements
type Hints struct {
	Array []interface{}
	Map   map[string]interface{}
}
type Arguments []string // CommandlineArgument

type CommandlineTool struct {
	Inputs       Inputs         `yaml:"inputs"`
	Outputs      Outputs        `yaml:"outputs"`
	Class        string         `yaml:"class"` // Must be "CommandLineTool"
	Id           *string        `yaml:"id"`
	Label        *string        `yaml:"label"`
	Doc          Strings        `yaml:"doc"`
	Requirements Requirements   `yaml:"requirements"`
	Hints        Hints          `yaml:"hints"`
	CWLVersion   *string        `yaml:"cwlVersion"`
	Intent       Strings        `yaml:"intent"`
	BaseCommand  Strings        `yaml:"baseCommand"`
	Arguments    Arguments      `yaml:"arguments"`
	Stdin        *CWLExpression `yaml:"stdin"`
	Stderr       *CWLExpression `yaml:"stderr"`
	Stdout       *CWLExpression `yaml:"stdout"`
}