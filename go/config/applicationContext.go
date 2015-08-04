// applicationContext.go
package config

var propertySource PropertySource

func init(){ 
	propertySource = PropertySource{}
}
func GetPropertySource() *PropertySource{
	return &propertySource
}


type PropertySource struct {
	configmap map[interface{}]interface{}
	profile   string // is this private visibility really need?
}

func (p PropertySource) GetProfile() string{
	return p.profile
}
func (p *PropertySource) SetProfile(profile string){
	p.profile = profile
}

//TODO make function to apply dev environment profile 