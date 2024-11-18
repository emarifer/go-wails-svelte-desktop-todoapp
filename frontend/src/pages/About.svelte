<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { _ } from "svelte-i18n";
    import octocat from "../assets/images/github_octocat.png";
    import logo from "../assets/images/logo-universal.png";

    let mounted = false;
    onMount(() => {
        mounted = true;
    });

    const year = new Date().getFullYear();
</script>

{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex flex-col h-full"
    >
        <img alt="Wails logo" src={logo} class="mt-1.5 h-16 block mx-auto" />

        <div class="pt-12 flex flex-col justify-between h-full">
            <div class="flex flex-col items-center gap-4">
                <h1 class="text-3xl text-sky-600 font-bold">
                    {$_("title_about_page")}
                </h1>

                <h3 class="text-xl font-medium text-amber-600 mt-8">
                    {$_("explanation")} 😀
                </h3>
                <p class="text-slate-600 font-medium mt-4">{$_("more_info")}</p>
                <a
                    class="hover:text-sky-500 ease-in duration-300"
                    href="http://"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    https://github.com/emarifer/go-wails-svelte-desktop-todoapp
                </a>
            </div>

            <div class="h-12 pl-4 text-start text-sm">
                <a href="http://" target="_blank" rel="noopener noreferrer">
                    ⚡ {$_("credits")} emarifer&nbsp;
                    <img class="w-8 inline" src={octocat} alt="GitHub logo" />
                    &nbsp;|&nbsp;&nbsp;Copyright © {year} - MIT Licensed
                </a>
            </div>
        </div>
    </div>
{/if}
