package main

const validJsonArray0 = `
    [
   {
      "a":1,
      "b":3
   },
   {
      "a":1,
      "b":3
   }
]
`

const validJsonMap0 = `
    {
   "foo":{
      "a":1,
      "b":2
   },
   "bar":{
      "a":1,
      "b":2
   }
}`

const validJsonMap1 = `
    {
   "foo":{
      "a":1,
      "b":2
   }
}`

var validMaps = []string{validJsonMap0, validJsonMap1}

var validArrays = []string{validJsonArray0}
