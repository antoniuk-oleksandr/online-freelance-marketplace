<script lang="ts">
    import SearchFilterArrayBlockDataSelectButtonLayout
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataSelectButton/SearchFilterArrayBlockDataSelectButtonLayout.svelte";
    import Button from "@/common-components/Button/Button.svelte";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import {submitSearchPage} from "@/pages/search/helpers.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {searchFilterArrayStore} from "@/pages/search/stores/search-filter-array-store.ts";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";
    import {handleSearchArrSelectButtonClick} from "@/pages/search/handlers.ts";

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);
</script>

<SearchFilterArrayBlockDataSelectButtonLayout>
    <Button
            clickAction={() => handleSearchArrSelectButtonClick(searchPageParams)}
            color="grey"
            styles="w-full">
        Select
    </Button>
</SearchFilterArrayBlockDataSelectButtonLayout>