<script lang="ts">
    import { Link } from "svelte-routing";
    import type { GetPropsParams } from "../../node_modules/svelte-routing/types/Link";
    import { _ } from "svelte-i18n";

    const getPropsLink =
        (path: string) =>
        ({ href, isPartiallyCurrent, isCurrent }: GetPropsParams) => {
            const isActive =
                href === path ? isCurrent : isPartiallyCurrent || isCurrent;
            return isActive
                ? {
                      class: "[&>span]:text-slate-300 bg-primary rounded-t -mb-[1px]",
                  }
                : { class: "-mb-[1px]" };
        };
</script>

<nav class="flex flex-wrap pl-2 mb-0 list-none border-b border-primary">
    <Link getProps={getPropsLink("/")} to="/">
        <span
            class="block py-1 px-4 border border-transparent rounded-t hover:border-b-primary hover:border-t-slate-300 hover:border-x-slate-300"
        >
            {$_("home")}
        </span>
    </Link>

    <Link getProps={getPropsLink("/about")} to="/about">
        <span
            class="block py-1 px-4 border border-transparent rounded-t hover:border-b-primary hover:border-t-slate-300 hover:border-x-slate-300"
        >
            {$_("about")}
        </span>
    </Link>
</nav>

<!-- ACTIVE CLASS FOR A LINK:
    https://github.com/EmilTholin/svelte-routing/issues/63
-->
