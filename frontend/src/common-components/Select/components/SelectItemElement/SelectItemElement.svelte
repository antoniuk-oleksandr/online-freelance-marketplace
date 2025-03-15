<script lang="ts">
    import type {SelectItem} from "@/types/SelectItem";
    import Icon from "@iconify/svelte";
    import {handleSelectClick} from "@/common-components/Select/handlers";

    type SelectItemElementProps = SelectItem & {
        selectedItem: string,
        setSelectedItem?: (item: string) => void,
    };

    const {title, clickAction, selectedItem, setSelectedItem}: SelectItemElementProps = $props();
</script>

<button
        class="text-start text-base flex gap-x-3 items-center hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover px-3 rounded-[4px] duration-200 ease-out py-1"
        onclick={() => handleSelectClick(clickAction, title, setSelectedItem)}
>
    {#if selectedItem === title}
        <Icon class="text-cyan-500" icon="akar-icons:radio-fill" width="16" height="16"/>
    {:else}
        <Icon icon="akar-icons:radio" width="16" height="16"/>
    {/if}
    <span>{title}</span>
</button>
