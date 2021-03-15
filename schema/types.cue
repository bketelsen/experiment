package schema

import (
  "strings"
)


#World: {
  name:     string
  camelName:  strings.ToCamel(name)
  TitleName:  strings.ToTitle(name)
  UPPER_NAME: strings.ToUpper(name)

  person:  string
  sites: [string]
  personPath: "\(person)"
}
 
#Person: {
  first_name: string
  last_name: string
  age: int 
}

#Site: {
  name: string
  environments: [#Environment]
}

#Environment: {
  name: string
  url: string
}