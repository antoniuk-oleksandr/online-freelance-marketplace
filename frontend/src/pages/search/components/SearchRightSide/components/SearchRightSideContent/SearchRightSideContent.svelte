<script lang="ts">
    import SearchSubmitResetButtonBlock
        from "@/pages/search/components/SearchSubmitResetButtonBlock/SearchSubmitResetButtonBlock.svelte";
    import SearchFilterArrayBlockList
        from "@/pages/search/components/SearchFilterArrayBlockList/SearchFilterArrayBlockList.svelte";
    import SearchRightSideCloseButton
        from "@/pages/search/components/SearchRightSideCloseButton/SearchRightSideCloseButton.svelte";
    import SearchFromToBlockList from "@/pages/search/components/SearchFromToBlockList/SearchFromToBlockList.svelte";
    import type {GetFilterParamsRequestResponse} from "@/types/GetFilterParamsRequestResponse.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import SearchRightSideContentLayout
        from "@/pages/search/components/SearchRightSide/components/SearchRightSideContent/SearchRightSideContentLayout.svelte";
    import SearchRightSideHeader from "@/pages/search/components/SearchRightSideHeader/SearchRightSideHeader.svelte";

    type SearchRightSideProps = {
        isFiltersModalOpen: boolean,
        defaultFilterParams: GetFilterParamsRequestResponse["data"],
    }

    const props: SearchRightSideProps = $props();

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);
</script>

{#if searchPageParams}
    <SearchRightSideContentLayout>
        <SearchRightSideHeader/>
        <SearchRightSideCloseButton {...props}/>
        <SearchFilterArrayBlockList
                searchPageParams={searchPageParams}
                defaultFilterParams={props.defaultFilterParams}
        />
        <SearchFromToBlockList searchPageParams={searchPageParams}/>
        <SearchSubmitResetButtonBlock searchPageParams={searchPageParams}/>
    </SearchRightSideContentLayout>
{/if}
