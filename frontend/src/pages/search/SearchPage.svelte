<script lang="ts">
    import {getSearchPageParams} from "@/pages/search/helpers.ts";
    import SearchPageLayout from "@/pages/search/SearchPageLayout.svelte";
    import SearchLeftSide from "@/pages/search/components/SearchLeftSide/SearchLeftSide.svelte";
    import SearchRightSide from "@/pages/search/components/SearchRightSide/SearchRightSide.svelte";
    import {useRouter} from "svelte-routing";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {getFilterParamsRequest} from "@/api/get-filter-params-request.ts";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";
    import {getSearchRequest} from "@/api/get-search-request.ts";

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);

    let searchRequestResponse = $state();

    useRouter().routerBase.subscribe(() => {
        const params = getSearchPageParams();

        searchStore.set(params);
        searchRequestResponse = getSearchRequest();
    });

    let defaultFilterParams = $state<GetFilterParamsRequestResponse | undefined>();

    getFilterParamsRequest().then((response) => {
        defaultFilterParams = response;

    });

    let isFiltersModalOpen = $state(false);
    searchFilterDrawerStore.subscribe((value) => isFiltersModalOpen = value);
</script>

{#if searchPageParams && defaultFilterParams && defaultFilterParams.status === 200}
    <SearchPageLayout>
        <SearchLeftSide
                searchRequestResponse={searchRequestResponse}/>
        <SearchRightSide
                defaultFilterParams={defaultFilterParams.data}
                isFiltersModalOpen={isFiltersModalOpen}
        />
    </SearchPageLayout>
{/if}
