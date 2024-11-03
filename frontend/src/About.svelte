<script lang="ts">
    import logo from "./assets/images/logo-universal.png";
    import octocat from "./assets/images/github_octocat.png";

    const year = new Date().getFullYear();
</script>

<div class="flex flex-col items-center gap-4">
    <img alt="Wails logo" src={logo} class="h-16 block mx-auto" />

    <h1 class="text-3xl text-sky-600 font-bold">About</h1>

    <h3 class="text-xl font-medium text-amber-600 mt-8">
        Wails Todoapp is a simple task manager so you don't forget anything 😀
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

<div class="absolute bottom-0 left-0 z-10 h-12 pl-4 text-center text-sm">
    <a href="http://" target="_blank" rel="noopener noreferrer">
        ⚡ Made by emarifer&nbsp;
        <img class="w-8 inline" src={octocat} alt="GitHub logo" />
        &nbsp;|&nbsp;&nbsp;Copyright © {year} - MIT Licensed
    </a>
</div>
