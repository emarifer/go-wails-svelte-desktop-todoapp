<script lang="ts">
    import Swal from "sweetalert2";
    import Todo from "./Todo.svelte";
    import RemoveButton from "./RemoveButton.svelte";
    import { main } from "../wailsjs/go/models";

    export let todos: main.Task[] = [];
    export let removeTodo = (i: number): void => {};
    export let handleChange = (status: boolean, id: number): void => {};
</script>

<!-- TODO: should find a way to invoke removeMe or get the index after button click -->
<ul
    class="w-4/5 mx-auto max-h-64 overflow-y-auto flex flex-col gap-2 bg-slate-700 rounded-xl py-4 px-12"
>
    {#each todos as todo (todo.id)}
        <li class="block">
            <div class="flex justify-between items-center">
                <Todo task={todo} toggle={handleChange} />
                <RemoveButton
                    on:click={() => {
                        Swal.fire({
                            title: "Do you want to perform this action?",
                            text: `Are you sure you want to delete the task with ID #${todo.id}?`,
                            icon: "warning",
                            background: "#1D232A",
                            color: "#A6ADBA",
                            showCancelButton: true,
                            confirmButtonColor: "#3085d6",
                            cancelButtonColor: "#d33",
                            confirmButtonText: "Yes, delete it!",
                        }).then((result) => {
                            if (result.value) {
                                removeTodo(todo.id);
                                Swal.fire({
                                    title: "Deleted!",
                                    text: "The task has been deleted",
                                    icon: "success",
                                    background: "#1D232A",
                                    color: "#A6ADBA",
                                    confirmButtonColor: "#3085d6",
                                });
                            }
                        });
                    }}
                />
            </div>
        </li>
    {/each}
</ul>

<!-- 
https://svelte.dev/tutorial/svelte/keyed-each-blocks
https://gist.github.com/collardeau/6a0c9777246db4f7b1764b3ccafdf822
 -->
