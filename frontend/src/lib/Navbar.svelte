<script lang="ts">
    import { link } from "svelte-spa-router";
    import active from "svelte-spa-router/active";
    import { _ } from "svelte-i18n";

    const activeStyle =
        "[&>span]:text-slate-300 bg-primary rounded-t -mb-[1px]";
    const inactiveStyle = "-mb-[1px]";
</script>

<nav class="flex flex-wrap pl-2 mb-0 list-none border-b border-primary">
    <a
        href="/"
        use:link
        use:active={{
            className: activeStyle,
            inactiveClassName: inactiveStyle,
        }}
    >
        <span
            class="block py-1 px-4 border border-transparent rounded-t hover:border-b-primary hover:border-t-slate-300 hover:border-x-slate-300"
        >
            {$_("home")}
        </span>
    </a>
    <a
        href="/about"
        use:link
        use:active={{
            className: activeStyle,
            inactiveClassName: inactiveStyle,
        }}
    >
        <span
            class="block py-1 px-4 border border-transparent rounded-t hover:border-b-primary hover:border-t-slate-300 hover:border-x-slate-300"
        >
            {$_("about")}
        </span>
    </a>
</nav>

<!-- ACTIVE CLASS FOR A LINK:
    https://github.com/ItalyPaleAle/svelte-spa-router?tab=readme-ov-file#highlight-active-links
-->
