// Package action contains the type Action and some basic actions.
//
// The actions are functions that act on the file system (although they do not
// necessarily have to do it).
//
// Logging Actions
//
// Actions related to logging tasks accept a special syntax (known as Shell
// Syntax) that allows you to insert the value of the environment variables
// into the text provided, the syntax is as follows:
//  $Var // Prints the value of Var.
//  ${Var} // Like the previous one.
//  $$Var // Prints $Var.
//  $NoExistentVariable // Do not print anything.
//
// In addition, special environment variables are defined to provide more
// information:
//  - Path: The absolute path of current file.
//  - BaseName: The last element of the current path.
//  - Extension: The extension of the path in uppercase and without the dot.
//  - ShellPath: The absolute path of the form: ~/very/long/path
//  - ShortPath: The absolute path of the form: ~/v/l/path
//
// Logging actions are the following:
//  - Print
//  - Write
//
// Each logging action has its own way of formatting the arguments, see their
// descriptions for more details.
package action

// Action represents the signature of an action, where path refers to the
// action path and arguments of the action. Returns a newPath chain that
// specifies the new acting path, for example, when the file moves, an empty
// string indicates that the file was deleted. If the file path is not
// modified, you must return where it was provided in path.
type Action func(path string, args []string) (newPath string)
