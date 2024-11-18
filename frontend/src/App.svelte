<script lang="ts">
  import { Route, Router } from "svelte-routing";
  import Navbar from "./lib/Navbar.svelte";
  import Menu from "./lib/Menu.svelte";
  import Language from "./lib/Language.svelte";
  import Home from "./pages/Home.svelte";
  import About from "./pages/About.svelte";
</script>

<main class="pt-2 pb-8 flex flex-col gap-4 overflow-hidden h-full relative">
  <Menu />

  <Language />

  <Router>
    <Navbar />

    <Route path="/">
      <Home />
    </Route>
    <Route path="/about">
      <About />
    </Route>
  </Router>
</main>
