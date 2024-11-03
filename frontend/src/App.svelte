<script lang="ts">
  import Tab1 from "./Home.svelte";
  import Tab2 from "./About.svelte";
  import Tabs from "./Tabs.svelte";

  // List of tab items with labels, values and assigned components
  let items = [
    { label: "Home", value: 1, component: Tab1 },
    { label: "About", value: 2, component: Tab2 },
  ];
</script>

<main class="pt-2 pb-8 flex flex-col gap-4">
  <Tabs {items} />
</main>
