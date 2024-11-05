<script lang="ts">
    import { main } from "../wailsjs/go/models";
    import { _ } from "svelte-i18n";

    export let task = new main.Task();
    export let checked = task.status;
    export let toggle = (st: boolean, id: number) => {};
</script>

<div class="flex justify-between items-center w-full">
    <div class="flex gap-2 justify-start items-center w-fit">
        <label class="cursor-pointer label">
            <input
                title={$_("checkbox")}
                bind:checked
                on:change={() => toggle(checked, task.id)}
                type="checkbox"
                class="checkbox checkbox-info"
            />
        </label>
        <div
            class="block badge badge-accent badge-outline max-w-56 text-ellipsis overflow-hidden whitespace-nowrap text-left px-2"
            title={task.title}
        >
            {task.title}
        </div>
    </div>
    <span class="text-xs font-light text-lime-300 pr-8">{task.created_at}</span>
</div>

<!-- 
    https://stackoverflow.com/questions/76993944/how-to-force-the-contents-of-a-div-to-be-on-one-line
    https://tailwindcss.com/docs/text-overflow#ellipsis

    https://stackoverflow.com/questions/68134909/svelte-use-function-from-a-parent-component
    https://hackernoon.com/svelte-communication-between-components-193i32l5
 -->
