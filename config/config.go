package config

type Config struct {
	getterPrefix string
	setterPrefix string
	ptrReceiver  bool
}

func (c *Config) PtrReceiver() bool {
	return c.ptrReceiver
}

func (c *Config) SetPtrReceiver(ptrReceiver bool) {
	c.ptrReceiver = ptrReceiver
}

func (c *Config) SetterPrefix() string {
	return c.setterPrefix
}

func (c *Config) SetSetterPrefix(setterPrefix string) {
	c.setterPrefix = setterPrefix
}

func (c *Config) GetterPrefix() string {
	return c.getterPrefix
}

func (c *Config) SetGetterPrefix(getterPrefix string) {
	c.getterPrefix = getterPrefix
}
