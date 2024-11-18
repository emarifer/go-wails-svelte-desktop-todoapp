<script lang="ts">
    import Swal from "sweetalert2";
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import { _ } from "svelte-i18n";
    import logo from "../assets/images/logo-universal.png";
    import Todos from "../lib/Todos.svelte";
    import {
        AddTodo,
        GetAllTodos,
        RemoveTodo,
        UpadateTodo,
    } from "../../wailsjs/go/main/App";
    import { main } from "../../wailsjs/go/models";

    let mounted: boolean = false;
    let allTodos: main.Task[] = [];
    let inputRef: HTMLInputElement | null = null;
    let newTodo: string = "";

    onMount(() => {
        mounted = true;
        GetAllTodos().then((result) => (allTodos = result));
    });

    function addTodo(): void {
        if (newTodo.length < 3) {
            newTodo = "";
            Swal.fire({
                title: $_("incorrect_title"),
                text: $_("incorrect_text"),
                icon: "info",
                background: "#1D232A",
                color: "#A6ADBA",
                confirmButtonColor: "#3085d6",
                didClose: () => inputRef?.focus(),
            });
            return;
        }
        AddTodo(newTodo).then((result) => (allTodos = result));
        newTodo = "";
        inputRef?.focus();
    }

    function removeTodo(i: number): void {
        console.log(`removing todo: ${i}`);
        RemoveTodo(i).then((result) => (allTodos = result));
    }

    function handleChange(status: boolean, id: number): void {
        console.log(status);
        UpadateTodo(status, id).then((result) => (allTodos = result));
        inputRef?.focus();
    }
</script>

<!-- svelte-ignore a11y_autofocus -->
<!-- svelte-ignore a11y_consider_explicit_label -->
{#if mounted}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 200 }}
        class="flex flex-col items-center gap-4"
    >
        <img alt="Wails logo" src={logo} class="mt-1.5 h-16 block mx-auto" />
        <h1 class="text-3xl text-sky-600 font-bold">Wails Todoapp</h1>
        <div class="flex gap-4 mt-6 w-fit mx-auto" id="input">
            <input
                bind:this={inputRef}
                bind:value={newTodo}
                class="input input-bordered input-primary bg-slate-800"
                id="newTodo"
                type="text"
                placeholder={$_("placeholder")}
                autofocus
            />
            <button
                data-tip={$_("add_task")}
                class="btn btn-outline btn-primary tooltip"
                on:click={addTodo}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="currentColor"
                    class="fill-white"
                    viewBox="0 0 16 16"
                >
                    <path
                        d="M9.293 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4.707A1 1 0 0 0 13.707 4L10 .293A1 1 0 0 0 9.293 0M9.5 3.5v-2l3 3h-2a1 1 0 0 1-1-1M8.5 7v1.5H10a.5.5 0 0 1 0 1H8.5V11a.5.5 0 0 1-1 0V9.5H6a.5.5 0 0 1 0-1h1.5V7a.5.5 0 0 1 1 0"
                    />
                </svg>
            </button>
        </div>
        {#if allTodos.length === 0}
            <p class="text-lg">{$_("add_some_todos")} 😀</p>
        {:else}
            <Todos
                {inputRef}
                todos={allTodos.reverse()}
                {removeTodo}
                {handleChange}
            />
        {/if}
    </div>
{/if}

<!-- https://github.com/sveltejs/svelte/issues/11841 -->
