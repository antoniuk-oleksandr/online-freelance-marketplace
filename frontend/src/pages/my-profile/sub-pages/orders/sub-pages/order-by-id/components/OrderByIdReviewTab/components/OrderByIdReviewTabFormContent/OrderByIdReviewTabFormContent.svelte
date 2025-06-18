<script lang="ts">
  import Title from '@/common-components/Title/Title.svelte'
  import OrderByIdReviewTabRatingSection from '../OrderByIdReviewTabRatingSection/OrderByIdReviewTabRatingSection.svelte'
  import OrderByIdReviewTabSubmitButton from '../OrderByIdReviewTabSubmitButton/OrderByIdReviewTabSubmitButton.svelte'
  import OrderByIdReviewTabMessageSection from '../OrderByIdReviewTabMessageSection/OrderByIdReviewTabMessageSection.svelte'
</script>

<Title text="Review your order" description="Share your experience with this service" />
<OrderByIdReviewTabRatingSection />
<OrderByIdReviewTabMessageSection />
<OrderByIdReviewTabSubmitButton />
