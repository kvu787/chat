Chat implements a basic file-based chat service for users logged into the 
same machine. 

`input` starts a prompt to write into a shared chat log. 

`display` prints the chat log and continually checks for new writes.

**Sample usage**

Alice and Bob are logged into the same machine. Each have root permissions.
They agree to store their chat log in the file `/cht`.

Alice runs the following commands in separate terminals:

```bash
input -f /cht -p alice
display -f /cht
```

Bob also runs the following commands in separate terminals:

```bash
input -f /cht -p bob
display -f /cht
```

They can leave and reenter their chat by using the same chat log at `/cht`.