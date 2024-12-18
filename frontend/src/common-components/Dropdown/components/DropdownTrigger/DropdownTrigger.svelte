<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {handleDropdownTriggerClick} from "@/common-components/Dropdown/handlers.ts";
    import type {DropdownProps} from "@/types/DropdownProps.ts";

    type DropdownTriggerProps = LayoutProps & {
        isOpen: boolean,
        title: string,
        dropdownMenuProps: DropdownProps,
        setIsOpen?: (value: boolean) => void,
    }

    const {isOpen, setIsOpen, children, title, dropdownMenuProps}: DropdownTriggerProps = $props();
</script>

<button
        onclick={() => handleDropdownTriggerClick(isOpen, title, dropdownMenuProps, setIsOpen)}
        class="dropdown-trigger rounded-lg border bg-light-palette-background-block dark:bg-dark-palette-background-block border-light-palette-divider dark:border-dark-palette-divider">
    {@render children()}
</button>