// Code generated by go generate; DO NOT EDIT.
package define

// OptionTypeMap is the type map for options.
var OptionTypeMap = map[string]string{
	"md5":           "string",
	"storage_class": "string",
}

// WithMd5 will apply md5 value to Options
func WithMd5(v string) *Option {
	return &Option{
		Key:   "md5",
		Value: v,
	}
}

// WithStorageClass will apply storage_class value to Options
func WithStorageClass(v string) *Option {
	return &Option{
		Key:   "storage_class",
		Value: v,
	}
}