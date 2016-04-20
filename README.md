[![wercker status](https://app.wercker.com/status/49e5be1894a79de99e41f6f9b975eb0d/m "wercker status")](https://app.wercker.com/project/bykey/49e5be1894a79de99e41f6f9b975eb0d)

# World of FluxCraft - Event Processing

Event Processor for the World of FluxCraft game sample. This service will listen on the appropriate queues for inbound events in need of processing. The service
will then store the event in the event store, process it, and emit whatever outbound messages are necessary in order to allow clients to be informed as to state changes.

Downstream, the _reality_ service (game state cache) will be updated as well.
