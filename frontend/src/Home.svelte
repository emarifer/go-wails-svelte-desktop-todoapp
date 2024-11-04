<script lang="ts">
    import Swal from "sweetalert2";
    import { onMount } from "svelte";
    import { fade, fly } from "svelte/transition";
    import {
        AddTodo,
        GetAllTodos,
        RemoveTodo,
        UpadateTodo,
    } from "../wailsjs/go/main/App.js";
    import Todos from "./Todos.svelte";
    import { main } from "../wailsjs/go/models";

    import logo from "./assets/images/logo-universal.png";

    let visible = false;
    let allTodos: main.Task[] = [];
    let inputRef = null;

    onMount(async () => {
        visible = !visible;
        GetAllTodos().then((result) => (allTodos = result));
    });

    // let allTodos: string[] = ["Hello", "World", "!"]
    let newTodo: string = "";

    function addTodo(): void {
        if (newTodo.length < 3) {
            newTodo = "";
            Swal.fire({
                title: "Incorrect entry",
                text: "You have entered less than 3 characters",
                icon: "info",
                background: "#1D232A",
                color: "#A6ADBA",
                confirmButtonColor: "#3085d6",
                didClose: () => inputRef.focus(),
            });
            return;
        }
        AddTodo(newTodo).then((result) => (allTodos = result));
        newTodo = "";
        inputRef.focus();
    }

    function removeTodo(i: number): void {
        console.log(`removing todo: ${i}`);
        RemoveTodo(i).then((result) => (allTodos = result));
    }

    function handleChange(status: boolean, id: number): void {
        console.log(status);
        UpadateTodo(status, id).then((result) => (allTodos = result));
        inputRef.focus();
    }
</script>

{#if visible}
    <div
        in:fly={{ x: 75, duration: 1200 }}
        out:fade={{ duration: 20 }}
        class="flex flex-col items-center gap-4"
    >
        <img alt="Wails logo" src={logo} class="h-16 block mx-auto" />
        <h1 class="text-3xl text-sky-600 font-bold">Wails Todoapp</h1>
        <div class="flex gap-4 mt-6 w-fit mx-auto" id="input">
            <!-- svelte-ignore a11y-autofocus -->
            <input
                bind:this={inputRef}
                bind:value={newTodo}
                class="input input-bordered input-primary bg-slate-800"
                id="newTodo"
                type="text"
                placeholder="Enter a task (min 3 chars)…"
                autofocus
            />
            <button
                title="Add Task"
                class="btn btn-outline btn-primary"
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
            <p class="text-lg">Add some todos 😀</p>
        {:else}
            <Todos
                ref={inputRef}
                todos={allTodos.reverse()}
                {removeTodo}
                {handleChange}
            />
        {/if}
    </div>
{/if}

<!-- 
https://svelte.dev/tutorial/svelte/keyed-each-blocks
https://gist.github.com/collardeau/6a0c9777246db4f7b1764b3ccafdf822

HOW TO FOCUS ON INPUT WITH SWEETALERT2:
https://es.stackoverflow.com/questions/416336/como-hacer-focus-en-input-con-sweetalert2-v10

TRANSITIONS:
https://svelte.dev/playground/in-and-out
https://www.youtube.com/watch?v=5oEo98BrRqs
 -->
