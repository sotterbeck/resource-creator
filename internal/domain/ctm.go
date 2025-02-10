package domain

import "fmt"

type CTM interface {
	GetProps() map[string]string
}

type CTMProps struct {
	props map[string]string
}

// NewCTMProps creates a new CTMProps instance with all the necessary properties.
// The additional map can be used to add extra properties depending on the method.
func NewCTMProps(method string, startTile, endTile int, additional map[string]string) (*CTMProps, error) {
	if startTile > endTile {
		return &CTMProps{}, fmt.Errorf("start tile greater than end tile")
	}

	props := map[string]string{
		"method": method,
		"tiles":  fmt.Sprintf("%d-%d", startTile, endTile),
	}

	for k, v := range additional {
		props[k] = v
	}

	return &CTMProps{props: props}, nil
}

func (c *CTMProps) GetProps() map[string]string {
	return c.props
}
