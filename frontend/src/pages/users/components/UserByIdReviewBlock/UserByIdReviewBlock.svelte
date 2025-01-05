<script lang="ts">
    import type {Review} from "@/types/Review.ts";
    import UserComponentLayout from "@/pages/users/components/UserComponentLayout.svelte";
    import NoReviewsMessage
        from "@/pages/users/components/UserByIdReviewBlock/components/NoReviewsMessage/NoReviewsMessage.svelte";
    import UserByIdReview
        from "@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/UserByIdReview.svelte";
    import Button from "@/common-components/Button/Button.svelte";
    import {handleMoreReviewsButtonClick} from "@/pages/users/handlers.ts";
    import ShowMoreReviewsButton
        from "@/pages/users/components/UserByIdReviewBlock/components/ShowMoreReviewsButton/ShowMoreReviewsButton.svelte";

    type UserByIdReviewProps = {
        reviews: Review[] | null,
        showMoreReviewsButtonAction: () => Promise<void>,
        hasMore: boolean,
        showServices?: boolean,
    }

    const {reviews, showServices, showMoreReviewsButtonAction, hasMore}: UserByIdReviewProps = $props();
</script>

<UserComponentLayout>
    <p class="text-xl font-bold">Reviews</p>
    <div class="flex flex-col gap-y-3">
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
        <ShowMoreReviewsButton
                buttonSuffix="reviews"
                showMoreReviewsButtonAction={showMoreReviewsButtonAction}
                hasMore={hasMore}
        />
    </div>
</UserComponentLayout>
