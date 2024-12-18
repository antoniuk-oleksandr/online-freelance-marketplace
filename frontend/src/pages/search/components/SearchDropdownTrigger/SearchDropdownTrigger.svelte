<script lang="ts">
    import SearchDropdownTriggerLayout
        from "@/pages/search/components/SearchDropdownTrigger/SearchDropdownTriggerLayout.svelte";
    import SearchDropdownTriggerIcon
        from "@/pages/search/components/SearchDropdownTriggerIcon/SearchDropdownTriggerIcon.svelte";

    type SearchDropdownTriggerProps = {
        selectedAdditionalItem: string,
        selectedItem: string,
    }

    const {selectedAdditionalItem, selectedItem}: SearchDropdownTriggerProps = $props();
</script>

<SearchDropdownTriggerLayout>
    <SearchDropdownTriggerIcon selectedAdditionalItem={selectedAdditionalItem}/>
    <span>{selectedItem}</span>
</SearchDropdownTriggerLayout>