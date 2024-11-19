<script lang="ts">
    import type {Review} from "@/types/Review.ts";
    import UserComponentLayout from "@/pages/users/components/UserComponentLayout.svelte";
    import NoReviewsMessage
        from "@/pages/users/components/UserByIdReviewBlock/components/NoReviewsMessage/NoReviewsMessage.svelte";
    import UserByIdReview
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/UserByIdReview.svelte";

    type UserByIdReviewProps = {
        reviews: Review[] | null,
        showServices?: boolean
    }

    const {reviews, showServices}: UserByIdReviewProps = $props();
</script>

<UserComponentLayout>
    <p class="text-xl font-bold">Reviews</p>
    <div class="flex flex-col gap-y-8">
        {#if !reviews || reviews.length === 0}
            <NoReviewsMessage/>
        {:else}
            {#each reviews as review}
                <UserByIdReview
                        review={review}
                        showServices={showServices}
                />
            {/each}
        {/if}
    </div>
</UserComponentLayout>
