<script lang="ts">
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import logo from "./assets/images/logo-universal.png";
    import octocat from "./assets/images/github_octocat.png";

    let visible = false;

    onMount(async () => {
        visible = !visible;
    });

    const year = new Date().getFullYear();
</script>

{#if visible}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 20 }}
        class="flex flex-col h-full justify-between"
    >
        <div class="flex flex-col items-center gap-4">
            <img alt="Wails logo" src={logo} class="h-16 block mx-auto" />

            <h1 class="text-3xl text-sky-600 font-bold">About</h1>

            <h3 class="text-xl font-medium text-amber-600 mt-8">
                Wails Todoapp is a simple task manager so you don't forget
                anything 😀
            </h3>
            <p class="text-slate-600 font-medium mt-4">More info:</p>
            <a
                class="hover:text-sky-500 ease-in duration-300"
                href="http://"
                target="_blank"
                rel="noopener noreferrer"
            >
                https://github.com/emarifer/go-wails-svelte-desktop-todoapp
            </a>
        </div>

        <div class="h-12 w-full pl-4 text-start text-sm">
            <a href="http://" target="_blank" rel="noopener noreferrer">
                ⚡ Made by emarifer&nbsp;
                <img class="w-8 inline" src={octocat} alt="GitHub logo" />
                &nbsp;|&nbsp;&nbsp;Copyright © {year} - MIT Licensed
            </a>
        </div>
    </div>
{/if}

<!-- 
https://stackoverflow.com/questions/59812003/tailwindcss-fixed-sticky-footer-on-the-bottom

TRANSITIONS:
https://svelte.dev/playground/in-and-out
https://www.youtube.com/watch?v=5oEo98BrRqs
 -->
