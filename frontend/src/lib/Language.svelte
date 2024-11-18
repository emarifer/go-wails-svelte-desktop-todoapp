<script lang="ts">
    import { onMount } from "svelte";
    import { locale } from "svelte-i18n";
    import { GetLanguage, SaveLanguage } from "../../wailsjs/go/main/App";

    type language = {
        code: string;
        name: string;
    };

    const languages: language[] = [
        { code: "en", name: "English" },
        { code: "es", name: "Español" },
    ];

    onMount(() => {
        GetLanguage().then((result) => {
            if (result.length != 0) locale.set(result);
        });
    });

    let selected = $state({ code: "", name: "🌐 Language" } as language);
    const handleChange = (newLocale: language) => {
        console.log("Language selected:", newLocale);
        locale.set(newLocale["code"]);
        SaveLanguage(newLocale["code"]);
    };
</script>

<select
    bind:value={selected}
    onchange={() => handleChange(selected)}
    class="absolute top-2 right-2 z-20 select select-primary w-fit max-w-xs"
>
    {#if selected.code == ""}
        <option disabled value={selected}>{selected.name}</option>
    {:else}
        <option disabled>🌐 Language</option>
    {/if}
    {#each languages as language}
        <option value={language}>{language.name}</option>
    {/each}
</select>
