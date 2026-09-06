import React from "react";
import {
  Link,
  Outlet,
  createRootRoute,
  createRoute,
  createRouter,
} from "@tanstack/react-router";

// gotossr calls this factory with createMemoryHistory on the server and
// createBrowserHistory in the browser. Props are identical on both sides.
export function createSSRRouter({ history, props }: { history: any; props: { message?: string } }) {
  const rootRoute = createRootRoute({
    component: () => <main><nav><Link to="/">Home</Link> | <Link to="/about">About</Link></nav><Outlet /></main>,
  });
  const indexRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "/",
    component: () => <><h1>{props.message}</h1><p>This HTML came from TanStack Router SSR.</p></>,
  });
  const aboutRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "/about",
    component: () => <h1>About: rendered at the requested route</h1>,
  });
  return createRouter({ routeTree: rootRoute.addChildren([indexRoute, aboutRoute]), history });
}

