# Definition

```ad-info
title: Definition
Streaming is when, instead of sending all data at ones, you **send it in chunks**. Perhaps, it’s a large object that **gets split into smaller parts**.

```

# Usage

- Transfer of video data => App using SignalR downloads the video in small chunks

```ad-question
title: Why raw data cannot be sent as individual messages
- Many messages sent in a short period of time => Exxcessive amount of data to deal with
- Every message in SignalR has *a JSON wrapper* => Computational overhead separating data from JSON envelope
- Raw bytes need to be assembled into something consumable

```

```ad-tip
title: Benefits of SignalR
- Improve performance by NOT using envelope in every message
- Easy to maintain message order

```

# Client streaming vs Server streaming

| Client-to-server                                                                                   | Server-to-client |
| -------------------------------------------------------------------------------------------------- | ---------------- |
| - Client controls the stream </br> - Clients open the stream, write data and then close the stream | - Clients initiate the streaming process </br> - Clients call a relevant endpoint and subscribe  to the streaming channel </br> - The channel keeps listening until there is no message being sent                 |


