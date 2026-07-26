import { Router, Route } from "@solidjs/router";

import Home from "../routes/Home";
import Settings from "../routes/Settings";
import Stats from "../routes/Stats";

// All top-level routes in one place, so adding or removing a page never
// requires touching main.jsx.
export default function AppRouter() {
  return (
    <Router>
      <Route path="/" component={Home} />
      <Route path="/settings" component={Settings} />
      <Route path="/stats" component={Stats} />
    </Router>
  );
}
