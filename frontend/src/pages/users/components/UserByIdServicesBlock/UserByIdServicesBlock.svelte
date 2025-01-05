<script lang="ts">
    import UserComponentLayout from "@/pages/users/components/UserComponentLayout.svelte";
    import UserByIdService
        from "@/pages/users/components/UserByIdServicesBlock/components/UserByIdService/UserByIdService.svelte";
    import type {UserService} from "@/types/UserService.ts";
    import ShowMoreReviewsButton
        from "@/pages/users/components/UserByIdReviewBlock/components/ShowMoreReviewsButton/ShowMoreReviewsButton.svelte";

    type UserByIdServicesBlockProps = {
        services: UserService[] | null,
        hasMore: boolean,
        showMoreServicesButtonAction: () => Promise<void>,
    }

    const {services, showMoreServicesButtonAction, hasMore}: UserByIdServicesBlockProps = $props();
</script>

{#if services && services.length > 0}
    <UserComponentLayout>
        <h2 class="text-xl font-bold">Services</h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
            {#each services as service}
                <UserByIdService service={service}/>
            {/each}
        </div>
        <ShowMoreReviewsButton
                buttonSuffix="services"
                showMoreReviewsButtonAction={showMoreServicesButtonAction}
                hasMore={hasMore}
        />
    </UserComponentLayout>
{/if}