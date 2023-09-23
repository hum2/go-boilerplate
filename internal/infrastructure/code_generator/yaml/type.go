package yaml

// パスファイル関連の構造体

type Root struct {
	Methods map[string]*Ref `yaml:",inline"`
}

type Ref struct {
	Ref string `yaml:"$ref"`
}

type Method struct {
	Summary     string               `yaml:"summary"`
	OperationID string               `yaml:"operationId"`
	Description string               `yaml:"description"`
	Tags        []string             `yaml:"tags"`
	Parameters  []*Ref               `yaml:"parameters,omitempty"`
	RequestBody *RequestBody         `yaml:"requestBody,omitempty"`
	Responses   map[string]*Response `yaml:"responses"`
}

type Response struct {
	Description string              `yaml:"description,omitempty"`
	Content     map[string]*Content `yaml:"content,omitempty"`
}

type Content struct {
	Schema *Schema `yaml:"schema"`
}

type Schema struct {
	Type  string `yaml:"type,omitempty"`
	Items *Ref   `yaml:"items,omitempty"`
	Ref   string `yaml:"$ref,omitempty"`
}

type RequestBody struct {
	Description string              `yaml:"description,omitempty"`
	Required    bool                `yaml:"required,omitempty"`
	Content     map[string]*Content `yaml:"content,omitempty"`
}

// リクエスト関連の構造体

type Property struct {
	Type        string `yaml:"type"`
	Format      string `yaml:"format,omitempty"`
	Description string `yaml:"description"`
}

type Properties map[string]Property

type Req struct {
	Properties Properties `yaml:"properties"`
	Required   []string   `yaml:"required"`
}

type Res struct {
	Properties Properties `yaml:"properties"`
	Required   []string   `yaml:"required"`
}
