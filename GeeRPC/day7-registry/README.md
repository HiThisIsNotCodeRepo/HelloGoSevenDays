# Registry

## How Registry work?

Instead of client directly call server, client will call registry to retrieve the list of server and pick one from the
list to call. The list is maintained by registry and every time when server starts up it will register itself to
registry.