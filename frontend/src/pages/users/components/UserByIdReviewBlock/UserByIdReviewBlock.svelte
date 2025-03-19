<script lang="ts">
  import type { Review } from '@/types/Review.ts'
  import UserComponentLayout from '@/pages/users/components/UserComponentLayout.svelte'
  import UserByIdReview from '@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/UserByIdReview.svelte'
  import ShowMoreReviewsButton from '@/pages/users/components/UserByIdReviewBlock/components/ShowMoreReviewsButton/ShowMoreReviewsButton.svelte'
  import NoContentMessage from '@/common-components/NoContentMessage/NoContentMessage.svelte'

  type UserByIdReviewProps = {
    reviews: Review[] | null
    showMoreReviewsButtonAction: () => Promise<void>
    hasMore: boolean
    showServices?: boolean
  }

  const { reviews, showServices, showMoreReviewsButtonAction, hasMore }: UserByIdReviewProps =
    $props()
</script>

<UserComponentLayout>
  <p class="text-xl font-bold">Reviews</p>
  <div class="flex flex-col gap-y-3">
    {#if !reviews || reviews.length === 0}
      <NoContentMessage text="No reviews yet" />
    {:else}
      {#each reviews as review}
        <UserByIdReview {review} {showServices} />
      {/each}
    {/if}
    <ShowMoreReviewsButton buttonSuffix="reviews" {showMoreReviewsButtonAction} {hasMore} />
  </div>
</UserComponentLayout>
