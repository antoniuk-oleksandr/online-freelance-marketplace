<script lang="ts">
    import SearchPageLayout from "@/pages/search/SearchPageLayout.svelte";
    import SearchLeftSide from "@/pages/search/components/SearchLeftSide/SearchLeftSide.svelte";
    import SearchRightSide from "@/pages/search/components/SearchRightSide/SearchRightSide.svelte";
    import {useRouter} from "svelte-routing";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {getFilterParamsRequest} from "@/api/get-filter-params-request.ts";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";
    import type {SearchRequestResponse} from "@/types/SearchRequestResponse.ts";
    import {searchCursorStore} from "@/pages/search/stores/search-cursor-store.ts";
    import type {SearchCursorStore} from "@/types/SearchCursorStore.ts";
    import {handleSearchRefresh} from "@/pages/search/handlers.ts";

    let searchCursorData = $state<SearchCursorStore | undefined>();
    searchCursorStore.subscribe((value) => searchCursorData = value);

    useRouter().routerBase.subscribe(() => {
        if (!searchCursorData) return;
        handleSearchRefresh();
    });

    let defaultFilterParams = $state<GetFilterParamsRequestResponse | undefined>();

    getFilterParamsRequest().then((response) => {
        defaultFilterParams = response;
    });

    let isFiltersModalOpen = $state(false);
    searchFilterDrawerStore.subscribe((value) => isFiltersModalOpen = value);
</script>

{#if defaultFilterParams && defaultFilterParams.status === 200}
    <SearchPageLayout>
        <SearchLeftSide/>
        <SearchRightSide
                defaultFilterParams={defaultFilterParams.data}
                isFiltersModalOpen={isFiltersModalOpen}
        />
    </SearchPageLayout>
{/if}
