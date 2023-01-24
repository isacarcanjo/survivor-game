package entitys

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("survivor not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot Be Deleted")

//ErrInvalidName name is invalid
var ErrInvalidName = errors.New("name is Invalid")

//ErrInvalidId id is invalid
var ErrInvalidId = errors.New("id is Invalid")
