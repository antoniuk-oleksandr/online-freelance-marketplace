<script lang="ts">
    import {resetSearchPage, submitSearchPage} from "@/pages/search/helpers.js";
    import Button from "@/common-components/Button/Button.svelte";
    import SearchSubmitResetButtonBlockLayout
        from "@/pages/search/components/SearchSubmitResetButtonBlock/SearchSubmitResetButtonBlockLayout.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import {searchFilterDrawerStore} from "@/pages/search/stores/search-filter-drawer-store.ts";

    type SearchSubmitButtonBlockProps = {
        searchPageParams: SearchPageParams
    }

    const {searchPageParams}: SearchSubmitButtonBlockProps = $props();
</script>

<SearchSubmitResetButtonBlockLayout>
    <Button
            styles="w-full"
            color="outline"
            clickAction={resetSearchPage}
    >Reset
    </Button>
    <Button
            styles="w-full"
            clickAction={() => {
                searchFilterDrawerStore.set(false);
                submitSearchPage(searchPageParams);
            }}
    >Submit
    </Button>
</SearchSubmitResetButtonBlockLayout>
