# Maps

In Go language, a map is a powerful, ingenious, and a versatile data structure. Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys. It is a reference to a hash table.

It is important to bear in mind that all keys and all values have to be the same type between them.

```$bash
// Map with keys type [string] and values type string as well
colors := map[string]string{
    "red": "#ff0000",
    "green": "#4bf745",
    "blue": "#232rs2",
}

===

colorsMap := make(map[string]string)
colorsMap["red"] = "#ff0000"
colorsMap["green"] = "#4bf745"
colorsMap["blue"] = "#232rs2"
```

It is possible delete a map position following next example:

```$bash
colors := map[string]string{
    "red": "#ff0000",
    "green": "#4bf745",
    "blue": "#232rs2",
}
delete(colors, "red")
```

In order to iterate a Map, it is required:

```$bash
colors := map[string]string{
    "red": "#ff0000",
    "green": "#4bf745",
    "blue": "#232rs2",
}
func printMap(m map[string]string) {
    for color,hex := range c {

    }
}
```

## Maps vs Structs

Maps:

- Keys have to be the same type
- Values have to be the same type
- Keys support indexing
- Reference Type (It is not required to work with pointers)
- Normally it is used to represent a collection of related properties

Structs:

- Values can be different types
- Keys don't support indexing
- It is a value type and requires pointers
