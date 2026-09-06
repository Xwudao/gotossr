# TanStack Router SSR example

```bash
pnpm install --dir frontend
go run .
```

`/` and `/about` are rendered with TanStack Router and `createMemoryHistory` inside gotossr's JavaScript runtime. The client bundle uses `createBrowserHistory` plus `hydrateRoot`, and receives the same serialized Go props.

## App contract

When `SPAHydrationMode: "tanstack"` is selected, `ClientAppPath` must named-export:

```ts
export function createSSRRouter({ history, props }) {
  return createRouter({ routeTree, history })
}
```

`history` is a TanStack history. `props` contains Go's `RenderConfig.Props` plus `__requestPath`.
