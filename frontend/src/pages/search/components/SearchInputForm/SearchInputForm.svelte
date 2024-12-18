<script lang="ts">
    import {searchStore} from "@/pages/search/stores/search-store.ts";
    import SearchInputFormLayout from "@/pages/search/components/SearchInputForm/SearchInputFormLayout.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import SearchInputFormInput from "@/pages/search/components/SearchInputFormInput/SearchInputFormInput.svelte";
    import SearchInputFormSubmitButton
        from "@/pages/search/components/SearchInputFormSubmitButton/SearchInputFormSubmitButton.svelte";

    let searchPageParams = $state<SearchPageParams | undefined>();
    searchStore.subscribe((value) => searchPageParams = value);
</script>

{#if searchPageParams}
    <SearchInputFormLayout searchPageParams={searchPageParams}>
        <SearchInputFormInput query={searchPageParams.query}/>
        <SearchInputFormSubmitButton/>
    </SearchInputFormLayout>
{/if}
