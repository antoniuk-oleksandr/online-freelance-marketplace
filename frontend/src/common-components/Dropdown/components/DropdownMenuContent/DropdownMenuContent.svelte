<script lang="ts">
    import type {DropdownMenuProps} from "@/types/DropdownMenuProps.ts";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import DropdownItemElementList
        from "@/common-components/Dropdown/components/DropdownItemElementList/DropdownItemElementList.svelte";
    import {handleDropdownBackdropClick} from "@/common-components/Dropdown/handlers.ts";

    const props: DropdownMenuProps = $props();

    const handleClick = (e: any) => handleDropdownBackdropClick(e, props.setIsOpen);

    $effect(() => {
        document.addEventListener("click", handleClick);

        return () => document.removeEventListener("click", handleClick);
    })
</script>

<DropdownItemElementList
        items={props.items}
        setIsOpen={props.setIsOpen}
        selectedItem={props.selectedItem}
        setSelectedItem={props.setSelectedItem}
/>
{#if props.additionalItems && props.selectedAdditionalItem}
    <DividerElement styles="dropdown-menu"/>
    <DropdownItemElementList
            selectedItem={props.selectedAdditionalItem}
            setSelectedItem={props.setSelectedAdditionalItem}
            items={props.additionalItems}
            setIsOpen={props.setIsOpen}
    />
{/if}