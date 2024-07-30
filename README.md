## pwgen

quickly create a random password

### Synopsis

A command line tool to quickly create a random password.

### Options

```
  -h, --help   help for pwgen
```

### SEE ALSO

* [pwgen number](#pwgen-number)	 - Generate a random number
* [pwgen phrase](#pwgen-phrase)	 - Generate a phrase of multiple words, leveraging the words API
* [pwgen random](#pwgen-random)	 - Generate a random string
* [pwgen completion](#pwgen-completion)	 - Generate the autocompletion script for the specified shell

---

## pwgen number

Generate a random number

```
pwgen number [flags]
```

### Options

```
  -c, --count int8   Number of digits (default 4)
  -h, --help         help for number
```

### SEE ALSO

* [pwgen](pwgen.md)	 - quickly create a random password

---

## pwgen phrase

Generate a phrase of multiple words, leveraging the words API

### Synopsis

This generates a phrase or list of words. This function does not
use a local dictionary, but fetches the words from the words API from
random-word-api.herokuapp.com. Please take this into consideration when
generating phrases to use as passwords.

```
pwgen phrase [flags]
```

### Options

```
  -c, --count int          Number of words to fetch (default 3)
  -h, --help               help for phrase
  -s, --separator string   Separator bewteen words (default "-")
```

### SEE ALSO

* [pwgen](pwgen.md)	 - quickly create a random password

---

## pwgen random

Generate a random string

```
pwgen random [flags]
```

### Options

```
  -c, --count int   Number of characters (default 12)
  -h, --help        help for random
  -l, --lowerCase   include lower case characters
  -n, --numbers     include numbers
  -s, --symbols     include special characters: 
  -u, --upperCase   include upper case characters
```

### SEE ALSO

* [pwgen](pwgen.md)	 - quickly create a random password

---

## pwgen completion

Generate the autocompletion script for the specified shell

### Synopsis

Generate the autocompletion script for pwgen for the specified shell.
See each sub-command's help for details on how to use the generated script.


### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [pwgen](pwgen.md)	 - quickly create a random password
* [pwgen completion bash](pwgen_completion_bash.md)	 - Generate the autocompletion script for bash
* [pwgen completion fish](pwgen_completion_fish.md)	 - Generate the autocompletion script for fish
* [pwgen completion powershell](pwgen_completion_powershell.md)	 - Generate the autocompletion script for powershell
* [pwgen completion zsh](pwgen_completion_zsh.md)	 - Generate the autocompletion script for zsh

