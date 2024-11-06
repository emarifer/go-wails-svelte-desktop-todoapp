<script lang="ts">
  import { locale, _ } from "svelte-i18n";
  import Tab1 from "./Home.svelte";
  import Tab2 from "./About.svelte";
  import Tabs from "./Tabs.svelte";
  import Menu from "./Menu.svelte";

  // List of tab items with labels, values and assigned components
  let items = [
    { label: $_("home"), value: 1, component: Tab1 },
    { label: $_("about"), value: 2, component: Tab2 },
  ];

  let selected = "🌐 Language";
  const languages = [
    { code: "en", name: "English" },
    { code: "es", name: "Español" },
  ];
  const handleChange = (newLocale: string) => {
    console.log("Language code selected:", newLocale);
    locale.set(newLocale);
    items = [
      { label: $_("home"), value: 1, component: Tab1 },
      { label: $_("about"), value: 2, component: Tab2 },
    ];
  };
  // $: console.log("Changed selected:", selected);
</script>

<main class="pt-2 pb-8 flex flex-col gap-4 overflow-hidden h-full relative">
  <Menu />
  <select
    bind:value={selected}
    on:change={() => handleChange(selected["code"])}
    class="absolute top-2 right-2 select select-primary w-fit max-w-xs"
  >
    <option disabled selected>🌐 Language</option>
    {#each languages as value}<option {value}>{value.name}</option>{/each}
  </select>
  <Tabs {items} />
</main>

<!-- REFERENCES:
  https://centus.com/blog/svelte-localization
  https://github.com/kaisermann/svelte-i18n
  https://stackoverflow.com/questions/66145620/svelte-i18n-with-snowpack-shows-blank-page-without-errors

  https://svelte.dev/playground/select-bindings?version=5.1.9
  https://svelte.dev/playground/8cf99db4663c4071a939497570b9b21d?version=3.2.2
-->
