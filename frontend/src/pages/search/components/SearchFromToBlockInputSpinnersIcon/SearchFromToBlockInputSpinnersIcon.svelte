<script lang="ts">
    import Icon from "@iconify/svelte";
    import {handleSearchFromToBlockInputSpinnersClick} from "@/pages/search/handlers.ts";

    type SearchFromToBlockInputProps = {
        value: number,
        setValue: (value: number) => void,
        direction: "up" | "down",
    }

    const {value, setValue, direction}: SearchFromToBlockInputProps = $props();
</script>

<Icon
        icon="fe:arrow-{direction}"
        width={24}
        height={12}
        class="hover:bg-dark-palette-action-selected dark:hover:bg-dark-palette-action-selected duration-200 ease-out cursor-pointer bg-light-palette-action-hover dark:bg-dark-palette-action-hover"
        onclick={() => handleSearchFromToBlockInputSpinnersClick(value, setValue, direction)}
/>
