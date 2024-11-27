package models

type PluginData struct {
	PluginName    string
	PluginVersion string
	PluginType    interface{}
	Parameters    map[string]interface{}
}
