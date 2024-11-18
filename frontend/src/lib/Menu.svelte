<script lang="ts">
    import Swal from "sweetalert2";
    import { _ } from "svelte-i18n";
    import { EventsEmit, EventsOn } from "../../wailsjs/runtime/runtime";

    EventsOn("exported", () =>
        Swal.fire({
            title: $_("data_export"),
            text: $_("data_export_text"),
            icon: "success",
            background: "#1D232A",
            color: "#A6ADBA",
            confirmButtonColor: "#3085d6",
        }),
    );

    function deleteAll() {
        Swal.fire({
            title: $_("erasing"),
            text: $_("erasing_text"),
            icon: "warning",
            background: "#1D232A",
            color: "#A6ADBA",
            showCancelButton: true,
            confirmButtonColor: "#3085d6",
            cancelButtonColor: "#d33",
            confirmButtonText: $_("erasing_confirm"),
        }).then((result) => {
            if (result.value) {
                EventsEmit("delete-all");
            }
        });
    }

    function importData() {
        Swal.fire({
            title: $_("import_title"),
            text: $_("import_text"),
            icon: "warning",
            background: "#1D232A",
            color: "#A6ADBA",
            showCancelButton: true,
            confirmButtonColor: "#3085d6",
            cancelButtonColor: "#d33",
            confirmButtonText: $_("erasing_confirm"),
        }).then((result) => {
            if (result.value) {
                EventsEmit("import");
            }
        });
    }

    function onKeyDown(e: KeyboardEvent) {
        // console.log(e.key);
        switch (e.key) {
            case "F12":
                deleteAll();
                break;
            case "F1":
                EventsEmit("export");
                break;
            case "F2":
                importData();
                break;
            case "Escape":
                EventsEmit("quit");
                break;
        }
    }
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_consider_explicit_label -->
<!-- svelte-ignore a11y_missing_attribute -->
<ul class="absolute top-14 left-2 menu bg-base-200 rounded-box">
    <li>
        <a
            onclick={deleteAll}
            class="tooltip tooltip-right"
            data-tip={`${$_("delete_all_data")} 'F12'`}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                viewBox="0 0 16 16"
            >
                <path
                    d="M.54 3.87.5 3a2 2 0 0 1 2-2h3.672a2 2 0 0 1 1.414.586l.828.828A2 2 0 0 0 9.828 3h3.982a2 2 0 0 1 1.992 2.181L15.546 8H14.54l.265-2.91A1 1 0 0 0 13.81 4H2.19a1 1 0 0 0-.996 1.09l.637 7a1 1 0 0 0 .995.91H9v1H2.826a2 2 0 0 1-1.991-1.819l-.637-7a2 2 0 0 1 .342-1.31zm6.339-1.577A1 1 0 0 0 6.172 2H2.5a1 1 0 0 0-1 .981l.006.139q.323-.119.684-.12h5.396z"
                />
                <path
                    d="M11.854 10.146a.5.5 0 0 0-.707.708L12.293 12l-1.146 1.146a.5.5 0 0 0 .707.708L13 12.707l1.146 1.147a.5.5 0 0 0 .708-.708L13.707 12l1.147-1.146a.5.5 0 0 0-.707-.708L13 11.293z"
                />
            </svg>
        </a>
    </li>
    <li>
        <a class="tooltip tooltip-right" data-tip={`${$_("export_data")} 'F1'`}>
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                viewBox="0 0 16 16"
            >
                <path
                    fill-rule="evenodd"
                    d="M10 12.5a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v2a.5.5 0 0 0 1 0v-2A1.5 1.5 0 0 0 9.5 2h-8A1.5 1.5 0 0 0 0 3.5v9A1.5 1.5 0 0 0 1.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0z"
                />
                <path
                    fill-rule="evenodd"
                    d="M15.854 8.354a.5.5 0 0 0 0-.708l-3-3a.5.5 0 0 0-.708.708L14.293 7.5H5.5a.5.5 0 0 0 0 1h8.793l-2.147 2.146a.5.5 0 0 0 .708.708z"
                />
            </svg>
        </a>
    </li>
    <li>
        <a
            onclick={importData}
            class="tooltip tooltip-right"
            data-tip={`${$_("import_data")} 'F2'`}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                viewBox="0 0 16 16"
            >
                <path
                    fill-rule="evenodd"
                    d="M10 3.5a.5.5 0 0 0-.5-.5h-8a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 1 1 0v2A1.5 1.5 0 0 1 9.5 14h-8A1.5 1.5 0 0 1 0 12.5v-9A1.5 1.5 0 0 1 1.5 2h8A1.5 1.5 0 0 1 11 3.5v2a.5.5 0 0 1-1 0z"
                />
                <path
                    fill-rule="evenodd"
                    d="M4.146 8.354a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L5.707 7.5H14.5a.5.5 0 0 1 0 1H5.707l2.147 2.146a.5.5 0 0 1-.708.708z"
                />
            </svg>
        </a>
    </li>
    <li>
        <a
            onclick={() => EventsEmit("quit")}
            class="tooltip tooltip-right"
            data-tip={`${$_("quit")} 'Escape'`}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                viewBox="0 0 16 16"
            >
                <path
                    d="M8.538 1.02a.5.5 0 1 0-.076.998 6 6 0 1 1-6.445 6.444.5.5 0 0 0-.997.076A7 7 0 1 0 8.538 1.02"
                />
                <path
                    d="M7.096 7.828a.5.5 0 0 0 .707-.707L2.707 2.025h2.768a.5.5 0 1 0 0-1H1.5a.5.5 0 0 0-.5.5V5.5a.5.5 0 0 0 1 0V2.732z"
                />
            </svg>
        </a>
    </li>
</ul>

<svelte:window on:keydown={onKeyDown} />

<!-- REFERENCES:
    https://github.com/tmclane/wails-events
-->
