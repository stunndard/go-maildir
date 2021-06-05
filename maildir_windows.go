package maildir

// The separator separates a messages unique key from its flags in the filename.
// Unfortunately, Windows doesn't allow ':' in file names
const separator rune = ';'
