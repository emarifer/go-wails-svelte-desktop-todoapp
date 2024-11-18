<script lang="ts">
  import Router from "svelte-spa-router";
  import Navbar from "./lib/Navbar.svelte";
  import Menu from "./lib/Menu.svelte";
  import Language from "./lib/Language.svelte";
  import Home from "./pages/Home.svelte";
  import About from "./pages/About.svelte";

  const routes = {
    "/": Home,
    "/about": About,
  };
</script>

<main class="pt-2 pb-8 flex flex-col gap-4 overflow-hidden h-full relative">
  <Menu />

  <Language />

  <Navbar />

  <Router {routes} />
</main>
