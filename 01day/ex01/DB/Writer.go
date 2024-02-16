package DB

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func (cookbook CookBook) Write(t string) error {
	if t == "json" {
		data, err := json.MarshalIndent(cookbook, "", "    ")
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(data))
	} else if t == "xml" {
		data, err := xml.MarshalIndent(cookbook, "", "    ")
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stdout, string(data))
	} else {
		return fmt.Errorf("wrong type of writing '%s': xml or json extension needed", t)
	}
	return nil
}
