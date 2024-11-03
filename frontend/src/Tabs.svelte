<script lang="ts">
  export let items = [];
  export let activeTabValue = 1;

  const handleClick = (tabValue) => () => (activeTabValue = tabValue);
</script>

<ul class="flex flex-wrap pl-0 mb-0 list-none border-b border-primary">
  {#each items as item}
    <li
      class={activeTabValue === item.value
        ? "[&>span]:text-slate-300 bg-primary rounded-t -mb-[1px]"
        : "-mb-[1px]"}
    >
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <span
        class="block cursor-pointer py-1 px-4 border border-transparent rounded-t hover:border-b-primary hover:border-t-slate-300 hover:border-x-slate-300"
        on:click={handleClick(item.value)}>{item.label}</span
      >
    </li>
  {/each}
</ul>
{#each items as item}
  {#if activeTabValue == item.value}
    <svelte:component this={item.component} />
  {/if}
{/each}

<!-- REFERENCES:
    https://svelte.dev/playground/cf05bd4a4ca14fb8ace8b6cdebbb58da?version=5.1.7
    https://dev.to/barim/tuto-create-tabs-in-svelte-n82
    https://svelte.dev/playground/b9b5c6737f0045f3bf4693386046f356?version=4.2.2

    https://stackoverflow.com/questions/72826605/how-to-style-nested-elements-based-on-parent-class-using-tailwind-css
 -->

<!-- <style>
	.box {
		margin-bottom: 10px;
		padding: 40px;
		border: 1px solid #dee2e6;
    border-radius: 0 0 .5rem .5rem;
    border-top: 0;
	}
  ul {
    display: flex;
    flex-wrap: wrap;
    padding-left: 0;
    margin-bottom: 0;
    list-style: none;
    border-bottom: 1px solid #dee2e6;
  }
	li {
		margin-bottom: -1px;
	}

  span {
    border: 1px solid transparent;
    border-top-left-radius: 0.25rem;
    border-top-right-radius: 0.25rem;
    display: block;
    padding: 0.5rem 1rem;
    cursor: pointer;
  }

  span:hover {
    border-color: #e9ecef #e9ecef #dee2e6;
  }

  li.active > span {
    color: #495057;
    background-color: #fff;
    border-color: #dee2e6 #dee2e6 #fff;
  }
</style> -->
