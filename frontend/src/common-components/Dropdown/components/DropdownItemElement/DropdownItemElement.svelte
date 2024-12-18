<script lang="ts">
    import type {DropdownItem} from "@/types/DropdownItem.ts";
    import Icon from "@iconify/svelte";
    import {handleDropdownClick} from "@/common-components/Dropdown/handlers.ts";

    type DropdownItemElementProps = DropdownItem & {
        selectedItem: string,
        setSelectedItem?: (item: string) => void,
    };

    const {title, clickAction, selectedItem, setSelectedItem}: DropdownItemElementProps = $props();
</script>

<button
        class="text-start text-base flex gap-x-2 items-center hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover px-3 rounded-[4px] duration-200 ease-out py-1"
        onclick={() => handleDropdownClick(clickAction, title, setSelectedItem)}
>
    {#if selectedItem === title}
        <Icon class="text-cyan-500" icon="akar-icons:radio-fill" width="16" height="16"/>
    {:else}
        <Icon icon="akar-icons:radio" width="16" height="16"/>
    {/if}
    <span>{title}</span>
</button>