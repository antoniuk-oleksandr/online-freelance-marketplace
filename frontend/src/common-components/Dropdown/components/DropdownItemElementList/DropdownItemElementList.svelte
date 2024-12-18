<script lang="ts">
    import DropdownItemElement
        from "@/common-components/Dropdown/components/DropdownItemElement/DropdownItemElement.svelte";
    import type {DropdownItem} from "@/types/DropdownItem.ts";

    type DropdownItemElementListProps = {
        items: DropdownItem[],
        setIsOpen?: (value: boolean) => void,
        selectedItem: string,
        setSelectedItem?: (item: string) => void,
    };

    const props: DropdownItemElementListProps = $props();
</script>

<div class="flex flex-col gap-y-1 dropdown-menu py-1 lg:py-0">
    {#each props.items as item}
        <DropdownItemElement {...props} {...item}/>
    {/each}
</div>
