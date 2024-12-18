<script lang="ts">
    import type {SearchFilterArrayStore} from "@/types/SearchFilterArrayStore.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import {resetSearchArrayParam, submitSearchPage} from "@/pages/search/helpers.ts";

    type SearchFilterArrayBlockDataTopLineResetButtonProps = {
        storeData: SearchFilterArrayStore

    }

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);

    const {storeData}: SearchFilterArrayBlockDataTopLineResetButtonProps = $props();
</script>

<button
        onclick={() => {
            resetSearchArrayParam(storeData)
            submitSearchPage(searchPageParams);
        }}
>reset
</button>
