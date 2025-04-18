<script lang="ts">
  import type { Review } from '@/types/Review.ts'
  import UserByIdReviewBottomContentLayout from '@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewBottomContent/UserByIdReviewBottomContentLayout.svelte'
  import UserByIdReviewColumn from '@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewColumn/UserByIdReviewColumn.svelte'
  import { getPackageDurationDifference } from '@/utils/utils'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import UserByIdReviewService from '@/pages/users/components/UserByIdReviewBlock/components/UserByIdReview/components/UserByIdReviewService/UserByIdReviewService.svelte'

  type UserByIdReviewBottomContentProps = {
    review: Review
    showServices?: boolean
  }

  const { review, showServices }: UserByIdReviewBottomContentProps = $props()
</script>

<UserByIdReviewBottomContentLayout>
  <div class="flex gap-y-1 md:gap-x-6 md:h-12 md:items-center flex-col md:flex-row items-start">
    <UserByIdReviewColumn label={'Price'} data={`$${review.service.price}`} />
    <DividerElement styles={'hidden md:block'} orientation={'vertical'} />
    <UserByIdReviewColumn
      label={'Duration'}
      data={getPackageDurationDifference(review.createdAt, review.endedAt)}
    />
    {#if showServices}
      <DividerElement styles={'hidden md:block'} orientation={'vertical'} />
      <p
        class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary mt-2 md:hidden"
      >
        Ordered
      </p>
      <UserByIdReviewService {...review.service} />
    {/if}
  </div>
</UserByIdReviewBottomContentLayout>
