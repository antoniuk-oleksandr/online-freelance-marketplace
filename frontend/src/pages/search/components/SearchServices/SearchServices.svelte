<script lang="ts">
    import UserByIdService
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/UserByIdService.svelte";
    import SearchServicesLayout from "@/pages/search/components/SearchServices/SearchServicesLayout.svelte";
    import SearchServicesNotFound from "@/pages/search/components/SearchServicesNotFound/SearchServicesNotFound.svelte";
    import type {SearchCursorStore} from "@/types/SearchCursorStore.ts";
    //@ts-ignore
    import {searchCursorStore} from "@/pages/search/stores/search-cursor-store.ts";

    let searchCursorData = $state<SearchCursorStore | undefined>();
    searchCursorStore.subscribe((value) => searchCursorData = value);
</script>

{#if !searchCursorData || !searchCursorData.searchRequestResponse}
    {null}
{:else if searchCursorData.searchRequestResponse.status !== 200}
    {null}
{:else if !searchCursorData.searchRequestResponse.data.services || searchCursorData.searchRequestResponse.data.services.length == 0}
    <SearchServicesNotFound/>
{:else}
    <SearchServicesLayout>
        {#each searchCursorData.searchRequestResponse.data.services as service}
            <UserByIdService
                    size="small"
                    service={service}
            />
        {/each}
    </SearchServicesLayout>
{/if}