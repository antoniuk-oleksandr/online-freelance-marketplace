<script lang="ts">
    import UserByIdService
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/UserByIdService.svelte";
    import type {UserService} from "@/types/UserService.ts";
    import {handleSearchScroll} from "@/pages/search/handlers.ts";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import SearchServicesLayout from "@/pages/search/components/SearchServices/SearchServicesLayout.svelte";
    import SearchServicesNotFound from "@/pages/search/components/SearchServicesNotFound/SearchServicesNotFound.svelte";

    type UserServicesProps = {
        searchRequestResponse: any;
    }

    let reached = $state(false);
    const setReached = (value: boolean) => reached = value;

    const {searchRequestResponse}: UserServicesProps = $props();

    $effect(() => {
        const handleFunc = () => handleSearchScroll(reached, setReached);
        document.addEventListener("scroll", handleFunc);

        return () => document.removeEventListener("scroll", handleFunc);
    })

</script>

{#if searchRequestResponse.status === 404}
    <SearchServicesNotFound/>
{:else }
    <SearchServicesLayout>
        {#each searchRequestResponse.data.services as service}
            <UserByIdService
                    size="small"
                    service={service}
            />
        {/each}
    </SearchServicesLayout>
{/if}
