---
title: Nuxt.js Walkthrough
slug: nuxt
heading: Nuxt.js Walkthrough
createdAt: 2024-08-13T18:48:40.767Z
updatedAt: 2024-08-13T18:48:40.767Z
---

## Overview

Our Node.js SDK integrates seamlessly with [Nuxt.js](https://nuxt.com), providing both frontend session replays and server-side monitoring capabilities.

## Installation

```shell
npm install @highlight-run/node
```

## Server Instrumentation

Create a server plugin:

```typescript
// server/plugins/highlight.ts
import { H, type NodeOptions } from '@highlight-run/node'

export default defineNitroPlugin((nitroApp) => {
	const highlightConfig: NodeOptions = {
		projectID: process.env.HIGHLIGHT_PROJECT_ID,
	}

	if (!H.isInitialized()) {
		H.init(highlightConfig)
	}

	nitro.hooks.hook("request", (event) => {
		const headers = event?.node.req.headers
		const parsed = H.parseHeaders(headers)
		if (parsed) {
			H.startActiveSpan(`${event.method} ${event.path}`, {
				secureSessionId: parsed.secureSessionId,
				requestId: parsed.requestId,
			})
		}
	})

	nitro.hooks.hook('afterResponse', () => {
		H.endActiveSpans()
	})

	nitro.hooks.hook('error', async (error, { event }) => {
		const headers = event?.node.req.headers!
		const parsed = H.parseHeaders(headers)
		if (parsed !== undefined) {
			H.consumeError(error, parsed?.secureSessionId, parsed?.requestId)
		}
	})
})
```

## Validation

Create a test API route:

```typescript
// server/api/test-error.ts
export default defineEventHandler(() => {
	throw new Error('Test server-side error')
})
```

Test with:

```bash
curl http://localhost:3000/api/test-error
```

Check your Highlight dashboard for the log at https://app.highlight.io/sessions.