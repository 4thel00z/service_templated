package examples

type PostMultiPartMessage struct {
	// With mapstructure tag you can override the name of the field
	File []byte `multipart:"required" mapstructure:"file"`
}
