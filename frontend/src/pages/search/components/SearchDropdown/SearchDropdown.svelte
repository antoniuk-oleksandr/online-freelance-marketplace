<script lang="ts">
    import Dropdown from "@/common-components/Dropdown/Dropdown.svelte";
    import {getSelectedSearchDropdownItem, makeASearchDropdownItemList} from "@/pages/search/helpers.ts";
    import SearchDropdownTrigger from "@/pages/search/components/SearchDropdownTrigger/SearchDropdownTrigger.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {searchStore} from "@/pages/search/stores/search-store.ts";

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);

    const items = $derived(makeASearchDropdownItemList(searchPageParams));
    const additionalItems = $derived(makeASearchDropdownItemList(searchPageParams, true));
</script>

<Dropdown
        title="Sorting"
        selectedItem={getSelectedSearchDropdownItem(searchPageParams)}
        items={items}
        additionalItems={additionalItems}
        selectedAdditionalItem={getSelectedSearchDropdownItem(searchPageParams, true)}
>
    <SearchDropdownTrigger
            selectedItem={getSelectedSearchDropdownItem(searchPageParams)}
            selectedAdditionalItem={getSelectedSearchDropdownItem(searchPageParams, true)}
    />
</Dropdown>

